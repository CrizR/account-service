package data

type DataAccess interface {
	// Not entirely sure what to the return type is going to be here
	CreateUser([]string) (string, error)
	FindAllUsers(string) (string, error)
	FindUserById(string) (string, error)
	FindUserByEmail(string) (string, error)
	UpdateUser(string) (string, error)
	RemoveUser(string) (string, error)
}
