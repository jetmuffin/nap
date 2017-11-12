package console

import (
	"github.com/JetMuffin/nap/pkg/api/router"
	"github.com/gorilla/websocket"
)

type consoleRouter struct {
	backend  Backend
	exec     map[string]string
	inputChan map[string] chan []byte
	outputChan map[string] chan []byte
	upgrader websocket.Upgrader
	routes   []router.Route
}

func NewRouter(b Backend) router.Router {
	r := &consoleRouter{
		backend:  b,
		upgrader: websocket.Upgrader{},
		exec:     make(map[string]string),
		inputChan: make(map[string] chan []byte),
		outputChan: make(map[string] chan []byte),
	}
	r.initRoutes()
	return r
}

func (cr *consoleRouter) Routes() []router.Route {
	return cr.routes
}

func (cr *consoleRouter) initRoutes() {
	cr.routes = []router.Route{
		router.NewGetRoute("/console/hello", cr.hello),
		router.NewGetRoute("/console/ws", cr.handleTaskConsole),
	}
}
