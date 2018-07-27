package data

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	//"firebase.google.com/go/auth"
)

type Firebase struct {
	app *firebase.App
}

func (fb *Firebase) New() {
	opt := option.WithCredentialsFile("keys/ecclesia-firebase-key.json")

	// TODO: set FIREBASE_CONFIG as an envornment variable so config can
	// 		 be passed in as nil.
	var err error
	fb.app, err = firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

}

func (fb *Firebase) CreateUser([]string) (string, error) {

	return "", nil
}

func (fb *Firebase) FindAllUsers(string) (string, error) {

	return "", nil
}

func (fb *Firebase) FindUserById(string) (string, error) {

	return "", nil
}

func (fb *Firebase) FindUserByEmail(string) (string, error) {

	return "", nil
}

func (fb *Firebase) UpdateUser(string) (string, error) {

	return "", nil
}

func (fb *Firebase) RemoveUser(string) (string, error) {

	return "", nil
}
