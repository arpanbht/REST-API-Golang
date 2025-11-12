package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/arpanbht/REST-API-Golang/internal/config"
	student "github.com/arpanbht/REST-API-Golang/internal/http/handlers/students"
)

func main() {
	// load configuration
	cfg := config.MustLoad()

	// database setup

	// setup router
	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.New())

	// setup server
	server := http.Server{
		Addr:    cfg.HTTPServer.Addr,
		Handler: router,
	}

	slog.Info("Starting server", slog.String("address", cfg.HTTPServer.Addr))

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		server.ListenAndServe()
	}()

	<-done
	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	slog.Info("🛑 Shutting down server gracefully...")
	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Failed to shutdown properly 🚨", slog.String("Error", err.Error()))
	}

	slog.Info("👋 Server stopped cleanly.")

}
