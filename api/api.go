package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	a "neurokin/auth"
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
	router.Get("/about", h.Make(h.HandleAbout))
	router.Post("/waitlist", h.Make(h.HandleWaitlist(deps)))

	router.HandleFunc("/login", h.Make(h.HandleLogin(deps)))
	router.HandleFunc("/signup", h.Make(h.HandleSignup(deps)))
  router.HandleFunc("/logout", h.Make(h.HandleLogout()))

	router.HandleFunc("/dashboard", a.WithJWT(h.Make(h.HandleDashboard(deps)), s.store))

	router.HandleFunc("/verify", h.Make(h.HandleVerify))

  router.HandleFunc("/quiz/{slug}", h.Make(h.HandleQuizSlug(deps)))

	return http.ListenAndServe(s.listenAddr, router)
}
