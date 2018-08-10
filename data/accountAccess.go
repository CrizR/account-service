package data

import (
	"github.com/ecclesia-dev/account-service/models"
)

type AccountAccess interface {
	CreateAccount(account models.Account) error
	GetAllAccounts() ([]models.Account, error)
	GetAccountByID(id string) (models.Account, error)
	GetAccountByEmail(email string) (models.Account, error)
	UpdateAccount(id string, updates map[string]interface{}) error
	RemoveAccount(id string) error
	GetToken(ID string) (string, error)
	Login(username string, password string) (string, error)
	Logout(ID string) (error)
}
