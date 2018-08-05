package data

import (
	"github.com/ecclesia-dev/account-service/models"
)

type DataAccess interface {
	CreateAccount(models.Account) error
	GetAllAccounts(string) ([]models.Account, error)
	GetAccountByID(string, string) (models.Account, error)
	GetAccountByEmail(string, string) (models.Account, error)
	UpdateAccount(string, string, map[string]interface{}) error
	ChangeEmail(string) error
	ChangePassword(string) error
	RemoveAccount(string) error
}
