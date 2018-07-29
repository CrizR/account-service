package controllers

import (
	"github.com/ecclesia-dev/account-service/models"
)

type ControllerAccess interface {
	CreateAccount([]string) error
	FindAllAccount(string) ([]models.Account, error)
	FindAccountByID(string) (models.Account, error)
	FindAccountByEmail(string) (models.Account, error)
	UpdateAccount(string) error
	RemoveAccount(string) error
}
