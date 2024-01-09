package entity

import "time"

type Transaction struct {
	From      *Account
	To        *Account
	Amount    float64
	CreatedAt int64
}

func NewTransaction(from, to *Account, amount float64) *Transaction {
	return &Transaction{
		from, to, amount, time.Now().Unix(),
	}
}
