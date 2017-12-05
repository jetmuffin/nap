package mesos

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
)

// Client manages the communication with the Mesos cluster.
type Client struct {
	sync.Mutex

	Master *url.URL
	HTTP   *http.Client
}

// Pid is the process if per machine.
type Pid struct {
	// Role of a PID
	Role string
	// Host / IP of the PID
	Host string
	// Port of the PID.
	// If no Port is available the standard port (5050) will be used.
	Port int
}

// NewClient returns a new  Mesos information client.
// addresses has to be the the URL`s of the single nodes of the
// Mesos cluster. It is recommended to apply all nodes in case of failures.
func NewClient(master *url.URL, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	client := &Client{
		Master: master,
		HTTP:   httpClient,
	}
	return client
}

func (c *Client) doRequest(endpoint string) ([]byte, error) {
	resp, err := c.HTTP.Get(endpoint)

	if err != nil {
		return []byte{}, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	return body, err
}
