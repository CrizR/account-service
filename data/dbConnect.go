package data

import (
	"context"
	"log"

	"firebase.google.com/go"
	"google.golang.org/api/option"
	//"firebase.google.com/go/auth"
	"user-service/models"
	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type Firebase struct {
	app    *firebase.App
	client *firestore.Client
}

func (fb *Firebase) New() {
	opt := option.WithCredentialsFile("keys/ecclesia-firebase-key.json")

	// TODO: set FIREBASE_CONFIG as an envornment variable so config can be passed in as nil.
	var err error
	fb.app, err = firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalln(err)
	}
	fb.client, err = fb.app.Firestore(context.Background())
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
}

func (fb *Firebase) CreateUser(user map[string]interface{}) (error) {
	var err error
	_, _, err = fb.client.Collection("users").Add(context.Background(), user)
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}
	return err
}

func (fb *Firebase) FindAllUsers(string) ([]models.Account, error) {
	iter := fb.client.Collection("users").Documents(context.Background())
	var accounts []models.Account
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		accounts = append(accounts, fb.converter(doc.Data()))
	}

	return accounts, nil
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

func (fb * Firebase) converter(data map[string]interface{}) (models.Account) {
	return models.Account{
		ID:          data["ID"].(int),
		AccountType: data["AccountType"].(models.AccountType),
		Email:       data["Email"].(string),
		Password:    data["Password"].(string),
		FirstName:   data["FirstName"].(string),
		LastName:    data["LastName"].(string),
		Bio:         data["Bio"].(string),
		Industry:    data["Industry"].(string),
		Education:   data["Education"].(string),
		State:       data["State"].(string),
		Reputation:  data["Reputation"].(int),
		Interests:   data["Interests"].([]string),
	}
}
