package account

import (
	"github.com/go-playground/validator"
	"github.com/mahdimehrabi/account-notification/account/internal/entity"
	"time"
)

type CreateDTO struct {
	AccountID int64   `json:"accountID" validate:"required"`
	Balance   float64 `json:"balance" validate:"required"`
}

func (c CreateDTO) ToModel() *entity.Account {
	return entity.NewAccount(c.AccountID, c.Balance)
}

func (c CreateDTO) Validate() error {
	//validating is concern of DTO not controller (solid (Single Responsibility Principle))
	validate := validator.New()
	return validate.Struct(c)
}

type TransactionDTO struct {
	FromID   int64   `json:"fromID" validate:"required"`
	ToID     int64   `json:"toID" validate:"required"`
	Amount   float64 `json:"amount" validate:"required"`
	CreateAt int64   `json:"-"`
}

func (c TransactionDTO) Validate() error {
	//validating is concern of DTO not controller (solid (Single Responsibility Principle))
	validate := validator.New()
	return validate.Struct(c)
}

func (c TransactionDTO) ToModel() *entity.Transaction {
	return &entity.Transaction{
		From:      entity.NewAccount(c.FromID, 0),
		To:        entity.NewAccount(c.ToID, 0),
		Amount:    c.Amount,
		CreatedAt: time.Now().Unix(),
	}
}
