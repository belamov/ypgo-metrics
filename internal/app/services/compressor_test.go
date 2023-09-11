package services

import (
	"bytes"
	"compress/gzip"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGzipCompressor_GetCompressedReader(t *testing.T) {
	data := "some data"
	compressor := NewGzipCompressor()
	compressedReader, err := compressor.GetCompressedReader([]byte(data))
	require.NoError(t, err)

	decompressor, err := gzip.NewReader(compressedReader)
	defer func(decompressor *gzip.Reader) {
		err := decompressor.Close()
		assert.NoError(t, err)
	}(decompressor)

	require.NoError(t, err)

	var decompressed bytes.Buffer
	n, err := decompressed.ReadFrom(decompressor)
	require.NoError(t, err)
	assert.EqualValues(t, len(data), n)
	assert.Equal(t, data, decompressed.String())
}

func TestGzipCompressor_SetHeader(t *testing.T) {
	req, err := http.NewRequest("post", "some url", nil)
	require.NoError(t, err)

	compressor := NewGzipCompressor()
	compressor.SetHeader(req)

	header := req.Header.Get("Content-Encoding")
	assert.NotEmpty(t, header)
	assert.Equal(t, header, "gzip")
}
