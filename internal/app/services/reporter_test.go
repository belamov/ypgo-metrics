package services

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/belamov/ypgo-metrics/internal/app/resources"

	"github.com/stretchr/testify/assert"
)

type NilCompressor struct{}

func (n NilCompressor) GetCompressedReader(data []byte) (io.Reader, error) {
	return bytes.NewReader(data), nil
}

func (n NilCompressor) SetHeader(req *http.Request) {
}

func NewNilCompressor() *NilCompressor {
	return &NilCompressor{}
}

type recordingTransport struct {
	req *http.Request
}

func (t *recordingTransport) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	t.req = req
	return &http.Response{StatusCode: http.StatusOK}, nil
}

func TestHTTPReporter_Report(t *testing.T) {
	tr := recordingTransport{}

	client := &http.Client{Transport: &tr}
	updateURL := "update_url"

	reporter := NewHTTPReporter(client, updateURL, NewNilCompressor())

	value := new(float64)
	*value = 10

	metrics := []resources.Metric{{
		ID:    "name",
		MType: "type",
		Delta: nil,
		Value: value,
	}}

	reporter.Report(metrics)

	assert.Equal(t,
		updateURL,
		tr.req.URL.String(),
	)

	expectedReq, err := json.Marshal(metrics[0])
	assert.NoError(t, err)
	reqBody, err := io.ReadAll(tr.req.Body)
	assert.NoError(t, err)
	err = tr.req.Body.Close()
	assert.NoError(t, err)

	assert.JSONEq(t, string(expectedReq), string(reqBody))
}
