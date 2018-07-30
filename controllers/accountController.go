package controllers

import (
	"github.com/ecclesia-dev/account-service/data"
	"github.com/ecclesia-dev/account-service/models"
)

type AccountController struct {
	data data.DataAccess
}

func NewAccountController() AccountController {
	return AccountController{data: data.NewFirebase()}
}

// TODO:
// 	* Add validation to all inputs here
//	* Handel errors and error logging

func (acct *AccountController) CreateAccount(acctInfo models.Account) error {
	return acct.data.CreateUser(acctInfo)
}

func (acct *AccountController) FindAllAccounts() ([]models.Account, error) {
	return acct.data.FindAllUsers()
}

func (acct *AccountController) FindAccountByID(id string) (models.Account, error) {
	return acct.data.FindUserByID(id)
}

func (acct *AccountController) FindAccountByEmail(email string) (models.Account, error) {
	return acct.data.FindUserByEmail(email)
}

func (acct *AccountController) UpdateAccount(id string, data map[string]interface{}) error {
	return acct.data.UpdateUser(id, data)
}

func (acct *AccountController) RemoveAccount(id string) error {
	return acct.data.RemoveUser(id)
}
