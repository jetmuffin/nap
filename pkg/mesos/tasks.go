package mesos

import (
	"encoding/json"
	"fmt"
)

// TasksWrap is the wrapper for task endpoint of mesos.
type TasksWrap struct {
	List []Task `json:"tasks"`
}

// ListTasks returns list of tasks, as tasks endpoint `/tasks` does not have
// enough information, we collects task list from state endpoint.
func (c *Client) ListTasks() ([]Task, error) {
	body, err := c.doRequest(c.masterStateEndpoint())
	if err != nil {
		return []Task{}, err
	}

	var state State
	err = json.Unmarshal(body, &state)
	if err != nil {
		return []Task{}, err
	}

	var tasks []Task
	for _, framework := range state.Frameworks {
		for _, task := range framework.Tasks {
			tasks = append(tasks, task)
		}
	}
	return tasks, nil
}

// GetTaskByID returns specific task with given ID.
func (c *Client) GetTaskByID(ID string) (Task, error) {
	tasks, err := c.ListTasks()
	if err != nil {
		return Task{}, err
	}

	for _, task := range tasks {
		if task.ID == ID {
			return task, nil
		}
	}

	return Task{}, fmt.Errorf("task ID %s not found", ID)
}

// GetTaskContainerName returns container name of given task.
func (c *Client) GetTaskContainerName(task Task) (string, error) {
	slaveID := task.SlaveID

	for _, status := range task.Statuses {
		if status.State == "TASK_RUNNING" {
			containerName := fmt.Sprintf("mesos-%s.%s", slaveID, status.ContainerStatus.ContainerID.Value)
			return containerName, nil
		}
	}

	return "", fmt.Errorf("task %s is not running", task.ID)
}