package models

// AccountType represents authorization level of an account.
type AccountType int

const (
	AdminAccount AccountType = iota
	StandardAccount
	AdvancedAccount
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
