package server

import (
	"net/http"
	"time"

	"github.com/Hendra-Huang/go-standard-layout/log"
)

// Options of
type Options struct {
	ListenAddress string
}

// Server struct
type Server struct {
	options Options
	birth   time.Time
}

// New web handler
func New(opts Options) *Server {
	s := &Server{
		options: opts,
		birth:   time.Now(),
	}

	return s
}

// Serve web service
func (s *Server) Serve(handler http.Handler) error {
	log.Infof("Starting webserver on %s", s.options.ListenAddress)

	return http.ListenAndServe(s.options.ListenAddress, handler)
}
