package dpi

import (
	infrastructures "github.com/mahdimehrabi/account-notification/notification/internal/infrastructure"
	notificationRepo "github.com/mahdimehrabi/account-notification/notification/internal/repository/notification"
	message "github.com/mahdimehrabi/account-notification/notification/internal/repository/notification/file"
)

// this helps increase performance/scalability/reliability because requests will send more parallel and requests
// won't send and relies on only one socket
const socketConnectionCount = 1000

// DPI singleton dependency injection
// for instances that use limited resources
var AppDPI *DPI

// we will define singleton dependencies here
type DPI struct {
	NotificationRepo notificationRepo.Notification
}

func NewDPI(env *infrastructures.Env) *DPI {

	return &DPI{NotificationRepo: message.NewFile()}
}
