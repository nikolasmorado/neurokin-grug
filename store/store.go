package store

import (
	"context"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"

	t "neurokin/types"
)

func InitFirebase(f string, i string) (s *Store, err error) {
	ctx := context.Background()
	opt := option.WithCredentialsFile(f)

	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, err
	}

	aC, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}

	fC, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}

	return &Store{
		db:   fC,
		auth: aC,
	}, nil
}

type Store struct {
	db   *firestore.Client
	auth *auth.Client
}

func (s *Store) CreateWaitlist(email string) error {
	ctx := context.Background()

	doc := s.db.Collection("waitlist").Doc("waitlist")

	_, err := doc.Update(ctx, []firestore.Update{
		{Path: "emails", Value: firestore.ArrayUnion(email)},
	})

	return err
}

func (s *Store) Login(email, password string) (string, error) {
	ctx := context.Background()

	params := (&auth.UserToCreate{}).
		Email(email).
		Password(password)

	u, err := s.auth.CreateUser(ctx, params)
	if err != nil {
		return "", err
	}

	return u.UID, nil
}

func (s *Store) CreateAccount(account *t.Account, password string) error {
	ctx := context.Background()

	params := (&auth.UserToCreate{}).
		Email(account.Email).
		Password(password)

	u, err := s.auth.CreateUser(ctx, params)
	if err != nil {
		return err
	}

	account.Id = u.UID

	doc := s.db.Collection("accounts").Doc(account.Id)

	_, err = doc.Set(ctx, account)

	return err
}

func (s *Store) GetAccount(id string) (*t.Account, error) {
	ctx := context.Background()

	doc := s.db.Collection("accounts").Doc(id)

	snap, err := doc.Get(ctx)
	if err != nil {
		return nil, err
	}

	var acc t.Account

	if err := snap.DataTo(&acc); err != nil {
		return nil, err
	}

	return &acc, nil
}
