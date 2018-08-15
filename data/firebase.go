package data

import (
	"context"

	"cloud.google.com/go/firestore"
	"firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/ecclesia-dev/account-service/models"
	log "github.com/sirupsen/logrus"
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
	// TODO: set FIREBASE_CONFIG as an envornment variable so config can be passed in as nil.
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Failed to initialize firebase app: %v\n", err)
	}
	client, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	auth, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("Failed to initialize Auth: %v\n", err)
	}

	log.Info("Firebase connection initialized.")
	return Firebase{app: app, client: client, auth: auth}
}

func (fb Firebase) CreateAccount(account models.Account) (string, error) {
	var token string
	params := (&auth.UserToCreate{}).Email(account.Email).Password(account.Password)
	account.Password = ""
	user, err := fb.auth.CreateUser(context.Background(), params)
	if err != nil {
		log.Errorf("Failed to create authentificaiton detials new account: %v", err)
		return token, err
	}
	user_data := account.Map()
	delete(user_data, "Password")
	delete(user_data, "ID")
	_, err = fb.client.Collection("users").Doc(user.UID).Set(context.Background(), user_data, firestore.MergeAll)
	if err != nil {
		log.Errorf("Failed to add account to firestore: %v", err)
		return token, err
	}
	token, err = fb.auth.CustomToken(context.Background(), user.UID)
	if err != nil {
		log.Errorf("Failed to create token for account: %v", err)
		return token, err
	}
	return token, err
}

func (fb Firebase) GetAllAccounts() ([]models.Account, error) {
	var accounts []models.Account
	iter := fb.client.Collection("users").Documents(context.Background())
	var err error
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Warn("Failed to iterate through accounts")
		}
		accounts = append(accounts, models.NewAccount(doc.Data()))
	}
	return accounts, err
}

func (fb Firebase) GetAccountByID(id string) (models.Account, error) {
	var data map[string]interface{}
	var dsnap *firestore.DocumentSnapshot
	dsnap, err := fb.client.Collection("users").Doc(id).Get(context.Background())
	if err != nil {
		log.Warn(err)
	}
	data = dsnap.Data()
	return models.NewAccount(data), err
}

func (fb Firebase) GetAccountByEmail(email string) (models.Account, error) {
	var data map[string]interface{}
	iter := fb.client.Collection("users").Where("email", "==", email).Limit(1).Documents(context.Background())
	doc, err := iter.Next()
	if err != nil {
		log.Warn(err)
	} else {
		data = doc.Data()
	}
	return models.NewAccount(data), err
}

func (fb Firebase) RemoveAccount(id string) error {
	err := fb.auth.DeleteUser(context.Background(), id)
	_, err = fb.client.Collection("users").Doc(id).Delete(context.Background())
	if err != nil {
		log.Errorf("Failed to Remove User: %v", err)
	}
	return err
}

func (fb Firebase) UpdateAccount(id string, updates map[string]interface{}) error {

	if email, ok := updates["Email"].(string); ok {
		fb.changeEmail(id, email)
	}
	if password, ok := updates["Password"].(string); ok {
		fb.changePassword(id, password)
		delete(updates, "Password")
	}
	_, err := fb.client.Collection("users").Doc(id).Set(context.Background(), updates, firestore.MergeAll)
	if err != nil {
		log.Errorf("Failed to Update User: %v", err)
	}

	return err
}

func (fb Firebase) changeEmail(id string, newEmail string) error {
	params := (&auth.UserToUpdate{}).Email(newEmail)
	_, err := fb.auth.UpdateUser(context.Background(), id, params)
	if err != nil {
		log.Errorf("Failed to change email: %v", err)
	}
	return err
}

func (fb Firebase) changePassword(id string, newPassword string) error {
	params := (&auth.UserToUpdate{}).
		Password(newPassword)
	_, err := fb.auth.UpdateUser(context.Background(), id, params)
	if err != nil {
		log.Errorf("Failed to change password: %v", err)
	}
	return err
}
