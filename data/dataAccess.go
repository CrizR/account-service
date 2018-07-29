package data


import (
	"github.com/ecclesia-dev/account-service/models"
)
type DataAccess interface {
	// Not entirely sure what to the return type is going to be here
	CreateUser(models.Account) (error)
	FindAllUsers() ([]models.Account, error)
	FindUserById(string) (models.Account, error)
	FindUserByEmail(string) (models.Account, error)
	UpdateUser(string, map[string]interface{}) (error)
	RemoveUser(string) (error)

}
