package notification

import (
	"errors"
	"github.com/mahdimehrabi/account-notification/notification/application/socket/dpi"

	"github.com/mahdimehrabi/account-notification/notification/internal/entity"
	"github.com/mahdimehrabi/account-notification/notification/internal/repository/notification"
	"github.com/rs/zerolog/log"
)

var ErrMessageInternal = errors.New("error")

type Notification struct {
	messageRepo notification.Notification
}

func NewNotification() *Notification {
	return &Notification{
		messageRepo: dpi.AppDPI.NotificationRepo,
	}
}

func (m Notification) Save(msg *entity.Transaction) error {
	if err := m.messageRepo.Save(msg); err != nil {
		log.Printf("error happened in sending notification to file: %s", err.Error())
		return ErrMessageInternal
	}
	return nil
}
