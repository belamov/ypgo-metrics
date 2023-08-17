package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/belamov/ypgo-metrics/internal/app/resources"
	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func (s *HandlersTestSuite) TestCounterMetricAddJSON() {
	metricName := "metricName"
	metricValue := int64(10)

	s.mockService.EXPECT().UpdateCounterMetric(gomock.Any(), metricName, metricValue).Times(1)

	req := resources.Metric{
		ID:    metricName,
		MType: "counter",
		Delta: &metricValue,
		Value: nil,
	}
	body, err := json.Marshal(req)
	require.NoError(s.T(), err)

	result, response := s.testRequest(
		http.MethodPost,
		"/update/",
		string(body),
		nil,
	)
	_ = result.Body.Close()

	assert.Equal(s.T(), http.StatusOK, result.StatusCode)
	assert.JSONEq(s.T(), string(body), response)
}

func (s *HandlersTestSuite) TestCounterMetricAddServiceErrorJSON() {
	metricName := "metricName"
	metricValue := int64(10)

	s.mockService.EXPECT().UpdateCounterMetric(gomock.Any(), metricName, metricValue).
		Times(1).
		Return(errors.New("unexpected error"))

	req := resources.Metric{
		ID:    metricName,
		MType: "counter",
		Delta: &metricValue,
		Value: nil,
	}
	body, err := json.Marshal(req)
	require.NoError(s.T(), err)

	result, _ := s.testRequest(
		http.MethodPost,
		"/update/",
		string(body),
		nil,
	)
	_ = result.Body.Close()

	assert.Equal(s.T(), http.StatusInternalServerError, result.StatusCode)
}

func (s *HandlersTestSuite) TestCounterMetricAddNotFloatJSON() {
	body, err := json.Marshal(`{"id":"metric", "type":"counter", "delta":"not float"}`)
	require.NoError(s.T(), err)

	result, _ := s.testRequest(
		http.MethodPost,
		"/update/",
		string(body),
		nil,
	)
	_ = result.Body.Close()

	assert.Equal(s.T(), http.StatusBadRequest, result.StatusCode)
}

func (s *HandlersTestSuite) TestCounterMetricWrongFormatJSON() {
	body, err := json.Marshal(`{"id":"metric", "delta":1}`)
	require.NoError(s.T(), err)

	result, _ := s.testRequest(
		http.MethodPost,
		"/update/",
		string(body),
		nil,
	)
	_ = result.Body.Close()

	assert.Equal(s.T(), http.StatusBadRequest, result.StatusCode)
}

func (s *HandlersTestSuite) TestGaugeMetricAddJSON() {
	metricName := "metricName"
	metricValue := float64(10)

	s.mockService.EXPECT().UpdateGaugeMetric(gomock.Any(), metricName, metricValue).Times(1)

	req := resources.Metric{
		ID:    metricName,
		MType: "gauge",
		Delta: nil,
		Value: &metricValue,
	}
	body, err := json.Marshal(req)
	require.NoError(s.T(), err)

	result, response := s.testRequest(
		http.MethodPost,
		"/update/",
		string(body),
		nil,
	)
	_ = result.Body.Close()

	assert.Equal(s.T(), http.StatusOK, result.StatusCode)
	assert.JSONEq(s.T(), string(body), response)
}

func (s *HandlersTestSuite) TestGaugeMetricAddServiceErrorJSON() {
	metricName := "metricName"
	metricValue := float64(10)

	s.mockService.EXPECT().UpdateGaugeMetric(gomock.Any(), metricName, metricValue).
		Times(1).
		Return(errors.New("unexpected error"))

	req := resources.Metric{
		ID:    metricName,
		MType: "counter",
		Delta: nil,
		Value: &metricValue,
	}
	body, err := json.Marshal(req)
	require.NoError(s.T(), err)

	result, _ := s.testRequest(
		http.MethodPost,
		"/update/",
		string(body),
		nil,
	)
	_ = result.Body.Close()

	assert.Equal(s.T(), http.StatusInternalServerError, result.StatusCode)
}

func (s *HandlersTestSuite) TestGaugeMetricAddNotFloatJSON() {
	body, err := json.Marshal(`{"id":"metric", "type":"gauge", "value":"not float"}`)
	require.NoError(s.T(), err)

	result, _ := s.testRequest(
		http.MethodPost,
		"/update/",
		string(body),
		nil,
	)
	_ = result.Body.Close()

	assert.Equal(s.T(), http.StatusBadRequest, result.StatusCode)
}

func (s *HandlersTestSuite) TestGaugeMetricWrongFormatJSON() {
	body, err := json.Marshal(`{"id":"metric", "value":1}`)
	require.NoError(s.T(), err)

	result, _ := s.testRequest(
		http.MethodPost,
		"/update/",
		string(body),
		nil,
	)
	_ = result.Body.Close()

	assert.Equal(s.T(), http.StatusBadRequest, result.StatusCode)
}

func (s *HandlersTestSuite) TestUnknownMetricTypeJSON() {
	metricName := "metricName"
	metricValue := int64(10)

	req := resources.Metric{
		ID:    metricName,
		MType: "unknown",
		Delta: &metricValue,
		Value: nil,
	}
	body, err := json.Marshal(req)
	require.NoError(s.T(), err)

	result, _ := s.testRequest(
		http.MethodPost,
		"/update/",
		string(body),
		nil,
	)
	_ = result.Body.Close()

	assert.Equal(s.T(), http.StatusBadRequest, result.StatusCode)
}
