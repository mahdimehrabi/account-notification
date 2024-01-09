package notification

import (
	"errors"
	"github.com/google/uuid"
	"github.com/mahdimehrabi/account-notification/account/external/utils"
	"github.com/mahdimehrabi/account-notification/account/internal/entity"
	"github.com/rs/zerolog/log"
	"time"
)

const (
	workerCount            = 50
	saveNotificationMethod = "save_notification"
	saveDeadlineDuration   = 5 * time.Second
)

var ErrResourceNotAvailable = errors.New("resource is not available")

type NotificationSender struct {
	queue   chan entity.Transaction
	sockets []*Socket
}

func NewNotificationSender(sockets []*Socket) *NotificationSender {
	b := &NotificationSender{
		sockets: sockets,
		queue:   make(chan entity.Transaction, 10000),
	}
	go b.SaveQueue()
	return b
}

func (b NotificationSender) SaveQueue() {
	for i := 0; i < workerCount; i++ {
		go b.savingWorker()
	}
}

func (b NotificationSender) savingWorker() {
	for {
		notification := <-b.queue
		socket := b.sockets[utils.RandomNumber(len(b.sockets)-1)]
		id := uuid.New().String() //we get response base on this
		deadline := time.NewTicker(saveDeadlineDuration)
		done := make(chan bool)
		go func(ch chan bool) {
			if _, err := socket.SendWaitJSON(notification, saveNotificationMethod, id); err != nil {
				log.Printf("failed to save notification %s trying again,err:%s", id, err.Error())
				return //done will not be filled , and deadline will be exceeded
			}
			done <- true
		}(done)
		select {
		case <-done:
			log.Printf("notificaton with timestamp %d saved successfully🥳 \n", notification.CreatedAt)
		case <-deadline.C: //deadline exceeded
			time.Sleep(1 * time.Microsecond)
			b.queue <- notification
		}
	}
}

func (b NotificationSender) Send(transaction entity.Transaction) error {
	if len(b.queue) >= 10000 {
		return ErrResourceNotAvailable
	}
	b.queue <- transaction
	return nil
}
