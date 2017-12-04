package mesos

import "github.com/JetMuffin/nap/apis/router"

type mesosRouter struct {
	routes []router.Route
}

// NewRouter initializes a new mesos router
func NewRouter() router.Router {
	r := &mesosRouter{}
	r.initRoutes()
	return r
}

// Routes returns the available routes to the mesos controller
func (mr *mesosRouter) Routes() []router.Route {
	return mr.routes
}

func (mr *mesosRouter) initRoutes() {
	mr.routes = []router.Route{
		router.NewGetRoute("/state", mr.handleUsage),
	}
}
