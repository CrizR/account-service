package data

import (
	"context"
	"log"
	"cloud.google.com/go/firestore"
	"firebase.google.com/go"
	"firebase.google.com/go/auth"
	"account-service/models"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type Firebase struct {
	app    *firebase.App
	client *firestore.Client
	auth   *auth.Client
}

func NewFirebase() Firebase {
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
	return Firebase{app: app, client: client, auth: auth}
}

func (fb Firebase) CreateAccount(account models.Account) (string, error) {
	var token string
	params := (&auth.UserToCreate{}).
		Email(account.Email).
		Password(account.Password).
		EmailVerified(true)
	account.Password = ""
	user, err := auth.Client.CreateUser(nil, context.Background(), params)
	if err != nil {
		log.Fatalf("Failed adding user: %v", err)
	}
	account.ID = user.UID
	_, err = fb.client.Collection("users").Doc(user.UID).Set(context.Background(), account, firestore.MergeAll)
	if err != nil {
		log.Fatalf("Failed adding user: %v", err)
	}
	token, err = auth.Client.CustomToken(nil, context.Background(), user.UID)
	if err != nil {
		log.Fatalf("Failed adding user: %v", err)
	}
	return token, err
}

func (fb Firebase) GetAllAccounts(token string) ([]models.Account, error) {
	_, err := fb.auth.VerifyIDToken(context.Background(), token)
	if err != nil {
		log.Fatalf("Failed to iterate: %v", err)
		return nil, err
	}
	var accounts []models.Account
	iter := fb.client.Collection("users").Documents(context.Background())
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		accounts = append(accounts, models.NewAccount(doc.Data()))
	}
	return accounts, err
}

func (fb Firebase) GetAccountByID(token string, id string) (models.Account, error) {
	var data map[string]interface{}
	_, err := fb.auth.VerifyIDToken(context.Background(), token)
	if err != nil {
		log.Fatalf("Failed to iterate: %v", err)
	} else {
		var dsnap *firestore.DocumentSnapshot
		dsnap, err = fb.client.Collection("users").Doc(id).Get(context.Background())
		if err != nil {
			log.Fatalf("Failed to Retrieve ID: %v", err)
		} else {
			data = dsnap.Data()
		}
	}
	return models.NewAccount(data), err
}

func (fb Firebase) GetAccountByEmail(token string, email string) (models.Account, error) {
	var data map[string]interface{}
	_, err := fb.auth.VerifyIDToken(context.Background(), token)
	if err != nil {
		log.Fatalf("Failed to iterate: %v", err)
	} else {
		iter := fb.client.Collection("users").Where("email", "==", email).Limit(1).Documents(context.Background())
		doc, err := iter.Next()
		if err != nil {
			log.Fatal(err)
		} else {
			data = doc.Data()
		}
	}
	return models.NewAccount(data), err
}

func (fb Firebase) UpdateAccount(token string, id string, updates map[string]interface{}) error {
	_, err := fb.auth.VerifyIDToken(context.Background(), token)
	if _, ok := updates["Email"]; ok {
		fb.ChangeEmail(token, id)
	}
	if _, ok := updates["Password"]; ok {
		fb.ChangePassword(token, id)
		delete(updates, "Password")
	}
	if err != nil {
		log.Fatalf("Failed to iterate: %v", err)
	} else {
		_, err = fb.client.Collection("users").Doc(id).Set(context.Background(), updates, firestore.MergeAll)
		if err != nil {
			log.Fatalf("Failed to Update User: %v", err)
		}
	}
	return err
}

func (fb Firebase) RemoveAccount(token string, id string) error {
	_, err := fb.auth.VerifyIDToken(context.Background(), token)
	if err != nil {
		log.Fatalf("Failed to iterate: %v", err)
	} else {
		err = fb.auth.DeleteUser(context.Background(), id)
		_, err = fb.client.Collection("users").Doc(id).Delete(context.Background())
		if err != nil {
			log.Fatalf("Failed to Remove User: %v", err)
		}
	}
	return err
}

func (fb Firebase) ChangeEmail(token string, newEmail string) error {
	jwt, err := fb.auth.VerifyIDToken(context.Background(), token)
	if err != nil {
		log.Fatalf("Failed to iterate: %v", err)
	} else {
		params := (&auth.UserToUpdate{}).
			Email(newEmail)
		_, err = fb.auth.UpdateUser(context.Background(), jwt.UID, params)
		if err != nil {
			log.Fatalf("Failed to Remove User: %v", err)
		}
	}
	return err
}

func (fb Firebase) ChangePassword(token string, newPassword string) error {
	jwt, err := fb.auth.VerifyIDToken(context.Background(), token)
	if err != nil {
		log.Fatalf("Failed to iterate: %v", err)
	} else {
		params := (&auth.UserToUpdate{}).
			Password(newPassword)
		_, err = fb.auth.UpdateUser(context.Background(), jwt.UID, params)
		if err != nil {
			log.Fatalf("Failed to Remove User: %v", err)
		}
	}
	return err
}

