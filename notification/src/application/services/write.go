package services

import (
	"fmt"
	"sword-health/notification/application/data_model"
	"sword-health/notification/application/dto"
	"sword-health/notification/application/repositories"
	"sword-health/notification/domain"
	grpc_user "sword-health/notification/infra/grpc/client/user"

	"github.com/go-redis/redis"
)

type WriteService struct {
	redis                  *redis.Client
	userClient             *grpc_user.UserClient
	notificationRepository repositories.Repository
}

func (WriteService) New(
	redis *redis.Client,
	userClient *grpc_user.UserClient,
	repository repositories.Repository,
) *WriteService {
	return &WriteService{
		redis:                  redis,
		userClient:             userClient,
		notificationRepository: repository,
	}
}

func (us *WriteService) Create(notificationDTO dto.CreateNotificationDTO) (notification *data_model.Notification, err error) {

	newNotification, err := domain.Create(
		notificationDTO.NotificationType,
		notificationDTO.Content,
		notificationDTO.FromId,
	)

	notification, err = us.notificationRepository.Add(&newNotification)
	
	go us.ClearCache(notificationDTO.FromId, int(notification.ID))
	
	return notification, err
}

func (us *WriteService) MarkAsRead(userLoggedId int, id int) (err error) {

	user, err := us.userClient.Get(userLoggedId)

	if err != nil {
		return err
	}

	notification, err := us.notificationRepository.FindOne(id)

	if err != nil {
		return err
	}

	err = notification.MarkAsRead(user.GetIsManager())

	if err != nil {
		return err
	}

	dataModel, err := us.notificationRepository.Update(notification)

	go us.ClearCache(int(user.GetId()), int(dataModel.ID))

	if err != nil {
		return err
	}

	return err
}

func (us *WriteService) ClearCache(userId int, notificationId int) {

	keys := []string{
		fmt.Sprintf("notification.%d", notificationId),
		fmt.Sprintf("notification.list.user.%d", userId),
	}

	us.redis.Del(keys...)
}
