package notification

import "github.com/mahdimehrabi/account-notification/notification/internal/entity"

type Notification interface {
	Save(transaction *entity.Transaction) error
}
