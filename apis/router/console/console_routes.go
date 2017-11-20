package console

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/JetMuffin/nap/apis/utils"
	"github.com/JetMuffin/nap/pkg/docker"
	"github.com/Sirupsen/logrus"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
)

func (cr *consoleRouter) hello(w http.ResponseWriter, req *http.Request) {
	utils.WriteJSON(w, http.StatusOK, "hello")
}

func (cr *consoleRouter) handleTaskConsole(w http.ResponseWriter, req *http.Request) {
	// TODO(jetmuffin: add origin to white list
	req.Header.Del("Origin")

	c, err := cr.upgrader.Upgrade(w, req, nil)
	if err != nil {
		logrus.Error("Upgrader error: ", err)
		return
	}
	defer c.Close()

	// Parse request parameters
	if err := utils.ParseForm(req); err != nil {
		c.WriteMessage(1, cr.handleError(err))
		return
	}

	// Get task
	taskID := req.Form.Get("task_id")
	task, err := cr.backend.GetTaskByID(taskID)
	if err != nil {
		c.WriteMessage(1, cr.handleError(errors.New("task not found")))
	}

	// Get container id.
	containerID, err := cr.backend.TaskContainerName(task)
	if err != nil {
		c.WriteMessage(1, cr.handleError(errors.New("task container not found, please check the task_id parameter")))
	}

	// Get slave node
	slave, err := cr.backend.TaskSlave(task)
	if err != nil {
		c.WriteMessage(1, cr.handleError(errors.New("task is not running on any slave now")))
	}

	cli := docker.NewDockerClient(slave)

	var lock sync.Mutex

	input := make(chan []byte)
	execID, err := cli.CreateExec(containerID, "/bin/bash")
	if err != nil {
		logrus.Error(err)
	}
	output, err := cli.ExecStart(execID, input)
	if err != nil {
		logrus.Error(err)
	}

	for {
		go func() {
			for {
				if data, ok := <-output; ok {
					lock.Lock()
					c.WriteMessage(websocket.TextMessage, cr.handleResponse(data))
					lock.Unlock()
				} else {
					break
				}
			}
		}()

		mt, rawMessage, err := c.ReadMessage()
		if err != nil {
			logrus.Error("Read error: ", err)
			cr.handleClose(input)
			break
		}

		var message Message
		json.Unmarshal([]byte(rawMessage), &message)

		switch message.Type {
		case PingMessage:
			c.WriteMessage(mt, cr.handlePing())
		case InitMessage:
			logrus.Infof("Console of task %s inited.", taskID)
		case InputMessage:
			cr.handleInput(input, message.Content)
		case ResizeMessage:
			cr.handleResize(cli, execID, message.Content)
		}
	}
}

func (cr *consoleRouter) handlePing() []byte {
	resp := Message{
		Type: PongMessage,
	}
	b, _ := json.Marshal(resp)
	return b
}

func (cr *consoleRouter) handleInput(input chan []byte, cmd string) {
	input <- []byte(cmd)
}

func (cr *consoleRouter) handleResize(cli *docker.DockerClient, execID string, content string) {
	var resize ResizeContent
	json.Unmarshal([]byte(content), &resize)

	cli.ExecResize(execID, resize.Columns, resize.Rows)
}

func (cr *consoleRouter) handleError(err error) []byte {
	resp := Message{
		Type:    ErrorMessage,
		Content: fmt.Sprintf("Failed to launch container terminal: %v", err),
	}
	b, _ := json.Marshal(resp)
	return b
}

func (cr *consoleRouter) handleResponse(data []byte) []byte {
	resp := Message{
		Type:    OutputMessage,
		Content: base64.StdEncoding.EncodeToString(data),
	}
	b, _ := json.Marshal(resp)
	return b
}

func (cr *consoleRouter) handleClose(input chan []byte) {
	input <- []byte("EOF")
	close(input)
}
