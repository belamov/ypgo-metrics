package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/belamov/ypgo-metrics/internal/app/handlers"
	"github.com/belamov/ypgo-metrics/internal/app/services"
)

type HTTPServer struct {
	server *http.Server
}

func NewHTTPServer(addr string, service services.MetricServiceInterface) *HTTPServer {
	return &HTTPServer{
		server: &http.Server{
			Addr:              addr,
			Handler:           handlers.NewRouter(service),
			ReadHeaderTimeout: 1 * time.Second,
		},
	}
}

func (s *HTTPServer) Run() {
	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint
		log.Println("Shutting Down Server")
		// We received an interrupt signal, shut down.
		if err := s.shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()

	if err := s.server.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		fmt.Println(err)
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}

	<-idleConnsClosed
	log.Println("Goodbye")
}

func (s *HTTPServer) shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
