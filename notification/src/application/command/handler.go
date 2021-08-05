package command

import (
	"encoding/json"
	"log"
	"sword-health/notification/application/dto"
	"sword-health/notification/application/services"
)

type NotificationHandler struct {
	notificationWriteService *services.WriteService
	notificationReadService  *services.ReadService
}

func (NotificationHandler) New(
	notificationWriteService *services.WriteService,
	notificationReadService *services.ReadService,
) *NotificationHandler {
	return &NotificationHandler{
		notificationWriteService: notificationWriteService,
		notificationReadService:  notificationReadService,
	}

}

func (uh *NotificationHandler) Exec(cmd string, body []byte) {
	switch cmd {
	case "notification.create":
		notificationDTO := dto.NotificationCreateDTO{}

		if err := json.Unmarshal(body, &notificationDTO); err == nil {
			uh.notificationWriteService.Create(notificationDTO)
		}

	default:
		log.Println("no implemented.")
	}
}

func (h *NotificationHandler) Read() *services.ReadService {

	return h.notificationReadService
}

func (h *NotificationHandler) Write() *services.WriteService {

	return h.notificationWriteService
}
