package handlers

import (
	"net/http"
	"neurokin/views/auth"
)


func HandleLogin(w http.ResponseWriter, r *http.Request) error {
  return Render(w, r, auth.Login())
}

func HandleSignup(w http.ResponseWriter, r *http.Request) error {
  return Render(w, r, auth.Signup())
}
