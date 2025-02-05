package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kerezsiz42/kinko/internal/server"
)

func main() {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: server.Mux(),
	}

	go func() {
		slog.Info("Starting server on port 8080")
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			slog.Error("Error starting server", "error", err)
		}
	}()

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)
	<-signalCh

	slog.Info("Server shutdown started")
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		slog.Error("Error shutting down server", "error", err)
	}

	slog.Info("Server shut down gracefully")
}
