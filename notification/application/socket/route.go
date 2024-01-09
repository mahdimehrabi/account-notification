package socket

import (
	"net"

	generalController "github.com/mahdimehrabi/account-notification/notification/application/socket/controller/general"
	"github.com/mahdimehrabi/account-notification/notification/application/socket/controller/notification"
	"github.com/mahdimehrabi/account-notification/notification/application/socket/dto"
)

func HandleRoute(conn net.Conn, req dto.Request) {
	switch req.Method {
	case "save_notification":
		msg := notification.NewNotification() // transient dependency injection one instance per endpoint call to get best performance
		msg.Save(conn, req)
	case "ping":
		g := generalController.NewGeneral() // transient dependency injection one instance per endpoint call to get best performance
		g.Ping(conn, req)
	default:
		g := generalController.NewGeneral() // transient dependency injection one instance per endpoint call to get best performance
		g.NotDefined(conn, req)
	}
}
