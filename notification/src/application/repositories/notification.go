package repositories

import (
	"encoding/json"
	"fmt"
	"sword-health/task/application/data_model"
	"sword-health/task/domain"
	"time"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type NotificationRepository struct {
	redis *redis.Client
	db    *gorm.DB
}

func (NotificationRepository) New(
	redis *redis.Client,
	db *gorm.DB,
) *NotificationRepository {
	return &NotificationRepository{
		redis: redis,
		db:    db,
	}
}

func (r *NotificationRepository) Save(model *domain.NotificationModel) (notification *data_model.Notification, err error) {
	dataModel := model.GetDataModel()
	ttl, _ := time.ParseDuration("300s")

	key := fmt.Sprintf("notification.%d", dataModel)

	dataJson, err := json.Marshal(dataModel)

	if err == nil {
		r.redis.Set(key, dataJson, ttl)
	}

	return dataModel, err
}

func (r *NotificationRepository) FindOne(id int, userLoggedId int, isManager bool) (notification *domain.NotificationModel, err error) {

	var dataModel data_model.Notification
	ttl, _ := time.ParseDuration("300s")
	notification = &domain.NotificationModel{}

	key := fmt.Sprintf("notification.%d", id)
	data, err := r.redis.Get(key).Bytes()

	if len(data) == 0 {
		err = r.db.Where(&data_model.Notification{ID: uint(id)}).
			Take(&dataModel).Error

		if err != nil {
			return notification, err
		}

		if !isManager {
			return &domain.NotificationModel{}, gorm.ErrRecordNotFound
		}

		data, err = json.Marshal(&dataModel)

		if err != nil {
			return notification, gorm.ErrInvalidData
		}
		r.redis.Set(key, data, ttl)
	}

	err = json.Unmarshal(data, &dataModel)

	if err != nil {
		return notification, err
	}

	return (domain.NotificationModel{}).Load(&dataModel), err

}

func (r *NotificationRepository) ListNotifications(condictions *data_model.Notification, limit int) (notifications []*data_model.Notification) {
	dataModels := []*data_model.Notification{}

	r.db.Where(condictions).
		Find(&dataModels).
		Limit(limit)

	for _, data := range dataModels {
		model := (domain.NotificationModel{}).Load(data)
		notifications = append(notifications, model.GetDataModel())
	}

	return notifications
}
