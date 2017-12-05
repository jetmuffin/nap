package docker

import (
	"github.com/go-check/check"
	"net/url"
	"testing"
)

func TestDockerClient(t *testing.T) {
	check.TestingT(t)
}

type DockerClientSuite struct {
	client Client
}

var _ = check.Suite(&DockerClientSuite{})

func (s *DockerClientSuite) TestParseHostURL(c *check.C) {
	testcases := []struct {
		host        string
		expected    *url.URL
		expectedErr string
	}{
		{
			host:        "",
			expectedErr: "unable to parse docker host ``",
		},
		{
			host:        "foobar",
			expectedErr: "unable to parse docker host `foobar`",
		},
		{
			host:     "foo://bar",
			expected: &url.URL{Scheme: "foo", Host: "bar"},
		},
		{
			host:     "tcp://localhost:2476",
			expected: &url.URL{Scheme: "tcp", Host: "localhost:2476"},
		},
		{
			host:     "tcp://localhost:2476/path",
			expected: &url.URL{Scheme: "tcp", Host: "localhost:2476", Path: "/path"},
		},
	}

	for _, testcase := range testcases {
		actual, err := ParseHostURL(testcase.host)
		if testcase.expectedErr != "" {
			c.Assert(err, check.ErrorMatches, testcase.expectedErr)
		}
		c.Assert(actual, check.DeepEquals, testcase.expected)
	}
}

func (s *DockerClientSuite) TestGetAPIPath(c *check.C) {
	testcases := []struct {
		version  string
		path     string
		query    url.Values
		expected string
	}{
		{"", "/containers/json", nil, "/containers/json"},
		{"", "/containers/json", url.Values{}, "/containers/json"},
		{"", "/containers/json", url.Values{"s": []string{"c"}}, "/containers/json?s=c"},
		{"1.22", "/containers/json", nil, "/v1.22/containers/json"},
		{"1.22", "/containers/json", url.Values{}, "/v1.22/containers/json"},
		{"1.22", "/containers/json", url.Values{"s": []string{"c"}}, "/v1.22/containers/json?s=c"},
		{"v1.22", "/containers/json", nil, "/v1.22/containers/json"},
		{"v1.22", "/containers/json", url.Values{}, "/v1.22/containers/json"},
		{"v1.22", "/containers/json", url.Values{"s": []string{"c"}}, "/v1.22/containers/json?s=c"},
		{"v1.22", "/networks/kiwl$%^", nil, "/v1.22/networks/kiwl$%25%5E"},
	}

	for _, testcase := range testcases {
		cli := Client{version: testcase.version, basePath: "/"}
		actual := cli.getAPIPath(testcase.path, testcase.query)
		c.Assert(actual, check.Equals, testcase.expected)
	}
}
