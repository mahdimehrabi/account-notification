package notification

import (
	"net"

	"github.com/mahdimehrabi/account-notification/notification/application/socket/dto"
	"github.com/mahdimehrabi/account-notification/notification/application/socket/response"
	"github.com/mahdimehrabi/account-notification/notification/internal/service/notification"
	"github.com/mitchellh/mapstructure"
)

type Notification struct {
	messageService *notification.Notification
}

func NewNotification() *Notification {
	return &Notification{
		notification.NewNotification(),
	}
}

func (g Notification) Save(conn net.Conn, req dto.Request) {
	msgReq := dto.Transaction{}
	if err := mapstructure.Decode(req.Data, &msgReq); err != nil {
		response.BadRequestErrorResponse(conn, req.ID)
		return
	}
	msgEnt := msgReq.ToModel()
	if err := g.messageService.Save(msgEnt); err != nil {
		response.InternalErrorResponse(conn, req.ID)
		return
	}
	response.SuccessResponse(conn, nil, req.ID, "notification saved successfully")
}
