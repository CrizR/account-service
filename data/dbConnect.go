package data

import (
	//"fmt"

	//"golang.org/x/net/context"

	firebase "firebase.google.com/go"
	//"firebase.google.com/go/auth"

	//option "google.golang.org/api/option"
)

type DB struct {
	ecclesiaDb *firebase.App
}

//The fuck, why is this not working >
opt, err := option.WithCredentialsFile("keys/ecclesia-firebase-key.json")

app, err := firebase.NewApp(context.Background(), nil, opt)
if err != nil {
log.Fatalf("error initializing app: %v\n", err)
}

func CreateUser([]string) (result string, err error) {

	return result, err

}

func FindAllUsers(string) (result string, err error) {

	return result, err
}

func FindUserById(string) (result string, err error) {

	return result, err

}

func FindUserByEmail(string) (result string, err error) {

	return result, err

}

func UpdateUser(string) (result string, err error) {

	return result, err

}
func RemoveUser(string) (result string, err error) {

	return result, err

}