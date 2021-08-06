package command

import (
	"encoding/json"
	"log"
	"sword-health/notification/application/dto"
)

type NotificationHandler struct {
	notificationWriteService Write
	notificationReadService  Read
}

func (NotificationHandler) New(
	notificationWriteService Write,
	notificationReadService Read,
) *NotificationHandler {
	return &NotificationHandler{
		notificationWriteService: notificationWriteService,
		notificationReadService:  notificationReadService,
	}

}

func (uh *NotificationHandler) Exec(cmd string, body []byte) {

	switch cmd {
	case "notification.create":
		notificationDTO := dto.CreateNotificationDTO{}

		if err := json.Unmarshal(body, &notificationDTO); err == nil {
			uh.notificationWriteService.Create(notificationDTO)
		}

	default:
		log.Println("no implemented.")
	}
}

func (h *NotificationHandler) Read() Read {

	return h.notificationReadService
}

func (h *NotificationHandler) Write() Write {

	return h.notificationWriteService
}
