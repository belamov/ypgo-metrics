package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/belamov/ypgo-metrics/internal/app/resources"
	"github.com/rs/zerolog/log"
)

func (h *Handler) UpdateMetricJSON(w http.ResponseWriter, r *http.Request) {
	var req resources.Metric
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&req); err != nil {
		log.Err(err).Msg("cannot decode request JSON body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var updateErr error

	switch req.MType {
	case TypeCounter:
		updateErr = h.metricsService.UpdateCounterMetric(r.Context(), req.ID, *req.Delta)
	case TypeGauge:
		updateErr = h.metricsService.UpdateGaugeMetric(r.Context(), req.ID, *req.Value)
	default:
		http.Error(w, "unknown metric type", http.StatusBadRequest)
		return
	}

	if updateErr != nil {
		log.Err(updateErr).Msg("error updating metric")
		http.Error(w, updateErr.Error(), http.StatusInternalServerError)
		return
	}

	resp := req

	w.Header().Set("Content-Type", "application/json")

	enc := json.NewEncoder(w)
	if err := enc.Encode(resp); err != nil {
		log.Err(err).Msg("error encoding response")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
