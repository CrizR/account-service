package data

import (
	"log"
)

type DataAccess interface {
	CreateUser([]string) (result []dbModel.SkuData, err error)
	FindAllUsers(string) (result []dbModel.OraResponse, err error)
	FindUserById(string) (result []dbModel.OraResponse, err error)
	FindUserByEmail(string) (result []dbModel.OraResponse, err error)
	UpdateUser(string) (result []dbModel.OraResponse, err error)
	RemoveUser(string) (result []dbModel.OraResponse, err error)
}