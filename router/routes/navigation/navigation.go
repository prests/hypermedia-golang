package navigation

import "github.com/prests/hypermedia-golang/router"

type navigationRouter struct {
	routes []router.Route

	templateBasePath string
}

func NewRouter() router.Router {
	nr := &navigationRouter{
		templateBasePath: "",
	}

	nr.initRoutes()
	return nr
}

func (nr *navigationRouter) Routes() []router.Route {
	return nr.routes
}

func (nr *navigationRouter) initRoutes() {
	nr.routes = []router.Route{
		router.NewRoute("/", nr.landingPage),
	}
}
