package repositories

import (
	"encoding/json"
	"fmt"
	"sword-health/task/application/data_model"
	"sword-health/task/domain"
	grpc_user "sword-health/task/infra/grpc/client/user"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type TaskRepository struct {
	redis      *redis.Client
	db         *gorm.DB
	userClient *grpc_user.UserClient
}

func (TaskRepository) New(
	redis *redis.Client,
	db *gorm.DB,
	userClient *grpc_user.UserClient,
) *TaskRepository {
	return &TaskRepository{
		redis:      redis,
		db:         db,
		userClient: userClient,
	}
}

func (r *TaskRepository) Save(model *domain.TaskModel) (task *data_model.Task, err error) {
	dataModel := model.GetDataModel()

	r.db.Save(dataModel)

	user, err := r.userClient.Get(dataModel.OwnerId)

	if err == nil {
		dataModel.OwnerFirstName = user.GetFirstName()
		dataModel.OwnerLastName = user.GetLastName()
		dataModel.OwnerEmail = user.GetEmail()
	}

	key := fmt.Sprintf("task.%d.user.%d", dataModel.ID, dataModel.OwnerId)

	dataJson, err := json.Marshal(dataModel)

	if err == nil {
		r.redis.Set(key, dataJson, 0)
	}

	return dataModel, err
}

func (r *TaskRepository) FindOne(condictions *data_model.Task) (task *domain.TaskModel, err error) {

	var dataModel data_model.Task

	key := fmt.Sprintf("task.%d.user.%d", condictions.ID, condictions.OwnerId)

	data, err := r.redis.Get(key).Bytes()

	if len(data) == 0 {
		err = r.db.Where(condictions).
			Take(&dataModel).Error

		user, err := r.userClient.Get(dataModel.OwnerId)

		if err != nil {
			return task, gorm.ErrInvalidData
		}
		if err == nil {
			dataModel.OwnerFirstName = user.GetFirstName()
			dataModel.OwnerLastName = user.GetLastName()
			dataModel.OwnerEmail = user.GetEmail()
		}
		if err != nil {
			return task, gorm.ErrRecordNotFound
		}

		data, err = json.Marshal(&dataModel)

		if err == nil {
			r.redis.Set(key, data, 0)
		}
	}

	err = json.Unmarshal(data, &dataModel)

	if err != nil {
		return task, err
	}

	return (domain.TaskModel{}).Load(&dataModel), err

}

func (r *TaskRepository) ListTasks(condictions *data_model.Task, limit int) (tasks []*data_model.Task) {

	r.db.Where(condictions).
		Find(&tasks).
		Limit(limit)

	for _, data := range tasks {
		model := (domain.TaskModel{}).Load(data)
		tasks = append(tasks, model.GetDataModel())
	}

	return tasks
}
