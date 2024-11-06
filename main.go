package main

import (
	"log"
	"log/slog"

	"RouteHub.Service.Dashboard/features"
	"RouteHub.Service.Dashboard/features/configuration"
	"RouteHub.Service.Dashboard/web"
	_ "github.com/lib/pq"
	"github.com/ory/graceful"
)

const (
	appName = "RouteHub.Service.Dashboard"
)

func main() {
	features.NewLoggerConfigurer(slog.LevelDebug)
	logger := features.GetLogger()

	err := configuration.Configure(configuration.WithDefaultServerConfig(), configuration.WithLogger(logger))
	if err != nil {
		logger.Error("Failed to configure server", "error", err)
		return
	}

	config := configuration.Get()
	app, err := web.NewApplication(config, logger)
	if err != nil {
		log.Fatalf("Failed to create application: %v", err)
	}

	server := graceful.WithDefaults(app.Echo.Server)

	logger.Info("main: Starting server")
	if err := graceful.Graceful(server.ListenAndServe, server.Shutdown); err != nil {
		logger.Error("main: Failed to gracefully shutdown", slog.Any("error", err))
	}
	logger.Info("main: Server stopped")
}
