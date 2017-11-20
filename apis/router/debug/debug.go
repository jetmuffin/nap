package debug

import (
	"github.com/JetMuffin/nap/apis/router"
	"net/http/pprof"
)

// NewRouter initializes a new debug router
func NewRouter() router.Router {
	r := &debugRouter{}
	r.initRoutes()
	return r
}

type debugRouter struct {
	routes []router.Route
}

func (dr *debugRouter) Routes() []router.Route {
	return dr.routes
}

func (dr *debugRouter) initRoutes() {
	dr.routes = []router.Route{
		router.NewGetRoute("/pprof/", pprof.Index),
		router.NewGetRoute("/pprof/cmdline", pprof.Cmdline),
		router.NewGetRoute("/pprof/profile", pprof.Profile),
		router.NewGetRoute("/pprof/symbol", pprof.Symbol),
		router.NewGetRoute("/pprof/trace", pprof.Trace),
	}
}
