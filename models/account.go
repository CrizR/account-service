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
	ID          int         `json:"id"`
	AccountType AccountType `json:"account_type"`
	Email       string      `json:"email"`
	Password    string      `json:"password"`
	FirstName   string      `json:"first_name"`
	LastName    string      `json:"last_name"`
	Bio         string      `json:"bio"`
	Industry    string      `json:"industry"`
	Education   string      `json:"education"`
	State       string      `json:"state"`
	Reputation  int         `json:"reputation"`
	Interests   []string    `json:"interests"`
}
