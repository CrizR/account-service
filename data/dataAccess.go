package data

import (
	"fmt"

	"golang.org/x/net/context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"

	"google.golang.org/api/option"
)

type DataAccess struct {
	ecclesiaDb *firebase.App
}

opt := option.WithCredentialsFile("keys/ecclesia-firebase-key.json")

app, err := firebase.NewApp(context.Background(), nil, opt)
if err != nil {
	log.Fatalf("error initializing app: %v\n", err)
}

