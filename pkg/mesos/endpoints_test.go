package mesos

import (
	"fmt"
	"github.com/go-check/check"
)

func (s *MesosClientSuite) TestMasterStateEndpoint(c *check.C) {
	c.Assert(s.client.masterStateEndpoint(), check.Equals, fmt.Sprintf("%s/master/state", s.addr))
}

func (s *MesosClientSuite) TestSlavesEndpoint(c *check.C) {
	c.Assert(s.client.slavesEndpoint(), check.Equals, fmt.Sprintf("%s/slaves", s.addr))
}

func (s *MesosClientSuite) TestTasksEndpoint(c *check.C) {
	c.Assert(s.client.tasksEndpoint(), check.Equals, fmt.Sprintf("%s/tasks", s.addr))
}

func (s *MesosClientSuite) TestFrameworksEndpoint(c *check.C) {
	c.Assert(s.client.frameworksEndpoint(), check.Equals, fmt.Sprintf("%s/frameworks", s.addr))
}

func (s *MesosClientSuite) TestMetricsEndpoint(c *check.C) {
	c.Assert(s.client.metricsEndpoint(), check.Equals, fmt.Sprintf("%s/metrics/snapshot", s.addr))
}
