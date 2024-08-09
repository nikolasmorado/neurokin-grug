package types

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type Storage interface {
	CreateWaitlist(email string) error

	Login(email, password string) (string, error)
	CreateAccount(account *Account, password string) error
	GetAccountById(id uuid.UUID) (*Account, error)
  GetAccountByEmail(email string) (*Account, error)

	CreateQuizAnswer(quizAnswer *QuizAnswer) error
	GetQuizAnswer(quizSlug string, userId uuid.UUID) (*QuizAnswer, error)
	UpdateQuizAnswer(quizAnswer *QuizAnswer) error
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
	Id        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Email     string    `json:"email"`
	Verified  bool      `json:"verified"`
	// Phone     string `json:"phone"`
}

type QuizQuestion struct {
	Question   string
	Answers    []string
	QuestionId string
}

type QuizAnswer struct {
	Id        uuid.UUID       `json:"id"`
	CreatedAt time.Time       `json:"created_at"`
	QuizSlug  string          `json:"quiz_slug"`
	Answers   json.RawMessage `json:"answers"`
	UserId    uuid.UUID       `json:"user_id"`
	Completed bool            `json:"completed"`
}


type Quiz struct {
  Name string
  Slug string
  Questions []QuizQuestion
}
