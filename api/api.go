package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	h "neurokin/handlers"
	t "neurokin/types"
	u "neurokin/util"
)

type Server struct {
	listenAddr string
	store      t.Storage
	public     http.Handler
}

func NewServer(listenAddr string, store t.Storage, public http.Handler) *Server {
	return &Server{
		listenAddr: listenAddr,
		store:      store,
		public:     public,
	}
}

func (s *Server) Start() error {
	router := chi.NewMux()

	deps := u.NewHandlerDependencies(s.store)

	router.Handle("/*", s.public)

	router.Get("/api/health", h.Make(h.HandleHealth))

	router.Get("/", h.Make(h.HandleHome))
	router.Post("/waitlist", h.Make(h.HandleWaitlist(deps)))

	router.Get("/login", h.Make(h.HandleLogin))
	router.Get("/signup", h.Make(h.HandleSignup))

	return http.ListenAndServe(s.listenAddr, router)
}
