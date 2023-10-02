package router

import "net/http"

type Router interface {
	Routes() []Route
}

type RouteHandler func (http.ResponseWriter, *http.Request)

type Route interface {
	Handler() RouteHandler
	Pattern() string
}

type route struct {
	pattern string
	handler RouteHandler
}

func NewRoute(pattern string, handler RouteHandler) Route {
	return &route{
		pattern,
		handler,
	}
}

func (r *route) Handler() RouteHandler {
	return r.handler
}

func (r *route) Pattern() string {
	return r.pattern
}