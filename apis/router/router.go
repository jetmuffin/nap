package router

import (
	"net/http"
)

// Router defines an interface to specify a group of routes
type Router interface {
	// Routes returns the list of routes to add to the nap server.
	Routes() []Route
}

// HandlerFunc is a wrapper for http response function.
type HandlerFunc func(w http.ResponseWriter, r *http.Request)

// Route defines an individual API route.
type Route struct {
	method  string
	path    string
	handler HandlerFunc
}

// Method returns the http method that the route responds to.
func (r *Route) Method() string {
	return r.method
}

// Path returns the subpath where the route responds to.
func (r *Route) Path() string {
	return r.path
}

// Handler returns the HandlerFunc to let the server wrap it in middlewares.
func (r *Route) Handler() HandlerFunc {
	return r.handler
}

// NewRoute initializes a new route for the router.
func NewRoute(method, path string, handler HandlerFunc) Route {
	return Route{method, path, handler}
}

// NewGetRoute initializes a new route with the http method GET.
func NewGetRoute(path string, handler HandlerFunc) Route {
	return NewRoute("GET", path, handler)
}

// NewPostRoute initializes a new route with the http method POST.
func NewPostRoute(path string, handler HandlerFunc) Route {
	return NewRoute("POST", path, handler)
}

// NewPutRoute initializes a new route with the http method PUT.
func NewPutRoute(path string, handler HandlerFunc) Route {
	return NewRoute("PUT", path, handler)
}

// NewDeleteRoute initializes a new route with the http method DELETE.
func NewDeleteRoute(path string, handler HandlerFunc) Route {
	return NewRoute("DELETE", path, handler)
}

// NewOptionsRoute initializes a new route with the http method OPTIONS.
func NewOptionsRoute(path string, handler HandlerFunc) Route {
	return NewRoute("OPTIONS", path, handler)
}

// NewHeadRoute initializes a new route with the http method HEAD.
func NewHeadRoute(path string, handler HandlerFunc) Route {
	return NewRoute("HEAD", path, handler)
}
