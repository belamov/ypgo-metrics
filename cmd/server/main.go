package main

import (
	"os"

	"github.com/belamov/ypgo-metrics/internal/app"

	"github.com/belamov/ypgo-metrics/internal/app/server"
	"github.com/belamov/ypgo-metrics/internal/app/services"
	"github.com/belamov/ypgo-metrics/internal/app/storage"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().Caller().Logger()

	repo := storage.NewInMemoryRepository()
	metricsService := services.NewMetricService(repo)

	config := app.BuildServerConfig()

	srv := server.NewHTTPServer(config.Address, metricsService)
	srv.Run()
}
