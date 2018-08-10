package data

import (
	"github.com/ecclesia-dev/account-service/models"
)

type DataAccess interface {
	CreateAccount(account models.Account) (string, error)
	GetAllAccounts() ([]models.Account, error)
	GetAccountByID(id string) (models.Account, error)
	GetAccountByEmail(email string) (models.Account, error)
	UpdateAccount(id string, updates map[string]interface{}) error
	RemoveAccount(id string) error
}
