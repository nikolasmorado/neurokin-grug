package store

import (
	"github.com/supabase-community/supabase-go"

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
		"representation",
		"",
	).Execute()

	if err != nil {
		return err
	}

	return nil
}

func (s *Store) Login(email, password string) (string, error) {
	return "", nil
}

func (s *Store) CreateAccount(account *t.Account, password string) error {
	return nil
}
