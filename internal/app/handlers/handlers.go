package handlers

import (
	"compress/gzip"
	"embed"
	"html/template"
	"net/http"

	"github.com/rs/zerolog/log"

	appMiddleware "github.com/belamov/ypgo-metrics/internal/app/handlers/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/belamov/ypgo-metrics/internal/app/services"
)

//go:embed templates/*.gohtml
var templatesDir embed.FS

type Handler struct {
	metricsService services.MetricServiceInterface
	templates      *template.Template
}

func NewHandler(service services.MetricServiceInterface) *Handler {
	templates, err := template.ParseFS(templatesDir, "templates/*.gohtml")
	if err != nil {
		log.Panic().Err(err).Msg("failed parsing templates")
	}
	return &Handler{
		metricsService: service,
		templates:      templates,
	}
}

func NewRouter(service services.MetricServiceInterface) http.Handler {
	h := NewHandler(service)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Compress(gzip.BestSpeed, "application/json", "text/html"))
	r.Use(appMiddleware.GzipDecompressor)
	r.Use(middleware.Heartbeat("/ping"))

	r.Post("/value/", h.GetMetricJSON)
	r.Post("/update/", h.UpdateMetricJSON)
	r.Get("/value/{metricType}/{metricName}", h.GetMetric)
	r.Post("/update/{metricType}/{metricName}/{metricValue}", h.UpdateMetric)
	r.Get("/", h.GetAllMetrics)

	return r
}
