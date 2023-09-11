package middleware

import (
	"bytes"
	"compress/gzip"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDecompressor(t *testing.T) {
	originalBody := "some text"
	buf := bytes.NewBuffer(nil)
	zb := gzip.NewWriter(buf)
	_, err := zb.Write([]byte(originalBody))
	require.NoError(t, err)
	err = zb.Close()
	require.NoError(t, err)

	r := httptest.NewRequest("POST", "https://some-url.example", buf)
	r.RequestURI = ""
	r.Header.Set("Content-Encoding", "gzip")

	h := func(w http.ResponseWriter, r *http.Request) {
		reqBody, err := io.ReadAll(r.Body)
		assert.NoError(t, err)
		if string(reqBody) != originalBody {
			t.Error("request body is not decompressed")
		}
	}

	decompressedHandler := GzipDecompressor(http.HandlerFunc(h))
	decompressedHandler.ServeHTTP(nil, r)
}

func TestDecompressorNoHeader(t *testing.T) {
	originalBody := "some text"
	buf := bytes.NewBuffer(nil)
	buf.Write([]byte(originalBody))

	r := httptest.NewRequest("POST", "https://some-url.example", buf)
	r.RequestURI = ""

	h := func(w http.ResponseWriter, r *http.Request) {
		reqBody, err := io.ReadAll(r.Body)
		assert.NoError(t, err)

		if string(reqBody) != originalBody {
			t.Error("request body is decompressed without required header")
		}
	}

	decompressedHandler := GzipDecompressor(http.HandlerFunc(h))
	decompressedHandler.ServeHTTP(nil, r)
}
