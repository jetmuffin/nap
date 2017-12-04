package tests

import (
	"github.com/go-check/check"
	"net"
	"github.com/JetMuffin/nap/apis"
	"github.com/JetMuffin/nap/apis/router"
)

type NapAPISuite struct {
	listenAddr string
	conn       net.Listener

	api *apis.Server
}

func init() {
	check.Suite(&NapAPISuite{})
}

func initRouter() []router.Router{
	routers := []router.Router{}

	return routers
}

func (s *NapAPISuite) SetUpTest(c *check.C) {
	s.listenAddr = "5678"
	conn, err := net.Listen("tcp", s.listenAddr)
	c.Assert(err, check.IsNil)
	s.conn = conn

	s.api.Accept(s.listenAddr, s.conn)
	s.api.InitRouter(initRouter()...)

	serveAPIWait := make(chan error)
	go s.api.Wait(serveAPIWait)
}

func (s *NapAPISuite) TearDownSuite(c *check.C) {
	s.conn.Close()
	s.api.Shutdown()
}
