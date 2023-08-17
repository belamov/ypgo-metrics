package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/belamov/ypgo-metrics/internal/app/services"
)

type Handler struct {
	metricsService services.MetricServiceInterface
}

func NewHandler(service services.MetricServiceInterface) *Handler {
	return &Handler{
		metricsService: service,
	}
}

func NewRouter(service services.MetricServiceInterface) http.Handler {
	h := NewHandler(service)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Heartbeat("/ping"))

	r.Post("/value/", h.GetMetricJSON)
	r.Post("/update/", h.UpdateMetricJSON)
	r.Get("/value/{metricType}/{metricName}", h.GetMetric)
	r.Post("/update/{metricType}/{metricName}/{metricValue}", h.UpdateMetric)

	return r
}
