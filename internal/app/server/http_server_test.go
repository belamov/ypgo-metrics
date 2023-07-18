package server

import (
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"syscall"
	"testing"
	"time"

	"github.com/belamov/ypgo-metrics/internal/app/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestHttpServer_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockService := mocks.NewMockMetricServiceInterface(ctrl)

	port := chooseRandomUnusedPort()
	serverAddress := fmt.Sprintf("0.0.0.0:%d", port)
	fmt.Println(serverAddress)
	server := NewHttpServer(serverAddress, mockService)

	done := make(chan struct{})
	go func() {
		<-done
		e := syscall.Kill(syscall.Getpid(), syscall.SIGINT)
		require.NoError(t, e)
	}()

	finished := make(chan struct{})
	go func() {
		server.Run()
		close(finished)
	}()

	// defer cleanup because require check below can fail
	defer func() {
		close(done)
		<-finished
	}()

	waitForHTTPServerStart(port)
	resp, err := http.Get(fmt.Sprintf("http://localhost:%d/", port))
	require.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func chooseRandomUnusedPort() (port int) {
	for i := 0; i < 10; i++ {
		port = 40000 + int(rand.Int31n(10000))
		if ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port)); err == nil {
			_ = ln.Close()
			break
		}
	}
	return port
}

func waitForHTTPServerStart(port int) {
	// wait for up to 10 seconds for server to start before returning it
	client := http.Client{Timeout: time.Second}
	defer client.CloseIdleConnections()
	for i := 0; i < 100; i++ {
		time.Sleep(time.Millisecond * 100)
		if resp, err := client.Get(fmt.Sprintf("http://localhost:%d", port)); err == nil {
			_ = resp.Body.Close()
			return
		}
	}
}
