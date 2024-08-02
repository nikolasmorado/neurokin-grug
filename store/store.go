package store

import (
	"github.com/supabase-community/supabase-go"
  sgtgt "github.com/supabase-community/gotrue-go/types"
	t "neurokin/types"
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
