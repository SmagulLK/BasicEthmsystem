// Package httpserver provides a simple HTTP server.
package httpserver

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"
)

//todo change default to getting env
const (
	defaultHost              string        = "0.0.0.0"
	defaultPort              string        = "8080"
	defaultMaxHeaderBytes    int           = 1 << 20 // 1MB
	defaultReadTimeout       time.Duration = 5 * time.Second
	defaultReadHeaderTimeout time.Duration = time.Minute
	defaultWriteTimeout      time.Duration = 5 * time.Second
)

// Server is a simple HTTP server.
type Server struct {
	host              string
	port              string
	maxHeaderBytes    int
	readTimeout       time.Duration
	readHeaderTimeout time.Duration
	writeTimeout      time.Duration
	httpServer        *http.Server
}

// New creates a new Server.
func New(httpHandler http.Handler, opts ...Option) Server {
	server := Server{
		host:              defaultHost,
		port:              defaultPort,
		maxHeaderBytes:    defaultMaxHeaderBytes,
		readTimeout:       defaultReadTimeout,
		readHeaderTimeout: defaultReadHeaderTimeout,
		writeTimeout:      defaultWriteTimeout,
	}

	for _, opt := range opts {
		opt(&server)
	}

	server.httpServer = &http.Server{
		Addr:              net.JoinHostPort(server.host, server.port),
		Handler:           httpHandler,
		MaxHeaderBytes:    server.maxHeaderBytes,
		ReadTimeout:       server.readTimeout,
		ReadHeaderTimeout: server.readHeaderTimeout,
		WriteTimeout:      server.writeTimeout,
	}

	return server
}

// Start starts the http server.
func (s Server) Start() error {
	fmt.Println("Starting http server successfully, ", s.httpServer.Addr)
	if err := s.httpServer.ListenAndServe(); err != nil {
		return fmt.Errorf("failed to start http server: %w", err)
	}

	return nil
}

// Stop stops the http server.
func (s Server) Stop(ctx context.Context, shutdownTimeout time.Duration) error {
	ctx, cancel := context.WithTimeout(ctx, shutdownTimeout)
	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		return fmt.Errorf("failed to shutdown http server: %w", err)
	}

	return nil
}
