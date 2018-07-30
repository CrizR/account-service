package data

import (
	"github.com/ecclesia-dev/account-service/models"
)

type DataAccess interface {
	// Not entirely sure what to the return type is going to be here
	CreateAccount(models.Account) error
	FindAllAccounts() ([]models.Account, error)
	FindAccountByID(string) (models.Account, error)
	FindAccountByEmail(string) (models.Account, error)
	UpdateAccount(string, map[string]interface{}) error
	RemoveAccount(string) error
}
