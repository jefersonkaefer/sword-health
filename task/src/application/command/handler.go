package command

import (
	"sword-health/task/application/data_model"
	"sword-health/task/application/dto"
)

type Handler interface {
	Exec(cmd string, body []byte)
	Read() Read
	Write() Write
}

type Read interface {
	FindOne(userLoggedId int, taskId int) (task *data_model.Task, err error)
	ListTasks(userLoggedId int, ownerId int, limit int) (tasks []*data_model.Task, err error)
}

type Write interface {
	Create(taskRequest dto.TaskCreateDTO) (task *data_model.Task, err error)
	Update(taskRequest dto.TaskUpdateDTO) (task *data_model.Task, err error)
	Delete(userLoggedId int, id int) (err error)
	Close(userId int, id int) (err error)
}
