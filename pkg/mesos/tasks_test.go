package mesos

import "github.com/go-check/check"

func (s *MesosClientSuite) TestListTasks(c *check.C) {
	tasks, err := s.client.ListTasks()

	c.Assert(err, check.IsNil)
	c.Assert(tasks, check.HasLen, 1)
}

func (s *MesosClientSuite) TestGetTaskByID(c *check.C) {
	task, err := s.client.GetTaskByID("simple-webserver.857d518f-241e-11e6-a10b-02424e0f635b")

	c.Assert(err, check.IsNil)
	c.Assert(task.Name, check.Equals, "simple-webserver")
}

func (s *MesosClientSuite) TestGetTaskByFakeID(c *check.C) {
	task, err := s.client.GetTaskByID("fake_id")

	c.Assert(err, check.ErrorMatches, `task ID .* not found`)
	c.Assert(task.ID, check.Equals, "")
}

func (s *MesosClientSuite) TestGetTaskContainerName(c *check.C) {
	task, err := s.client.GetTaskByID("simple-webserver.857d518f-241e-11e6-a10b-02424e0f635b")
	c.Assert(err, check.IsNil)

	container, err := s.client.GetTaskContainerName(task)
	c.Assert(err, check.IsNil)
	c.Assert(container, check.Equals, "mesos-69c30f85-3c9b-4cb3-8cfd-b7ff5ec68e74-S0.9fdf6891-90c9-4ace-8b9b-754d69bf68a8")
}
