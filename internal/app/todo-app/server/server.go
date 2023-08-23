package server

import (
	"context"
	"github.com/ilyabukanov123/todo-app/internal/app/todo-app/config"
	"net/http"
	"time"
)

// Server ...
type Server struct {
	config     *config.Config
	handler    http.Handler
	httpServer *http.Server
}

// NewServer ...
func NewServer(cfg *config.Config, h http.Handler, port string) *Server {
	return &Server{
		config:  cfg,
		handler: h,
		httpServer: &http.Server{
			Addr:           ":" + port,
			MaxHeaderBytes: 1 << 20,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second},
	}
}

// Start ...
func (s *Server) Start() error {
	if err := s.setupLogger(); err != nil {
		//throw error
		return err
	}
	return s.httpServer.ListenAndServe()
}

// Stop ...
func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

// setupLogger ...
func (s *Server) setupLogger() error {
	return nil
}
