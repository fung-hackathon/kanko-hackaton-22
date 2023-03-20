package infra

import (
	"context"
	"errors"
	"kanko-hackaton-22/app/config"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

var (
	ErrFirebaseInit = errors.New("failed to initialize firebase instance")
	ErrFirestore    = errors.New("failed to establish connection to Firestore")
)

type Firestore struct {
	Client  *firestore.Client
	Context context.Context
}

func Initialize() (*Firestore, error) {
	ctx := context.Background()
	opt := option.WithCredentialsFile(config.GOOGLE_APPLICATION_CREDENTIALS)
	//conf := &firebase.Config{ProjectID: "funhackathon22"}
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, ErrFirebaseInit
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Println(err)
		return nil, ErrFirestore
	}

	return &Firestore{Client: client, Context: ctx}, nil
}
