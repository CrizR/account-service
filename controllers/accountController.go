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
	return acct.data.CreateAccount(acctInfo)
}

func (acct *AccountController) GetAllAccounts() ([]models.Account, error) {
	return acct.data.GetAllAccounts()
}

func (acct *AccountController) GetAccountByID(id string) (models.Account, error) {
	return acct.data.GetAccountByID(id)
}

func (acct *AccountController) GetAccountByEmail(email string) (models.Account, error) {
	return acct.data.GetAccountByEmail(email)
}

func (acct *AccountController) UpdateAccount(id string, data map[string]interface{}) error {
	return acct.data.UpdateAccount(id, data)
}

func (acct *AccountController) RemoveAccount(id string) error {
	return acct.data.RemoveAccount(id)
}
