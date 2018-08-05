package data

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/ecclesia-dev/account-service/models"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type Firebase struct {
	app    *firebase.App
	client *firestore.Client
	auth   *auth.Client
}

func NewFirebase() DataAccess {
	opt := option.WithCredentialsFile("../keys/ecclesia-firebase-key.json")
	// TODO: set FIREBASE_CONFIG as an envornment variable so config can
	// 		 be passed in as nil.
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalln(err)
	}
	client, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	auth, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	return Firebase{app: app, client: client, auth:auth}
}

func (fb Firebase) CreateAccount(account models.Account) error {
	params := (&auth.UserToCreate{}).
		Email(account.Email).
		Password(account.Password).
		EmailVerified(true)
	account.Email = ""
	account.Password = ""
	//_, err := auth.Client.CreateUser(context.Background(), params)
	_, _, err := fb.client.Collection("users").Add(context.Background(), account)
	if err != nil {
		log.Fatalf("Failed adding user: %v", err)
	}
	return err
}

func (fb Firebase) GetAllAccounts() ([]models.Account, error) {
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
		accounts = append(accounts, converter(doc.Data()))
	}

	return accounts, nil
}

func (fb Firebase) GetAccountByID(id string) (models.Account, error) {
	dsnap, err := fb.client.Collection("users").Doc(id).Get(context.Background())
	if err != nil {
		log.Fatalf("Failed to Retrieve ID: %v", err)
		return models.Account{}, err
	}
	return converter(dsnap.Data()), nil
}

func (fb Firebase) GetAccountByEmail(email string) (models.Account, error) {
	iter := fb.client.Collection("users").Where("email", "==", email).Limit(1).Documents(context.Background())

	doc, err := iter.Next()
	if err != nil {
		log.Fatal(err)
	}
	return converter(doc.Data()), err

}

func (fb Firebase) UpdateAccount(id string, updates map[string]interface{}) error {
	_, err := fb.client.Collection("users").Doc(id).Set(context.Background(), updates, firestore.MergeAll)
	if err != nil {
		log.Fatalf("Failed to Update User: %v", err)
	}
	return err
}

func (fb Firebase) RemoveAccount(id string) error {
	_, err := fb.client.Collection("users").Doc(id).Delete(context.Background())
	if err != nil {
		log.Fatalf("Failed to Remove User: %v", err)
		return err
	}
	return nil
}

func converter(data map[string]interface{}) models.Account {

	return models.Account{
		ID:          data["id"].(string),
		AccountType: data["account_type"].(int64),
		Email:       data["email"].(string),
		Password:    data["password"].(string),
		FirstName:   data["first_name"].(string),
		LastName:    data["last_name"].(string),
		Bio:         data["bio"].(string),
		Industry:    data["industry"].(string),
		Education:   data["education"].(string),
		State:       data["state"].(string),
		Reputation:  data["reputation"].(int64),
		Interests:   data["interests"].([]interface{}),
	}
}
