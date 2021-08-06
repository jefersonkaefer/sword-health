package services

import (
	"encoding/json"
	"fmt"
	"sword-health/notification/application/data_model"
	"sword-health/notification/application/repositories"
	grpc_user "sword-health/notification/infra/grpc/client/user"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type ReadService struct {
	redis                  *redis.Client
	userClient             *grpc_user.UserClient
	notificationRepository repositories.Repository
}

func (ReadService) New(
	redis *redis.Client,
	userClient *grpc_user.UserClient,
	repository repositories.Repository,
) *ReadService {
	return &ReadService{
		redis:                  redis,
		userClient:             userClient,
		notificationRepository: repository,
	}
}

func (us *ReadService) FindOne(userLoggedId int, id int) (notification *data_model.Notification, err error) {
	
	notification = &data_model.Notification{}

	key := fmt.Sprintf("notification.%d", id)

	user, err := us.userClient.Get(userLoggedId)

	if err != nil {
		return notification, err
	}

	data, err := us.redis.Get(key).Bytes()

	if len(data) == 0 {

		if !user.GetIsManager() {
			return notification, gorm.ErrRecordNotFound
		}

		model, err := us.notificationRepository.FindOne(
			id,
		)

		notification = model.GetDataModel()

		data, err = json.Marshal(&notification)

		if err != nil {
			return notification, gorm.ErrInvalidData
		}

		us.redis.Set(key, data, 0)
	}

	err = json.Unmarshal(data, &notification)

	if err != nil {
		return notification, err
	}

	return notification, err
}

func (us *ReadService) ListNotifications(userLoggedId int, fromId int, limit int) (notifications []*data_model.Notification, err error) {

	userLogged, err := us.userClient.Get(userLoggedId)

	if err != nil {
		return notifications, err
	}

	if !userLogged.GetIsManager() {
		fromId = int(userLogged.GetId())
	}

	key := fmt.Sprintf("notification.list.user.%d", userLogged.GetId())

	data, err := us.redis.Get(key).Bytes()

	if len(data) == 0 {
		notificationModels, err := us.notificationRepository.ListNotifications(fromId, limit)

		users := make(map[int]*grpc_user.User)

		for _, notification := range notificationModels {

			user, ok := users[notification.GetFromId()]

			if !ok {
				user, err := us.userClient.Get(notification.GetFromId())
				users[int(user.GetId())] = user

				if err != nil {
					return notifications, err
				}
			}

			dataModel := notification.GetDataModel()

			dataModel.FromFullName = user.GetFullName()

			notifications = append(notifications, dataModel)

			data, err = json.Marshal(notifications)

			if err != nil {
				return notifications, err
			}

		}
		if err == nil {
			us.redis.Set(key, data, 0)
		}
	}
	err = json.Unmarshal(data, &notifications)

	if limit > 0 {
		count := len(notifications)

		if limit > count {
			limit = count
		}
		return notifications[0:limit], err
	}

	return notifications, err
}
