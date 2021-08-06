package command

import (
	"sword-health/notification/application/data_model"
	"sword-health/notification/application/dto"
)

type Handler interface {
	Exec(cmd string, body []byte)
	Read() Read
	Write() Write
}

type Read interface {
	FindOne(userLoggedId int, id int) (notification *data_model.Notification, err error)
	ListNotifications(userLoggedId int, fromId int, limit int) (notifications []*data_model.Notification, err error)
}

type Write interface {
	Create(notificationDTO dto.CreateNotificationDTO) (notification *data_model.Notification, err error)
	MarkAsRead(userLoggedId int, id int) (err error)
}
