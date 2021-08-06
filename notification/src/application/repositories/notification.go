package repositories

import (
	"sword-health/notification/application/data_model"
	"sword-health/notification/domain"

	"gorm.io/gorm"
)

type NotificationRepository struct {
	db *gorm.DB
}

func (NotificationRepository) New(
	db *gorm.DB,
) Repository {
	return &NotificationRepository{
		db: db,
	}
}

func (r *NotificationRepository) Add(model *domain.NotificationModel) (notification *data_model.Notification, err error) {

	dataModel := model.GetDataModel()

	err = r.db.Save(dataModel).Error

	return dataModel, err
}

func (r *NotificationRepository) FindOne(id int) (notification *domain.NotificationModel, err error) {

	var dataModel = data_model.Notification{}

	err = r.db.Where(
		&data_model.Notification{
			ID: uint(id),
		},
	).Take(&dataModel).Error

	if err != nil {
		return &domain.NotificationModel{}, err
	}

	return (domain.NotificationModel{}).Load(&dataModel), err

}

func (r *NotificationRepository) ListNotifications(fromId int, limit int) (notifications []*domain.NotificationModel, err error) {
	dataModels := []*data_model.Notification{}

	err = r.db.Where(data_model.Notification{
		FromId: fromId,
	}).
		Find(&dataModels).
		Limit(limit).
		Error

	for _, data := range dataModels {
		model := (domain.NotificationModel{}).Load(data)
		notifications = append(notifications, model)
	}

	return notifications, err
}

func (r *NotificationRepository) Update(model *domain.NotificationModel) (notification *data_model.Notification, err error) {

	dataModel := model.GetDataModel()
	err = r.db.Save(&dataModel).Error

	return dataModel, err
}
