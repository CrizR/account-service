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
	ID         int
	AccoutType AccountType
	Email      string
	Password   string
	FirstName  string
	LastName   string
	Bio        string
	Profession string
	Education  string
	State      string
	Reputation int
	Interests  []string
}
