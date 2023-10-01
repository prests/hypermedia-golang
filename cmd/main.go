package main

import (
	"os"

	"github.com/prests/hypermedia-golang/api"
)

func main() {
	server := api.NewServer()
	server.StartServer(os.Args[1])
}