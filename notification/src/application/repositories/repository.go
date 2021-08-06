package repositories

import (
	"sword-health/notification/application/data_model"
	"sword-health/notification/domain"
)

type Repository interface {
	Add(model *domain.NotificationModel) (notification *data_model.Notification, err error)
	FindOne(id int) (notification *domain.NotificationModel, err error)
	ListNotifications(fromId int, limit int) (notifications []*domain.NotificationModel, err error)
	Update(model *domain.NotificationModel) (notification *data_model.Notification, err error)
}
