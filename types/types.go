package types

import (
	"net/http"
)

type Storage interface {
	CreateWaitlist(email string) error

	CreateAccount(account *Account, password string) error
  GetAccount(email string) (*Account, error)

  Login(email, password string) (string, error)
}

type HTTPHandler func(w http.ResponseWriter, r *http.Request) error

type WaitlistRequest struct {
	Email string `json:"email"`
}

type RegisterRequest struct {
	Email        string `json:"email"`
	Password     string `json:"password"`
	ConfPassword string `json:"conf-password"`
}

type Account struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	// Phone     string `json:"phone"`
}
