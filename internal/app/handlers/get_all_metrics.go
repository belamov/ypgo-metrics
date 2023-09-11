package handlers

import (
	"net/http"

	"github.com/belamov/ypgo-metrics/internal/app/models"
)

type AllMetricsData struct {
	CounterMetrics []models.CounterMetric
	GaugeMetrics   []models.GaugeMetric
}

func (h *Handler) GetAllMetrics(w http.ResponseWriter, r *http.Request) {
	counterMetrics, gaugeMetrics, err := h.metricsService.GetAllMetrics(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := AllMetricsData{
		CounterMetrics: counterMetrics,
		GaugeMetrics:   gaugeMetrics,
	}

	w.Header().Set("Content-Type", "text/html")

	err = h.templates.ExecuteTemplate(w, "all_metrics.gohtml", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
