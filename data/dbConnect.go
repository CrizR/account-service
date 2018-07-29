package data

import (
	"context"
	"log"

	"firebase.google.com/go"
	"github.com/ecclesia-dev/account-service/models"
	"google.golang.org/api/option"
	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type Firebase struct {
	app    *firebase.App
	client *firestore.Client
}

func NewFirebase() DataAccess {
	opt := option.WithCredentialsFile("keys/ecclesia-firebase-key.json")
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

	return Firebase{app: app, client: client}
}

func (fb Firebase) CreateUser(user models.Account) (error) {
	var err error
	_, _, err = fb.client.Collection("users").Add(context.Background(), user)
	if err != nil {
		log.Fatalf("Failed adding user: %v", err)
	}
	return err
}

func (fb Firebase) FindAllUsers() ([]models.Account, error) {
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

func (fb Firebase) FindUserByID(id string) (models.Account, error) {
	dsnap, err := fb.client.Collection("users").Doc(id).Get(context.Background())
	if err != nil {
		log.Fatalf("Failed to Retrieve ID: %v", err)
		return models.Account{}, err
	}
	return converter(dsnap.Data()), nil
}

func (fb Firebase) FindUserByEmail(email string) (models.Account, error) {
	iter := fb.client.Collection("users").OrderBy("email", firestore.Asc).Where("email", "=", email).Limit(1).Documents(context.Background())
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
			return models.Account{}, err
		}
		return converter(doc.Data()), nil
	}
	return models.Account{}, nil
}

func (fb Firebase) UpdateUser(id string, updates map[string]interface{}) (error) {
	_, err := fb.client.Collection("users").Doc(id).Set(context.Background(), updates, firestore.MergeAll)
	if err != nil {
		log.Fatalf("Failed to Update User: %v", err)
		return err
	}
	return nil
}

func (fb Firebase) RemoveUser(id string) (error) {
	_, err := fb.client.Collection("users").Doc(id).Delete(context.Background())
	if err != nil {
		log.Fatalf("Failed to Remove User: %v", err)
		return err
	}
	return nil
}

func converter(data map[string]interface{}) (models.Account) {
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
