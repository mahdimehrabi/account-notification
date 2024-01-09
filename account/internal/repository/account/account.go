package account

import (
	"errors"
	"github.com/mahdimehrabi/account-notification/account/internal/entity"
)

var ErrNotFound = errors.New("account not found")
var ErrInsufficientAmount = errors.New("insufficient amount")

// I used contemplate methods and not transaction because it makes my code more clean
// for executing multiple queries inside of a repository method I rather use transaction not for roleback from service
type Account interface {
	Create(account *entity.Account) error
	List() ([]entity.Account, error)
	Send(transaction *entity.Transaction) error
	SendContemplate(transaction *entity.Transaction) error
}
