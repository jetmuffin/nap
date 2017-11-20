package console

import (
	"net/http"
	"github.com/JetMuffin/nap/pkg/api/utils"
	"errors"
	"github.com/JetMuffin/nap/pkg/docker"
	"github.com/Sirupsen/logrus"
	"encoding/json"
	"fmt"
	"encoding/base64"
	"github.com/gorilla/websocket"
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
	taskId := req.Form.Get("task_id")
	task, err := cr.backend.GetTaskByID(taskId)
	if err != nil {
		c.WriteMessage(1, cr.handleError(errors.New("Task not found.")))
	}

	// Get container id.
	containerId, err := cr.backend.TaskContainerName(task)
	if err != nil {
		c.WriteMessage(1, cr.handleError(errors.New("Task container not found, please check the task_id parameter.")))
	}

	// Get slave node
	slave, err := cr.backend.TaskSlave(task)
	if err != nil {
		c.WriteMessage(1, cr.handleError(errors.New("Task is not running on any slave now.")))
	}

	cli := docker.NewDockerClient(slave)

	var lock sync.Mutex

	input := make(chan []byte)
	execId, err := cli.CreateExec(containerId, "/bin/bash")
	if err != nil {
		logrus.Error(err)
	}
	output, err := cli.ExecStart(execId, input)
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

		var message ConsoleMessage
		json.Unmarshal([]byte(rawMessage), &message)

		switch message.Type {
		case PingMessage:
			c.WriteMessage(mt, cr.handlePing())
		case InitMessage:
			logrus.Infof("Console of task %s inited.", taskId)
		case InputMessage:
			cr.handleInput(input, message.Content)
		case ResizeMessage:
			cr.handleResize(cli, execId, message.Content)
		}
	}
}

func (cr *consoleRouter) handlePing() []byte {
	resp := ConsoleMessage{
		Type: PongMessage,
	}
	b, _ := json.Marshal(resp)
	return b
}

func (cr *consoleRouter) handleInput(input chan []byte, cmd string) {
	input <- []byte(cmd)
}

func (cr *consoleRouter) handleResize(cli *docker.DockerClient, execId string, content string) {
	var resize ResizeContent
	json.Unmarshal([]byte(content), &resize)

	cli.ExecResize(execId, resize.Columns, resize.Rows)
}

func (cr *consoleRouter) handleError(err error) []byte {
	resp := ConsoleMessage{
		Type:    ErrorMessage,
		Content: fmt.Sprintf("Failed to launch container terminal: %v", err),
	}
	b, _ := json.Marshal(resp)
	return b
}

func (cr *consoleRouter) handleResponse(data []byte) []byte {
	resp := ConsoleMessage{
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
