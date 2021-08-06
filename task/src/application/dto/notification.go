package dto

type CreateNotificationDTO struct {
	NotificationType string `json:"type"`
	Content          string `json:"content"`
	FromId           int    `json:"from_id"`
}
