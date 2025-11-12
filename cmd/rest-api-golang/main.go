package main

import (
	"fmt"
	"net/http"

	"github.com/arpanbht/REST-API-Golang/internal/config"
)

func main() {
	// load configuration
	cfg := config.MustLoad()

	// database setup

	// setup router
	router := http.NewServeMux()

	router.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to students API"))
		w.WriteHeader(http.StatusOK)

	})

	// setup server
	fmt.Printf("Server started %s\n", cfg.HTTPServer.Addr)
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	server.ListenAndServe()

}
