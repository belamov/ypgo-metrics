package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

const (
	TypeGauge   string = "gauge"
	TypeCounter string = "counter"
)

func (h *Handler) UpdateMetric(w http.ResponseWriter, r *http.Request) {
	metricType := chi.URLParam(r, "metricType")

	switch metricType {
	case TypeCounter:
		h.updateCounterMetric(r.Context(), w, r)
	case TypeGauge:
		h.updateGaugeMetric(r.Context(), w, r)
	default:
		http.Error(w, "unknown metric type", http.StatusBadRequest)
		return
	}
}

func (h *Handler) updateGaugeMetric(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	metricName := chi.URLParam(r, "metricName")
	metricValue, err := strconv.ParseFloat(chi.URLParam(r, "metricValue"), 64)
	if err != nil {
		http.Error(w, "metric value must be float", http.StatusBadRequest)
		return
	}

	err = h.metricsService.UpdateGaugeMetric(ctx, metricName, metricValue)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) updateCounterMetric(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	metricName := chi.URLParam(r, "metricName")
	metricValue, err := strconv.ParseInt(chi.URLParam(r, "metricValue"), 10, 64)
	if err != nil {
		http.Error(w, "metric value must be int", http.StatusBadRequest)
		return
	}

	err = h.metricsService.UpdateCounterMetric(ctx, metricName, metricValue)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
