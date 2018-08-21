package controllers

import (
	"github.com/ecclesia-dev/account-service/data"
	"github.com/ecclesia-dev/account-service/models"
)

type AccountController struct {
	access data.AccountAccess
}

func NewAccountController() AccountController {
	return AccountController{access: data.NewFirebase()}
}

// TODO:
// 	* Add validation to all inputs here
//	* Handel errors and error logging

func (acct *AccountController) CreateAccount(acctInfo models.Account) error {
	_, err := acct.access.CreateAccount(acctInfo)
	return err
}

func (acct *AccountController) GetAllAccounts() ([]models.Account, error) {
	return acct.access.GetAllAccounts()
}

func (acct *AccountController) GetAccountByID(id string) (models.Account, error) {
	return acct.access.GetAccountByID(id)
}

func (acct *AccountController) GetAccountByEmail(email string) (models.Account, error) {
	return acct.access.GetAccountByEmail(email)
}

func (acct *AccountController) UpdateAccount(id string, data map[string]interface{}) error {
	return acct.access.UpdateAccount(id, data)
}

func (acct *AccountController) RemoveAccount(id string) error {
	return acct.access.RemoveAccount(id)
}

func (acct *AccountController) GetToken(id string) (string, error) {
	return acct.access.GetToken(id)
}

func (acct *AccountController) Logout(id string) error {
	return acct.access.Logout(id)
}

func (acct *AccountController) Login(username string, password string) (string, error) {
	return acct.access.Login(username, password)
}
