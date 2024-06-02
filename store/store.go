package store

import (
	"cloud.google.com/go/firestore"
	"context"
	"google.golang.org/api/option"
)

func InitFirebase(f string) (s *Store, err error) {
	ctx := context.Background()
	opt := option.WithCredentialsFile(f)

	client, err := firestore.NewClient(ctx, "neurokin", opt)

	if err != nil {
		return nil, err
	}

	return &Store{
		db: client,
	}, nil
}

type Store struct {
	db *firestore.Client
}

func (s *Store) CreateWaitlist(email string) error {
	ctx := context.Background()

	_, _, err := s.db.Collection("waitlist").Add(ctx, map[string]interface{}{
		"email": email,
	})

	return err
}
