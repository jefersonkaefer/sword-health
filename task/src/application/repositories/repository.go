package repositories

import (
	"sword-health/task/application/data_model"
	"sword-health/task/domain"
)

type Repository interface {
	Save(model *domain.TaskModel) (task *data_model.Task, err error)
	FindOne(id int) (task *domain.TaskModel, err error)
	ListTasks(ownerId int, limit int) (tasks []*domain.TaskModel, err error)
	Delete(model *domain.TaskModel) (err error)
}
