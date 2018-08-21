package controllers

import (
	"github.com/ecclesia-dev/account-service/models"
)

type ControllerAccess interface {
	CreateAccount([]string) error
	GetAllAccount(string) ([]models.Account, error)
	GetAccountByID(string) (models.Account, error)
	GetAccountByEmail(string) (models.Account, error)
	UpdateAccount(string, map[string]interface{}) error
	RemoveAccount(string) error
	GetToken(id string) (string, error)
	Logout(ID string) (error)
	Login(username string, password string) (string, error)
}
