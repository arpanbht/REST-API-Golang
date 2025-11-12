package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	fmt.Printf("Server started %s\n", cfg.HTTPServer.Addr)

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		server.ListenAndServe()
	}()

	<-done
	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	fmt.Println("🛑 Shutting down server gracefully...")
	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Failed to shutdown properly 🚨", slog.String("Error", err.Error()))
	}

	fmt.Println("👋 Server stopped cleanly.")

}
