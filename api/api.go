package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nikolasmorado/neurokin-grug/handlers"
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

	router := chi.NewMux()

	router.Get("/api/health", handlers.Make(handlers.HandleHealth))

  return http.ListenAndServe(s.listenAddr, router)
}
