package mesos

import "github.com/go-check/check"

func (s *MesosClientSuite) TestListSlaves(c *check.C) {
	slaves, err := s.client.ListSlaves()

	c.Assert(err, check.IsNil)
	c.Assert(slaves, check.HasLen, 4)
}

func (s *MesosClientSuite) TestGetSlaveByID(c *check.C) {
	slave, err := s.client.GetSlaveByID("348a8c33-f98d-4ba2-bd4b-03b278332a65-S4")

	c.Assert(err, check.IsNil)
	c.Assert(slave.Hostname, check.Equals, "n170")
}

func (s *MesosClientSuite) TestGetSlaveByFakeID(c *check.C) {
	slave, err := s.client.GetSlaveByID("fake_id")

	c.Assert(err, check.ErrorMatches, `slave ID .* not found`)
	c.Assert(slave.ID, check.Equals, "")
}
