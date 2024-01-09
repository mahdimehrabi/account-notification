package dto

import "github.com/mahdimehrabi/account-notification/notification/internal/entity"

type Transaction struct {
	From      *entity.Account `json:"from"`
	To        *entity.Account `json:"to"`
	Amount    float64         `json:"amount"`
	CreatedAt int64           `json:"createdAt"`
}

func (m Transaction) ToModel() *entity.Transaction {
	return &entity.Transaction{
		From:      m.From,
		To:        m.To,
		CreatedAt: m.CreatedAt,
		Amount:    m.Amount,
	}
}
