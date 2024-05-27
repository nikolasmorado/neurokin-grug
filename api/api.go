package api

import (
	"github.com/go-chi/chi/v5"
)

type Server struct {
	listenAddr string
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
	}
}

func (s *Server) Start() error {

  router := chi.NewMux();

  router.Get("/api/health",

}
