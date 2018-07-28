package data

import "github.com/ecclesia-dev/account-service/models"

type DataAccess interface {
	// Not entirely sure what to the return type is going to be here
	CreateUser([]string) error
	FindAllUsers() ([]models.Account, error)
	FindUserByID(string) (models.Account, error)
	FindUserByEmail(string) (models.Account, error)
	UpdateUser(string) error
	RemoveUser(string) error
}
