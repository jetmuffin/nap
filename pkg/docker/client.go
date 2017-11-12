package docker

import (
	"strings"
	"fmt"
	"net/url"
	"net/http"
	"github.com/Sirupsen/logrus"
	"io/ioutil"
	"encoding/json"
	"time"
	"net/http/httputil"
	"net"
	"errors"
	. "github.com/JetMuffin/nap/pkg/types"
)

type DockerClient struct {
	Host string
}

type ContainerResponse struct {
	items []Container
}

type ExecResponse struct {
	Id string
}

type DockerContainer struct {
	Id      string
	Names   []string
	Image   string
	ImageID string
	Command string
	Created int64
	Status  string
}

func NewDockerClient(slave *MesosSlave) *DockerClient {
	endpoint := fmt.Sprintf("http://%s:%d", slave.Hostname, 2375)
	return &DockerClient{
		Host: endpoint,
	}
}

func (client *DockerClient) ListContainers() []Container {
	resp, err := http.Get(client.Host + "/containers/json")
	if err != nil {
		logrus.Println("get container error")
		return nil
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	items := make([]Container, 0)
	json.Unmarshal(body, &items)
	return items
}

func (client *DockerClient) CreateExec(id string, cmd string) (string, error) {
	var jsonBody = strings.NewReader(`{
		"AttachStdin": true,
		"AttachStdout": true,
		"AttachStderr": true,
		"DetachKeys": "ctrl-p,ctrl-q",
		"Tty": true,
		"Cmd": [
		"` + cmd + `"
		]
	}`)

	res, err := http.Post(client.Host+"/containers/"+id+"/exec", "application/json;charset=utf-8", jsonBody)

	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	var result ExecResponse
	json.Unmarshal(body, &result)
	return result.Id, nil
}

func (client *DockerClient) ExecStart(id string, input chan []byte) (chan []byte, error) {
	execUrl, _ := url.Parse(client.Host + "/exec/" + id + "/start")
	return client.connect(execUrl, input)
}

func (client *DockerClient) ExecResize(id string, width int, height int) error {
	execUrl := fmt.Sprintf(client.Host+"/exec/%s/resize?h=%d&w=%d", id, height, width)

	resp, err := http.Post(execUrl, "application/json;charset=utf-8", nil)

	if err != nil {
		return err
	}

	body, _ := ioutil.ReadAll(resp.Body)

	if len(body) == 0 {
		return nil
	}

	return errors.New(string(body))

}

func (client *DockerClient) connect(url *url.URL, input chan []byte) (chan []byte, error) {
	output := make(chan []byte)

	req, _ := http.NewRequest("POST", url.String(), strings.NewReader(
		`{
			"Detach": false,
			"Tty": true
		}`))
	dial, err := net.Dial("tcp", url.Host)
	if err != nil {
		logrus.Debug(err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	clientconn := httputil.NewClientConn(dial, nil)
	clientconn.Do(req)

	rwc, br := clientconn.Hijack()

	go func() {
		defer clientconn.Close()

		for {
			if data, ok := <-input; ok {
				rwc.Write(data)
			} else {
				break
			}
		}
	}()

	go func() {
		defer rwc.Close()

		for {
			buf := make([]byte, 1024)
			_, err := br.Read(buf)
			if err != nil {
				if err.Error() == "EOF" {
					output <- []byte("EOF")
					break
				}
				logrus.Debug("Read Err: " + err.Error())
				break
			}

			output <- buf

			//Equal EOF
			if buf[0] == 69 && buf[1] == 79 && buf[2] == 70 {
				close(output)
				break
			}

			time.Sleep(500)
			buf = nil
		}

	}()
	return output, nil
}
