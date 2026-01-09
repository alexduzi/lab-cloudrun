package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/alexduzi/labcloudrun/internal/client"
	"github.com/alexduzi/labcloudrun/internal/config"
	h "github.com/alexduzi/labcloudrun/internal/http"
)

// @title Weather
// @version 1.0
// @description API for fetching weather by zipcode
// @termsOfService http://swagger.io/terms/

// @contact.name Alex Duzi
// @contact.email duzihd@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http https
func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	slog.Info("Configuration loaded", "port", cfg.Port)

	// Initialize clients with config
	cepApiApiClient := client.NewCepClient(cfg)
	weatherApiClient := client.NewWeatherClient(cfg)

	// Initialize HTTP handler
	h := h.NewHttpHandler(cfg, cepApiApiClient, weatherApiClient)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Port),
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
