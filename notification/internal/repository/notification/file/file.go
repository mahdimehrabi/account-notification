package message

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/mahdimehrabi/account-notification/notification/internal/entity"
	"github.com/mahdimehrabi/account-notification/notification/internal/repository/notification"
	"github.com/rs/zerolog/log"
)

const (
	workerCount          = 50
	saveDeadlineDuration = 5 * time.Second
)

var ErrResourceNotAvailable = errors.New("resource is not available")

type file struct {
	queue chan entity.Transaction
}

func NewFile() notification.Notification {
	d := &file{
		queue: make(chan entity.Transaction, 10000),
	}
	go d.SaveQueue()
	return d
}

func (r *file) SaveQueue() {
	for i := 0; i < workerCount; i++ {
		go r.savingWorker()
	}
}

func (r *file) savingWorker() {
	for {
		transaction := <-r.queue
		id := uuid.New().String()

		deadline := time.NewTicker(saveDeadlineDuration)
		done := make(chan bool)
		go func(ch chan bool) {
			if err := r.saveToFile(transaction); err != nil {
				log.Printf("failed to save notification %s trying again,err:%s", id, err.Error())
				return //if this happen transaction will go end of queue for trying again
			}
			done <- true
		}(done)
		select {
		case <-done:
			log.Printf("notification with timestamp %d saved successfully🥳 \n", transaction.CreatedAt)
		case <-deadline.C: // deadline exceeded
			time.Sleep(1 * time.Microsecond) // socket resend cool down
			r.queue <- transaction
		}
	}
}

func (r *file) saveToFile(transaction entity.Transaction) error {
	// Get today's date for the filename
	today := time.Now().Format("2006-01-02")
	directoryPath := fmt.Sprintf("logs/%d", transaction.From.AccountID)
	filename := fmt.Sprintf("%s/%s.json", directoryPath, today)

	// Create the directory if it doesn't exist
	if _, err := os.Stat(directoryPath); os.IsNotExist(err) {
		if err := os.MkdirAll(directoryPath, 0755); err != nil {
			return err
		}
	}

	// Open or create the file
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Convert transaction to JSON
	transactionJSON, err := json.Marshal(transaction)
	if err != nil {
		return err
	}

	// Append the JSON data to the file
	if _, err := file.WriteString(string(transactionJSON) + "\n"); err != nil {
		return err
	}

	return nil
}

func (r *file) Save(msg *entity.Transaction) error {
	if len(r.queue) >= 10000 {
		return ErrResourceNotAvailable
	}
	r.queue <- *msg
	return nil
}
