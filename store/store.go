package store

import (
	"encoding/json"
	t "neurokin/types"
	"strconv"

	"github.com/google/uuid"
	sgtgt "github.com/supabase-community/gotrue-go/types"
	"github.com/supabase-community/supabase-go"
)

func InitSupabase(u string, k string) (s *Store, err error) {
	client, err := supabase.NewClient(u, k, nil)

	if err != nil {
		return nil, err
	}

	return &Store{db: client}, nil
}

type Store struct {
	db *supabase.Client
}

func (s *Store) CreateWaitlist(email string) error {
	_, _, err := s.db.From("waitlist").Insert(
		[]map[string]interface{}{{"email": email}},
		false,
		"",
		"minimal",
		"",
	).Execute()

	if err != nil {
		return err
	}

	return nil
}

func (s *Store) Login(email, password string) (string, error) {
	res, err := s.db.Auth.SignInWithEmailPassword(email, password)

	if err != nil {
		return "", err
	}

	return res.AccessToken, nil
}

func (s *Store) CreateAccount(account *t.Account, password string) error {
	req := sgtgt.SignupRequest{
		Email:    account.Email,
		Password: password,
	}

	res, err := s.db.Auth.Signup(req)

	if err != nil {
		return err
	}

	account.Id = res.User.ID

	_, _, err = s.db.From("users").Insert(
		[]map[string]interface{}{{
			"id":    account.Id,
			"email": account.Email,
		}},
		false,
		"",
		"minimal",
		"",
	).Execute()

	if err != nil {
		return err
	}

	return nil

}

func (s *Store) CreateQuizAnswer(id, slug string) error {
	_, _, err := s.db.From("quiz_answers").Insert(
		[]map[string]interface{}{{
			"quiz_slug": slug,
			"user_id":   id,
		}},
		false,
		"",
		"minimal",
		"",
	).Execute()

	if err != nil {
		return err
	}

	return nil
}

func (s *Store) GetAccountByEmail(email string) (*t.Account, error) {
	res, _, err := s.db.From("users").Select("*", "exact", false).Eq(
		"email",
		email,
	).Single().Execute()

	if err != nil {
		return nil, err
	}

	var account t.Account
	err = json.Unmarshal(res, &account)

	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (s *Store) GetAccountById(id uuid.UUID) (*t.Account, error) {
	res, _, err := s.db.From("users").Select("*", "exact", false).Eq(
		"id",
		id.String(),
	).Execute()

	if err != nil {
		return nil, err
	}

	var account t.Account
	err = json.Unmarshal(res, &account)

	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (s *Store) GetQuizAnswer(quizSlug string, id uuid.UUID) (*t.QuizAnswer, error) {
	res, _, err := s.db.From("quiz_answers").Select("*", "exact", false).Eq(
		"quiz_slug",
		quizSlug,
	).Eq(
		"user_id",
		id.String(),
	).Single().Execute()

	if res == nil {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	var quizAnswer t.QuizAnswer
	err = json.Unmarshal(res, &quizAnswer)

	if err != nil {
		return nil, err
	}

	var answers []t.Answer
	err = json.Unmarshal(quizAnswer.RawAnswers, &answers)

	if err != nil {
		return nil, err
	}

	quizAnswer.DecodedAnswers = answers

	return &quizAnswer, nil
}

func (s *Store) UpdateQuizAnswer(quizAnswer *t.QuizAnswer) error {
  an, err := json.Marshal(quizAnswer.DecodedAnswers)

  if err != nil {
    return err
  }

  quizAnswer.RawAnswers = an

	_, _, err = s.db.From("quiz_answers").Update(
    []map[string]interface{}{{
      "answers": quizAnswer.RawAnswers,
      "completed": quizAnswer.Completed,
    }},
		"minimal",
		"exact",
	).Eq(
    "quiz_slug",
    quizAnswer.QuizSlug,
  ).Eq(
    "user_id",
    quizAnswer.UserId.String(),
  ).Eq(
    "id",
    strconv.Itoa(quizAnswer.Id),
  ).Execute()

	if err != nil {
		return err
	}

	return nil
}
