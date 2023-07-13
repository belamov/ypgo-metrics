package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/belamov/ypgo-metrics/internal/app/handlers"
	"github.com/belamov/ypgo-metrics/internal/app/services"
)

type HttpServer struct {
	server *http.Server
}

func NewHttpServer(service services.MetricServiceInterface) (Server, error) {
	return &HttpServer{
		server: &http.Server{
			Addr:              "0.0.0.0:8080",
			Handler:           handlers.NewRouter(service),
			ReadHeaderTimeout: 1 * time.Second,
		},
	}, nil
}

func (s *HttpServer) Run() {
	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint
		log.Println("Shutting Down Server")
		// We received an interrupt signal, shut down.
		if err := s.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()

	if err := s.server.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}

	<-idleConnsClosed
	log.Println("Goodbye")
}

func (s *HttpServer) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
