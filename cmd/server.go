package main

import (
	"log"
	"net/http"

	"bidhub.com/app"
)

func NewServer() http.Handler {
	mux := http.NewServeMux()
	app.AddRoutes(mux)
	return mux
}

func run() error {
	log.Println("Starting http server...")
	server := NewServer()
	httpServer := &http.Server{
		Addr:    ":6969",
		Handler: server,
	}
	return httpServer.ListenAndServe()
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
