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

	"github.com/alexduzi/labcloudrun/internal/client"
	h "github.com/alexduzi/labcloudrun/internal/http"
)

func main() {
	cepApiApiClient := client.NewCepClient()
	weatherApiClient := client.NewWeatherClient()

	h := h.NewHttpHandler("8080", cepApiApiClient, weatherApiClient)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", h.Addr),
		Handler: h.SetupRouter().Handler(),
	}

	go func() {
		slog.Info("server starting at", "addr", srv.Addr)

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("server failed to start", "err", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-stop

	slog.Info("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		slog.Error("server forced to shutdown", "err", err)
	} else {
		slog.Info("server gracefully stopped")
	}
}
