package mesos

import (
	"fmt"
	"github.com/go-check/check"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestMesosClient(t *testing.T) {
	check.TestingT(t)
}

type MesosClientSuite struct {
	client *Client

	mux    *http.ServeMux
	server *httptest.Server
	addr   *url.URL
}

var _ = check.Suite(&MesosClientSuite{})

func mockEndpoint(mux *http.ServeMux, endpoint string, filePath string) {
	mux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		c, _ := ioutil.ReadFile(filePath)
		fmt.Fprintf(w, string(c))
	})
}

func (s *MesosClientSuite) muxEndpoints() {
	mockEndpoint(s.mux, "/master/state", "../../tests/mock/master.state.json")
	mockEndpoint(s.mux, "/tasks", "../../tests/mock/tasks.json")
	mockEndpoint(s.mux, "/frameworks", "../../tests/mock/frameworks.json")
	mockEndpoint(s.mux, "/slaves", "../../tests/mock/slaves.json")
}

func (s *MesosClientSuite) SetUpSuite(c *check.C) {
	s.mux = http.NewServeMux()
	s.muxEndpoints()

	s.server = httptest.NewServer(s.mux)
	s.addr, _ = url.Parse(s.server.URL)
	s.client = NewClient(s.addr, nil)
}

func (s *MesosClientSuite) TearDownSuite(c *check.C) {
	s.server.Close()
}
