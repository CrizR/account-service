package data

import (
	"github.com/ecclesia-dev/account-service/models"
)

type DataAccess interface {
	CreateAccount(models.Account) error
	GetAllAccounts() ([]models.Account, error)
	GetAccountByID(id string) (models.Account, error)
	GetAccountByEmail(id string, email string) (models.Account, error)
	UpdateAccount(id string, updates map[string]interface{}) error
	changeEmail(id string, newEmail string) error
	changePassword(id string, newPassword string) error
	RemoveAccount(id string) error
}
