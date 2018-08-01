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
	ID          string        `json:"id" db:"id"`
	AccountType int64         `json:"account_type" db:"account_type"`
	Email       string        `json:"email" db:"email"`
	Password    string        `json:"password" db:"password"`
	FirstName   string        `json:"first_name" db:"first_name"`
	LastName    string        `json:"last_name" db:"last_name"`
	Bio         string        `json:"bio" db:"bio"`
	Industry    string        `json:"industry" db:"industry"`
	Education   string        `json:"education" db:"education"`
	State       string        `json:"state" db:"state"`
	Reputation  int64         `json:"reputation" db:"reputation"`
	Interests   []interface{} `json:"interests" db:"interests"`
}
