package models

// AccountType represents authorization level of an account.
type AccountType int64

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

func NewAccount(data map[string]interface{}) Account {
	if data == nil {
		return Account{}
	}

	return Account{
		ID:          data["id"].(string),
		AccountType: data["account_type"].(int64),
		Email:       data["email"].(string),
		Password:    data["password"].(string),
		FirstName:   data["first_name"].(string),
		LastName:    data["last_name"].(string),
		Bio:         data["bio"].(string),
		Industry:    data["industry"].(string),
		Education:   data["education"].(string),
		State:       data["state"].(string),
		Reputation:  data["reputation"].(int64),
		Interests:   data["interests"].([]interface{}),
	}
}

func (ac Account) Map() map[string]interface{} {
	return map[string]interface{}{
		"ID":          ac.ID,
		"AccountType": ac.AccountType,
		"Email":       ac.Email,
		"Password":    ac.Password,
		"FirstName":   ac.FirstName,
		"LastName":    ac.LastName,
		"Bio":         ac.Bio,
		"Industry":    ac.Industry,
		"Education":   ac.Education,
		"State":       ac.State,
		"Reputation":  ac.Reputation,
		"Interests":   ac.Interests,
	}
}
