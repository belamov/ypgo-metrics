package middleware

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"

	"github.com/rs/zerolog/log"
)

func GzipDecompressor(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contentEncoding := r.Header.Get("Content-Encoding")
		sendsGzip := strings.Contains(contentEncoding, "gzip")
		if sendsGzip {
			// TODO: sync.Pool for optimization?
			// https://www.sobyte.net/post/2022-06/go-sync-pool/
			// https://developer20.com/using-sync-pool/
			cr, err := newCompressReader(r.Body)
			if err != nil {
				log.Error().Err(err).Msg("failed init decompressor")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			r.Body = cr
			defer func(cr *compressReader) {
				err := cr.Close()
				if err != nil {
					log.Error().Err(err).Msg("failed closing compress reader")
				}
			}(cr)
		}

		next.ServeHTTP(w, r)
	})
}

type compressReader struct {
	r  io.ReadCloser
	zr *gzip.Reader
}

func newCompressReader(r io.ReadCloser) (*compressReader, error) {
	zr, err := gzip.NewReader(r)
	if err != nil {
		return nil, err
	}

	return &compressReader{
		r:  r,
		zr: zr,
	}, nil
}

func (c *compressReader) Read(p []byte) (n int, err error) {
	return c.zr.Read(p)
}

func (c *compressReader) Close() error {
	if err := c.r.Close(); err != nil {
		return err
	}
	return c.zr.Close()
}
