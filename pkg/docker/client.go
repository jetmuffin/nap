package docker

import (
	"encoding/json"
	"fmt"
	"github.com/JetMuffin/nap/pkg/mesos"
	"github.com/Sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

type Client struct {
	Host string
}

type ContainerResponse struct {
	items []Container
}

type ExecResponse struct {
	ID string
}

type Container struct {
	ID      string
	Names   []string
	Image   string
	ImageID string
	Command string
	Created int64
	Status  string
}

func NewClient(slave mesos.Slave) *Client {
	endpoint := fmt.Sprintf("http://%s:%d", slave.Hostname, 2375)
	return &Client{
		Host: endpoint,
	}
}

func (client *Client) ListContainers() []Container {
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
