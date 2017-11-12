package mesos


import (
	"fmt"
	. "github.com/JetMuffin/nap/pkg/types"
	"errors"
)

func (c *Client) ListTasks(filter func(t MesosTask) bool) []MesosTask {
	var filteredTasks []MesosTask

	state, err := c.GetStateFromLeader()
	if err != nil {
		return filteredTasks
	}

	// TODO(jetmuffin): list all tasks
	for _, framework := range state.Frameworks {
		for _, task := range framework.Tasks {
			if filter(task) {
				filteredTasks = append(filteredTasks, task)
			}
		}
	}

	return filteredTasks
}

func (c *Client) GetTaskByID(taskId string) (MesosTask, error) {
	tasks := c.ListTasks(func(t MesosTask) bool {
		return t.ID == taskId
	})

	if len(tasks) == 0 {
		return MesosTask{}, errors.New("Task not found.")
	}

	return tasks[0], nil
}

func (c *Client) TaskContainerName(task MesosTask) (string, error) {
	slaveId := task.SlaveID

	for _, status := range task.Statuses {
		if status.State == "TASK_RUNNING" {
			containerName := fmt.Sprintf("mesos-%s.%s", slaveId, status.ContainerStatus.ContainerID.Value)
			return containerName, nil
		}
	}

	return "", errors.New("Task is not running.")
}

func (c *Client) TaskSlave(task MesosTask) (*MesosSlave, error) {
	state, err := c.GetStateFromLeader()
	if err != nil {
		return nil, err
	}

	slave, err := c.GetSlaveByID(state.Slaves, task.SlaveID)
	if err != nil {
		return nil, err
	}

	return slave, nil
}
