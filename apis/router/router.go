package router

import (
	"net/http"
)

type Router interface {
	// Routes returns the list of routes to add to the nap server.
	Routes() []Route
}

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

type Route struct {
	method  string
	path    string
	handler HandlerFunc
}

func (r *Route) Method() string {
	return r.method
}

func (r *Route) Path() string {
	return r.path
}

func (r *Route) Handler() HandlerFunc {
	return r.handler
}

func NewRoute(method, path string, handler HandlerFunc) Route {
	var r Route = Route{method, path, handler}
	return r
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
