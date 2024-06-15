package handlers

import (
	"net/http"
	"neurokin/views/auth"

  t "neurokin/types"
  u "neurokin/util"
)

func HandleLogin(deps *u.HandlerDependencies) t.HTTPHandler {
  s := deps.Store

  return func(w http.ResponseWriter, r *http.Request) error {
    switch r.Method {
      case "GET":
        return RenderLogin(w, r)
      case "POST":
        return login(w, r, s)
      default:
        return RenderLogin(w, r)
    }
  }
}

func HandleSignup(deps *u.HandlerDependencies) t.HTTPHandler {
  s := deps.Store

  return func(w http.ResponseWriter, r *http.Request) error {
    switch r.Method {
      case "GET":
        return RenderSignup(w, r)
      case "POST":
        return register(w, r, s)
      default:
        return RenderSignup(w, r)
    }
  }
}

func login(w http.ResponseWriter, r *http.Request, s t.Storage) error {
  return u.WriteJSON(w, http.StatusOK, map[string]string{"message": "login"})
}

func register(w http.ResponseWriter, r *http.Request, s t.Storage) error {
  return u.WriteJSON(w, http.StatusOK, map[string]string{"message": "register"})
}

func RenderLogin(w http.ResponseWriter, r *http.Request) error {
  return Render(w, r, auth.Login())
}

func RenderSignup(w http.ResponseWriter, r *http.Request) error {
  return Render(w, r, auth.Signup())
}
