package repositories

import (
	"encoding/json"
	"fmt"
	"sword-health/task/application/data_model"
	"sword-health/task/domain"
	grpc_user "sword-health/task/infra/grpc/client/user"
	"time"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type TaskRepository struct {
	redis *redis.Client
	db    *gorm.DB
}

func (TaskRepository) New(
	redis *redis.Client,
	db *gorm.DB,
	userClient *grpc_user.UserClient,
) *TaskRepository {
	return &TaskRepository{
		redis: redis,
		db:    db,
	}
}

func (r *TaskRepository) Save(model *domain.TaskModel) (task *data_model.Task, err error) {
	dataModel := model.GetDataModel()

	r.db.Save(dataModel)

	return dataModel, err
}

func (r *TaskRepository) FindOne(id int) (task *domain.TaskModel, err error) {

	var dataModel data_model.Task
	ttl, _ := time.ParseDuration("300s")
	task = &domain.TaskModel{}

	key := fmt.Sprintf("task.%d", id)
	data, err := r.redis.Get(key).Bytes()

	if len(data) == 0 {
		err = r.db.Where(&data_model.Task{ID: uint(id)}).
			Take(&dataModel).Error

		if err != nil {
			return task, err
		}

		if err != nil {
			return task, gorm.ErrInvalidData
		}

		data, err = json.Marshal(&dataModel)

		if err != nil {
			return task, gorm.ErrInvalidData
		}
		r.redis.Set(key, data, ttl)
	}

	err = json.Unmarshal(data, &dataModel)

	if err != nil {
		return task, err
	}

	return (domain.TaskModel{}).Load(&dataModel), err

}

func (r *TaskRepository) ListTasks(ownerId int, limit int) (tasks []*domain.TaskModel, err error) {

	dataModels := []*data_model.Task{}

	if err != nil {
		return tasks, err
	}

	condictions := data_model.Task{
		OwnerId: ownerId,
	}

	r.db.Where(condictions).
		Find(&dataModels).
		Limit(limit)

	for _, data := range dataModels {
		model := (domain.TaskModel{}).Load(data)
		tasks = append(tasks, model)
	}

	return tasks, err
}

func (r *TaskRepository) Delete(model *domain.TaskModel) (err error) {

	dataModel := model.GetDataModel()

	return r.db.Delete(dataModel).Error
}
