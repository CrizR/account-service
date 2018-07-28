package data

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"github.com/ecclesia-dev/account-service/models"
	"google.golang.org/api/option"
)

type Firebase struct {
	app *firebase.App
}

func NewFirebase() DataAccess {
	opt := option.WithCredentialsFile("keys/ecclesia-firebase-key.json")

	// TODO: set FIREBASE_CONFIG as an envornment variable so config can
	// 		 be passed in as nil.
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	return Firebase{app: app}
}

func (fb Firebase) CreateUser([]string) error {

	return nil
}

func (fb Firebase) FindAllUsers() ([]models.Account, error) {

	return nil, nil
}

func (fb Firebase) FindUserByID(string) (models.Account, error) {

	return models.Account{
		ID:          0,
		AccountType: models.StandardAccount,
		Email:       "",
		Password:    "",
		FirstName:   "",
		LastName:    "",
		Bio:         "",
		Industry:    "",
		Education:   "",
		State:       "",
		Reputation:  0,
		Interests:   nil,
	}, nil
}

func (fb Firebase) FindUserByEmail(string) (models.Account, error) {

	return models.Account{
		ID:          0,
		AccountType: models.StandardAccount,
		Email:       "",
		Password:    "",
		FirstName:   "",
		LastName:    "",
		Bio:         "",
		Industry:    "",
		Education:   "",
		State:       "",
		Reputation:  0,
		Interests:   nil,
	}, nil
}

func (fb Firebase) UpdateUser(string) error {

	return nil
}

func (fb Firebase) RemoveUser(string) error {

	return nil
}
