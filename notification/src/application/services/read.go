package services

import (
	"sword-health/task/application/data_model"
	"sword-health/task/application/repositories"

	"gorm.io/gorm"
)

type ReadService struct {
	notificationRepository *repositories.NotificationRepository
}

func (ReadService) New(repository *repositories.NotificationRepository) *ReadService {
	return &ReadService{
		notificationRepository: repository,
	}
}

func (us *ReadService) ListNotifications(params *dto.FindNotification) []*data_model.Notification {

	notificationDataModel := &data_model.Notification{
		OwnerId: params.OwnerId,
	}

	if params.Role != "manager" {
		notificationDataModel.OwnerId = params.UserLoggedId
	}

	return us.notificationRepository.ListNotifications(notificationDataModel, params.Limit)
}

func (us *ReadService) FindOne(params *dto.FindNotification) (notification *data_model.Notification, err error) {

	model, err := us.notificationRepository.FindOne(
		params.Id,
		params.UserLoggedId,
		params.Role == "manager",
	)

	if err != nil {
		return &data_model.Notification{}, gorm.ErrRecordNotFound
	}

	return model.GetDataModel(), err
}
