package main

import (
	"github.com/belamov/ypgo-metrics/internal/app/server"
	"github.com/belamov/ypgo-metrics/internal/app/services"
	"github.com/belamov/ypgo-metrics/internal/app/storage"
)

func main() {
	repo := storage.NewInMemoryRepository()
	metricsService := services.NewMetricService(repo)

	srv, err := server.NewHttpServer(metricsService)
	if err != nil {
		panic(err)
	}

	srv.Run()
}
