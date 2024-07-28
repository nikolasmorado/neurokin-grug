package handlers

import (
	"net/http"
	"neurokin/views/auth"
	"strings"

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
	email := r.FormValue("email")
	password := r.FormValue("password")

	if !u.ValidateEmail(email) {
		return u.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Please enter a valid email address."})
	}

	token, err := s.Login(email, password)
	if err != nil {
		if strings.Contains(err.Error(), "EMAIL_NOT_FOUND") {
			return u.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Email not found."})
		}
		return u.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid email or password."})
	}

	// Set the token in a cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
	})

  w.WriteHeader(http.StatusOK)
  return nil
}

func register(w http.ResponseWriter, r *http.Request, s t.Storage) error {

	email := r.FormValue("email")

	if !u.ValidateEmail(email) {
		return u.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Please enter a valid email address."})
	}

	password := r.FormValue("password")
	confPassword := r.FormValue("conf-password")

	if password != confPassword {
		return u.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Passwords must match."})
	}

	acc := new(t.Account)

	acc.Email = email

	if err := s.CreateAccount(acc, password); err != nil {
		if strings.Contains(err.Error(), "EMAIL_EXISTS") {
			return u.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Email already exists."})
		}
		return err
	}

	return u.WriteJSON(w, http.StatusOK, acc)
}

func RenderLogin(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, auth.Login())
}

func RenderSignup(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, auth.Signup())
}
