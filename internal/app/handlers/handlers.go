package handlers

import (
	"net/http"

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

	mux := http.NewServeMux()
	mux.HandleFunc("/update/gauge/", h.AddGaugeMetric)
	mux.HandleFunc("/update/counter/", h.AddCounterMetric)
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusBadRequest)
	})

	return mux
}
