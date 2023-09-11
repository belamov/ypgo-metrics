package handlers

import (
	"fmt"
	"net/http"

	"github.com/belamov/ypgo-metrics/internal/app/models"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func (s *HandlersTestSuite) TestGetAllMetrics() {
	gaugeMetric := models.GaugeMetric{
		Name:  "gauge",
		Value: 1,
	}
	counterMetric := models.CounterMetric{
		Name:  "counter",
		Value: 2,
	}

	s.mockService.EXPECT().
		GetAllMetrics(gomock.Any()).
		Times(1).
		Return([]models.CounterMetric{counterMetric}, []models.GaugeMetric{gaugeMetric}, nil)

	result, response := s.testRequest(
		http.MethodGet,
		"/",
		"",
		nil,
	)
	_ = result.Body.Close()

	fmt.Println(response)
	assert.Equal(s.T(), http.StatusOK, result.StatusCode)
	goldenResponse := "<html lang=\"ru\">\n<body>\n<table>\n    <tr>\n        <th>Метрика</th>\n        <th>Тип</th>\n        <th>Значение</th>\n    </tr>\n    \n        <tr>\n            <td>\n                counter\n            </td>\n            <td>Counter</td>\n            <td>2</td>\n        </tr>\n    \n    \n        <tr>\n            <td>\n                gauge\n            </td>\n            <td>Gauge</td>\n            <td>1</td>\n        </tr>\n    \n</table>\n</body>\n</html>"
	assert.Equal(s.T(), goldenResponse, response)
}
