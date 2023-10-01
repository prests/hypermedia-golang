package api

import (
	"fmt"
	"net/http"
)

type Server interface {
	StartServer(port string) error
}

type server struct {}

func NewServer() Server {
	return &server{}
}

func (s *server) StartServer(port string) error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})

	return http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
