package models

/*
Account represents details for a particular user account.
*/
type Account struct {
	AccoutType string
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
