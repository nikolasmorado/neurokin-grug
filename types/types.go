package types

import (
	"net/http"
)

type Storage interface {
	CreateWaitlist(email string) error
}

type HTTPHandler func(w http.ResponseWriter, r *http.Request) error

type WaitlistRequest struct {
	Email string `json:"email"`
}

type Account struct {
	Id        string `json:"id"`
	Email     string `json:"email"`
  Phone     string `json:"phone"`
}
