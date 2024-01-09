package dpi

import (
	"github.com/mahdimehrabi/account-notification/account/external/notification"
	infrastructures "github.com/mahdimehrabi/account-notification/account/internal/infrastructure"
	accountRepo "github.com/mahdimehrabi/account-notification/account/internal/repository/account"
	"github.com/mahdimehrabi/account-notification/account/internal/repository/account/redis"
)

// this helps increase performance/scalability/reliability because requests will send more parallel and requests
// won't send and relies on only one socket
const socketConnectionCount = 1000

// DPI singleton dependency injection
// for instances that use limited resources
var AppDPI *DPI

// we will define singleton dependencies here
type DPI struct {
	AccountRepo        accountRepo.Account
	NotificationSender *notification.NotificationSender
}

func NewDPI(env *infrastructures.Env) *DPI {
	sockets := make([]*notification.Socket, socketConnectionCount)
	for i := 0; i < socketConnectionCount; i++ {
		sockets[i] = notification.NewSocket(env.NotificationAddr)
	}
	return &DPI{AccountRepo: redis.NewRedis(env.RedisAddr, env.RedisPassword, env.RedisDB),
		NotificationSender: notification.NewNotificationSender(sockets)}
}
