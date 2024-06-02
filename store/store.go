package store

import (
	"context"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

func InitFirebase(f string) (s *Store, err error) {
	ctx := context.Background()
	opt := option.WithCredentialsFile(f)

	client, err := firestore.NewClient(ctx, "neurokin-df808", opt)

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

	doc := s.db.Collection("waitlist").Doc("waitlist")

	_, err := doc.Update(ctx, []firestore.Update{
		{Path: "emails", Value: firestore.ArrayUnion(email)},
	})

	return err
}
