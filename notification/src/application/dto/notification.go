package dto

type CreateNotificationDTO struct {
	ID               int    `gorm:"primary_key;auto_increment;not_null"`
	NotificationType string `json:"type"`
	Status           string `json:"status"`
	Content          string `json:"content"`
	FromId           int    `json:"from_id"`
}

func CreateNotification(
	NotificationType string,
	Status string,
	Content string,
	FromId int,
) *CreateNotificationDTO {
	return &CreateNotificationDTO{
		NotificationType: NotificationType,
		Status:           Status,
		Content:          Content,
		FromId:           FromId,
	}
}
