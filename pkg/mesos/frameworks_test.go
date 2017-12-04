package mesos

import (
	"github.com/go-check/check"
)

func (s *MesosClientSuite) TestListFrameworks(c *check.C) {
	frameworks, err := s.client.ListFrameworks()

	c.Assert(err, check.IsNil)
	c.Assert(frameworks, check.HasLen, 1)
}

func (s *MesosClientSuite) TestGetFrameworkByID(c *check.C) {
	framework, err := s.client.GetFrameworkByID("8c527af7-e41b-4c59-a93c-2c5fee263df8-0000")

	c.Assert(err, check.IsNil)
	c.Assert(framework.Name, check.Equals, "chronos")
}

func (s *MesosClientSuite) TestGetFrameworkByFakeID(c *check.C) {
	framework, err := s.client.GetFrameworkByID("fake_id")

	c.Assert(err, check.ErrorMatches, `framework ID .* not found`)
	c.Assert(framework.ID, check.Equals, "")
}
