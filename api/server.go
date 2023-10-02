package api

import (
	"net/http"

	"github.com/prests/hypermedia-golang/router"
)

type server struct {
	http.Handler
}

func NewServer(fs http.Handler, routers ...router.Router) *server {
	s := new(server)

	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	for _, apiRouter := range routers {
		for _, r := range apiRouter.Routes() {
			mux.HandleFunc(r.Pattern(), r.Handler())
		}
	}

	s.Handler = mux

	return s
}
