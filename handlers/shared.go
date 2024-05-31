package handlers

import (
	"log/slog"
	"net/http"

  t "github.com/nikolasmorado/neurokin-grug/types"
	"github.com/a-h/templ"
)


func Make(h t.HTTPHandler) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    if err := h(w, r); err != nil {
      slog.Error("HTTP Handler error", "err", err, "path", r.URL.Path)
    }
  }
}

func Render(w http.ResponseWriter, r *http.Request, c templ.Component) error {
  return c.Render(r.Context(), w)
}

