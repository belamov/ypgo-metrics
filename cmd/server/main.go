package main

import (
	"github.com/belamov/ypgo-metrics/internal/app/server"
	"github.com/belamov/ypgo-metrics/internal/app/services"
	"github.com/belamov/ypgo-metrics/internal/app/storage"
)

func main() {
	repo := storage.NewInMemoryRepository()
	metricsService := services.NewMetricService(repo)

	serverAddress := "0.0.0.0:8080"

	srv := server.NewHttpServer(serverAddress, metricsService)
	srv.Run()
}
