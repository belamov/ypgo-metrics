package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/belamov/ypgo-metrics/internal/app/resources"
	"github.com/stretchr/testify/require"

	"github.com/belamov/ypgo-metrics/internal/app/services"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func (s *HandlersTestSuite) TestGetCounterMetricJSON() {
	metricName := "metricName"
	metricValue := int64(10)

	req := resources.Metric{
		ID:    metricName,
		MType: "counter",
		Delta: nil,
		Value: nil,
	}

	body, err := json.Marshal(req)
	require.NoError(s.T(), err)

	s.mockService.EXPECT().GetCounterMetric(gomock.Any(), metricName).Times(1).Return(metricValue, nil)

	result, response := s.testRequest(
		http.MethodPost,
		"/value/",
		string(body),
		nil,
	)
	_ = result.Body.Close()

	assert.Equal(s.T(), http.StatusOK, result.StatusCode)

	expectedResponse := req
	expectedResponse.Delta = &metricValue

	expectedJSON, err := json.Marshal(expectedResponse)
	assert.NoError(s.T(), err)

	assert.JSONEq(s.T(), string(expectedJSON), response)
}

func (s *HandlersTestSuite) TestGetUnknownCounterMetricJSON() {
	metricName := "metricName"

	req := resources.Metric{
		ID:    metricName,
		MType: "counter",
		Delta: nil,
		Value: nil,
	}

	body, err := json.Marshal(req)
	require.NoError(s.T(), err)

	s.mockService.EXPECT().GetCounterMetric(gomock.Any(), metricName).Times(1).Return(int64(0), services.ErrMetricNotFound)

	result, _ := s.testRequest(
		http.MethodPost,
		"/value/",
		string(body),
		nil,
	)
	_ = result.Body.Close()

	assert.Equal(s.T(), http.StatusNotFound, result.StatusCode)
}

func (s *HandlersTestSuite) TestGetGaugeMetricJSON() {
	metricName := "metricName"
	metricValue := float64(10)

	req := resources.Metric{
		ID:    metricName,
		MType: "gauge",
		Delta: nil,
		Value: nil,
	}
	body, err := json.Marshal(req)
	require.NoError(s.T(), err)

	s.mockService.EXPECT().GetGaugeMetric(gomock.Any(), metricName).Times(1).Return(metricValue, nil)

	result, response := s.testRequest(
		http.MethodPost,
		"/value/",
		string(body),
		nil,
	)
	_ = result.Body.Close()

	expectedResponse := req
	expectedResponse.Value = &metricValue

	expectedJSON, err := json.Marshal(expectedResponse)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), http.StatusOK, result.StatusCode)
	assert.JSONEq(s.T(), string(expectedJSON), response)
}

func (s *HandlersTestSuite) TestGetUnknownGaugeMetricJSON() {
	metricName := "metricName"

	req := resources.Metric{
		ID:    metricName,
		MType: "gauge",
		Delta: nil,
		Value: nil,
	}
	body, err := json.Marshal(req)
	require.NoError(s.T(), err)

	s.mockService.EXPECT().GetGaugeMetric(gomock.Any(), metricName).Times(1).Return(float64(0), services.ErrMetricNotFound)

	result, _ := s.testRequest(
		http.MethodPost,
		"/value/",
		string(body),
		nil,
	)
	_ = result.Body.Close()

	assert.Equal(s.T(), http.StatusNotFound, result.StatusCode)
}

func (s *HandlersTestSuite) TestGetUnknownTypeMetricJSON() {
	metricName := "metricName"

	req := resources.Metric{
		ID:    metricName,
		MType: "unknown",
		Delta: nil,
		Value: nil,
	}
	body, err := json.Marshal(req)
	require.NoError(s.T(), err)

	s.mockService.EXPECT().GetGaugeMetric(gomock.Any(), metricName).Times(1).Return(float64(0), services.ErrMetricNotFound)

	result, _ := s.testRequest(
		http.MethodPost,
		"/value/",
		string(body),
		nil,
	)
	_ = result.Body.Close()

	assert.Equal(s.T(), http.StatusBadRequest, result.StatusCode)
}

func (s *HandlersTestSuite) TestGetMetricUnexpectedErrorJSON() {
	metricName := "metricName"

	req := resources.Metric{
		ID:    metricName,
		MType: "gauge",
		Delta: nil,
		Value: nil,
	}
	body, err := json.Marshal(req)
	require.NoError(s.T(), err)

	s.mockService.EXPECT().GetGaugeMetric(gomock.Any(), metricName).Times(1).Return(float64(0), errors.New("unexpected"))

	result, response := s.testRequest(
		http.MethodPost,
		"/value/",
		string(body),
		nil,
	)
	_ = result.Body.Close()

	assert.Equal(s.T(), http.StatusInternalServerError, result.StatusCode)
	assert.Contains(s.T(), response, "unexpected")
}

func (s *HandlersTestSuite) TestGetMetricInvalidJSON() {
	body, err := json.Marshal(`{"name":"metric", "type":"counter"}`)
	require.NoError(s.T(), err)

	result, _ := s.testRequest(
		http.MethodPost,
		"/value/",
		string(body),
		nil,
	)
	_ = result.Body.Close()

	assert.Equal(s.T(), http.StatusBadRequest, result.StatusCode)
}
