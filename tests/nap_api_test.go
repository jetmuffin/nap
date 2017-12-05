package tests

import (
	"fmt"
	"github.com/JetMuffin/nap/apis"
	"github.com/JetMuffin/nap/apis/router"
	"github.com/go-check/check"
	"io/ioutil"
	"net"
	"net/http"
	"testing"
)

type NapAPISuite struct {
	listenAddr string
	conn       net.Listener
	endpoint   string

	api *apis.Server
}

var _ = check.Suite(&NapAPISuite{})

func TestNapAPI(t *testing.T) {
	check.TestingT(t)
}

func initRouter() []router.Router {
	routers := []router.Router{}

	return routers
}

func (s *NapAPISuite) SetUpSuite(c *check.C) {
	s.listenAddr = "0.0.0.0:5678"
	conn, err := net.Listen("tcp", s.listenAddr)
	s.conn = conn
	s.endpoint = "http://localhost:5678"
	c.Assert(err, check.IsNil)
	c.Assert(s.conn, check.NotNil)

	apiConfig := &apis.Config{
		LogLevel: "debug",
	}
	s.api = apis.New(apiConfig)

	s.api.Accept(s.listenAddr, s.conn)
	s.api.InitRouter(initRouter()...)

	serveAPIWait := make(chan error)
	go s.api.Wait(serveAPIWait)
}

func (s *NapAPISuite) TearDownSuite(c *check.C) {
	s.api.Shutdown()
	s.conn.Close()
}

func (s *NapAPISuite) TestConsoleHello(c *check.C) {
	resp, err := http.Get(fmt.Sprintf("%s/console/hello", s.endpoint))
	c.Assert(err, check.IsNil)

	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)
	c.Assert(err, check.IsNil)
}
