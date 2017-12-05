package docker

import (
	"encoding/json"
	"net/url"
	"strconv"
)

// ExecResponse represents response of exec api.
type ExecResponse struct {
	ID string
}

// ExecConfig isi the Config struct that holds the configuration
// for the exec feature of docker.
type ExecConfig struct {
	User         string
	Privileged   bool
	Tty          bool
	AttachStdin  bool
	AttachStderr bool
	AttachStdout bool
	Detach       bool
	DetachKeys   string
	Env          []string
	Cmd          []string
}

func (cli *Client) defaultExecConfig() ExecConfig {
	return ExecConfig{
		User:         "root",
		AttachStdin:  true,
		AttachStdout: true,
		DetachKeys:   "ctrl-p,ctrl-q",
		Tty:          true,
	}
}

// ExecCreate creates a new exec configuration to run an exec process.
func (cli *Client) ExecCreate(container string, cmd string) (string, error) {
	var response ExecResponse

	execConfig := cli.defaultExecConfig()
	execConfig.Cmd = append(execConfig.Cmd, cmd)

	resp, err := cli.post("/containers/"+container+"/exec", nil, execConfig, nil)
	if err != nil {
		return response.ID, err
	}
	err = json.NewDecoder(resp.body).Decode(&response)
	return response.ID, err
}

// ExecStart starts an exec process already created in the docker host and
// returns a channel to communicate with the process.
func (cli *Client) ExecStart(execID string, input chan []byte) (chan []byte, error) {
	return cli.postHijack("/exec/"+execID+"/start", nil, cli.defaultExecConfig(), nil, input)
}

// ExecResize changes the size of the tty for an exec process running inside a container.
func (cli *Client) ExecResize(execID string, width int, height int) error {
	query := url.Values{}
	query.Set("h", strconv.Itoa(int(height)))
	query.Set("w", strconv.Itoa(int(width)))

	_, err := cli.post("/exec/"+execID+"/resize", query, nil, nil)
	return err
}
