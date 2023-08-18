package server

import "github.com/ilyabukanov123/todo-app/internal/app/todo-app/config"

// Server ...
type Server struct {
	config *config.Config
}

// NewServer ...
func NewServer(c *config.Config) *Server {
	return &Server{
		config: c,
	}
}

// Start ...
func (s *Server) Start() error {
	if err := s.setupLogger(); err != nil {
		//throw error
		return err
	}
	return nil
}

// Stop ...
func (s *Server) Stop() error {
	return nil
}

// setupLogger ...
func (s *Server) setupLogger() error {
	return nil
}
