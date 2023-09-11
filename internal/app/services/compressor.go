package services

import (
	"bytes"
	"compress/gzip"
	"io"
	"net/http"

	"github.com/rs/zerolog/log"
)

type Compressor interface {
	GetCompressedReader(data []byte) (io.Reader, error)
	SetHeader(req *http.Request)
}

type GzipCompressor struct{}

func NewGzipCompressor() *GzipCompressor {
	return &GzipCompressor{}
}

func (g GzipCompressor) GetCompressedReader(data []byte) (io.Reader, error) {
	b := new(bytes.Buffer)
	// TODO: sync.pool for writer?
	// https://www.sobyte.net/post/2022-06/go-sync-pool/
	w, err := gzip.NewWriterLevel(b, gzip.BestSpeed)
	if err != nil {
		log.Error().Err(err).Msg("error init gzip writer")
		return nil, err
	}
	_, err = w.Write(data)
	if err != nil {
		log.Error().Err(err).Msg("error compressing data")
		return nil, err
	}
	err = w.Close()
	w.Reset(b)
	if err != nil {
		log.Error().Err(err).Msg("error closing writer")
		return nil, err
	}

	return b, nil
}

func (g GzipCompressor) SetHeader(req *http.Request) {
	req.Header.Set("Content-Encoding", "gzip")
}
