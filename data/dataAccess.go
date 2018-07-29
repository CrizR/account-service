package data


import (
	"github.com/ecclesia-dev/account-service/models"
)
type DataAccess interface {
	CreateUser(models.Account) (error)
	FindAllUsers() ([]models.Account, error)
	FindUserByID(string) (models.Account, error)
	FindUserByEmail(string) (models.Account, error)
	UpdateUser(string, map[string]interface{}) (error)
	RemoveUser(string) (error)
}
