package auth

import "github.com/JetMuffin/nap/apis/router"

type authRouter struct {
	routes    []router.Route
	oAuthAddr string
}

// NewRouter initializes a new auth router
func NewRouter(addr string) router.Router {
	r := &authRouter{
		oAuthAddr: addr,
	}
	r.initRoutes()
	return r
}

// Routes returns the available routes to the auth controller
func (ar *authRouter) Routes() []router.Route {
	return ar.routes
}

func (ar *authRouter) initRoutes() {
	ar.routes = []router.Route{
		router.NewGetRoute("/auth", ar.handleAuthorize),
	}
}
