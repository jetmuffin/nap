package docker

import (
	"github.com/go-check/check"
	"net/http"
	"fmt"
	"strings"
	"bytes"
	"encoding/json"
	"io/ioutil"
)

func (s *DockerClientSuite) TestExecCreate(c *check.C) {
	expectedURL := "/containers/container_id/exec"
	client := &Client{
		client: newMockClient(func(req *http.Request) (*http.Response, error) {
			if !strings.HasPrefix(req.URL.Path, expectedURL) {
				return nil, fmt.Errorf("expected URL '%s', got '%s'", expectedURL, req.URL)
			}
			if req.Method != "POST" {
				return nil, fmt.Errorf("expected POST method, got %s", req.Method)
			}
			// FIXME validate the content is the given ExecConfig ?
			if err := req.ParseForm(); err != nil {
				return nil, err
			}
			execConfig := &ExecConfig{}
			if err := json.NewDecoder(req.Body).Decode(execConfig); err != nil {
				return nil, err
			}
			if execConfig.User != "root" {
				return nil, fmt.Errorf("expected an execConfig with User == 'root', got %v", execConfig.User)
			}
			b, err := json.Marshal(ExecResponse{
				ID: "exec_id",
			})
			if err != nil {
				return nil, err
			}
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewReader(b)),
			}, nil
		}),
	}

	r, err := client.ExecCreate("container_id", "/bin/bash")
	c.Assert(err, check.IsNil)
	c.Assert(r, check.Equals, "exec_id")
}
