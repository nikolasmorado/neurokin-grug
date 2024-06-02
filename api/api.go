package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	h "neurokin/handlers"
)

type Server struct {
	listenAddr string
  public http.Handler
}

func NewServer(listenAddr string, public http.Handler) *Server {
	return &Server{
		listenAddr: listenAddr,
    public: public,
	}
}

func (s *Server) Start() error {

	router := chi.NewMux()

  router.Handle("/*", s.public)

	router.Get("/api/health", h.Make(h.HandleHealth))

  router.Get("/", h.Make(h.HandleHome))

  router.Post("/waitlist", h.Make(h.HandlePostWaitlist))

  return http.ListenAndServe(s.listenAddr, router)
}
