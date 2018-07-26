package data

type DataAccess interface {
	// Not entirely sure what to the return type is going to be here
	CreateUser([]string) (result string, err error)
	FindAllUsers(string) (result string, err error)
	FindUserById(string) (result string, err error)
	FindUserByEmail(string) (result string, err error)
	UpdateUser(string) (result string, err error)
	RemoveUser(string) (result string, err error)
}