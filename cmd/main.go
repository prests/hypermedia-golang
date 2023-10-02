package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/prests/hypermedia-golang/api"
	"github.com/prests/hypermedia-golang/router/routes/navigation"
)

func main() {
	fs := http.FileServer(http.Dir("web/assets"))

	navigationRouter := navigation.NewRouter()

	server := api.NewServer(fs, navigationRouter)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Args[1]), server))
}