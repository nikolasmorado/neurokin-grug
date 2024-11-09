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

	CreateQuizAnswer(id, slug string) error
	GetQuizAnswer(quizSlug string, id uuid.UUID) (*QuizAnswer, error)
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

type QuestionType string

const (
	QTypeMultipleChoice QuestionType = "mc"
	QTypeText           QuestionType = "text"
)

type QuizQuestion struct {
	Question string       `json:"question"`
	Answers  []string     `json:"answers"`
	Id       string       `json:"id"`
	Type     QuestionType `json:"type"`
}

type QuizAnswer struct {
	Id             int             `json:"id"`
	CreatedAt      time.Time       `json:"created_at"`
	QuizSlug       string          `json:"quiz_slug"`
	RawAnswers     json.RawMessage `json:"answers"`
	DecodedAnswers []Answer        `json:"decoded_answers"`
	UserId         uuid.UUID       `json:"user_id"`
	Completed      bool            `json:"completed"`
}

type Answer struct {
	QuestionId string `json:"question_id"`
	Answer     string `json:"answer"`
	Context    string `json:"context"`
	Bookmarked bool   `json:"bookmarked"`
}

type Quiz struct {
	Id        int            `json:"id"`
	Name      string         `json:"name"`
	Slug      string         `json:"slug"`
	Questions []QuizQuestion `json:"questions"`
}

type Orientation string

const (
	OrientationLTR Orientation = "ltr"
	OrientationRTL Orientation = "rtl"
)

type Section struct {
	Id          int         `json:"id"`
	CreatedAt   time.Time   `json:"created_at"`
	Subtitle    string      `json:"subtitle"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Cta         string      `json:"cta"`
	Link        string      `json:"link"`
	Orientation Orientation `json:"orientation"`
	BgColor     string      `json:"bg-color"`
	Img         string      `json:"img"`
}

type Tile struct {
	Id        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
}

type Task struct {
	Id          int       `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	Slug        string    `json:"slug"`
	Type        string    `json:"type"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Eta         string    `json:"eta"`
	Prereqs     []string  `json:"prereqs"`
}
