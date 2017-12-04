package docker

import (
	"encoding/json"
	"fmt"
	"github.com/JetMuffin/nap/pkg/mesos"
	"github.com/Sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

type DockerClient struct {
	Host string
}

type Container struct {
	Id      string
	Names   []string
	Image   string
	ImageID string
	Command string
	Created int64
	Status  string
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

func NewDockerClient(slave mesos.Slave) *DockerClient {
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
