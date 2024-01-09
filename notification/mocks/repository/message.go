package repository

import (
	"github.com/mahdimehrabi/account-notification/notification/internal/entity"
	"github.com/stretchr/testify/mock"
)

// MockMessage is a mock implementation of the Notification interface
type MockMessage struct {
	mock.Mock
}

// Save is a mock implementation for the Save method
func (m *MockMessage) Save(message *entity.Notification) error {
	args := m.Called(message)
	return args.Error(0)
}
