package repositories

import (
	"fmt"
	"sword-health/task/application/data_model"
	"sword-health/task/domain"
	grpc_user "sword-health/task/infra/grpc/client/user"

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
) Repository {
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

	err = r.db.Where(
		data_model.Task{
			ID: uint(id),
		},
	).Take(&dataModel).Error

	if err != nil {
		return task, err
	}
	fmt.Println(" DATA ", id)
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

	err = r.db.Delete(&dataModel).Error

	if err != nil {
		fmt.Println("ERROR ", err)
	}

	return err
}
