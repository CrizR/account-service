package models

// AccountType represents authorization level of an account.
type AccountType int

const (
	Admin AccountType = iota
	Standard
	Advanced
)

// Account represents details for a particular user account.
type Account struct {
	ID          int
	AccountType AccountType
	Email       string
	Password    string
	FirstName   string
	LastName    string
	Bio         string
	Industry    string
	Education   string
	State       string
	Reputation  int
	Interests   []string
}