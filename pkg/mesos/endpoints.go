package mesos

import (
	"fmt"
)

func (c *Client) masterStateEndpoint() string {
	return fmt.Sprintf("%s/master/state", c.Master.String())
}

func (c *Client) slavesEndpoint() string {
	return fmt.Sprintf("%s/slaves", c.Master.String())
}

func (c *Client) tasksEndpoint() string {
	return fmt.Sprintf("%s/tasks", c.Master.String())
}

func (c *Client) metricsEndpoint() string {
	return fmt.Sprintf("%s/metrics/snapshot", c.Master.String())
}

func (c *Client) frameworksEndpoint() string {
	return fmt.Sprintf("%s/frameworks", c.Master.String())
}
