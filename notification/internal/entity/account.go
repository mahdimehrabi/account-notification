package entity

type Account struct {
	AccountID int64
	Balance   float64
}

func NewAccount(accountID int64, balance float64) *Account {
	return &Account{AccountID: accountID, Balance: balance}
}
