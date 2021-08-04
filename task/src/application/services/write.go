package services

import (
	"errors"
	"fmt"
	"sword-health/task/application/data_model"
	"sword-health/task/application/dto"
	"sword-health/task/application/repositories"
	"sword-health/task/domain"
)

type WriteService struct {
	taskRepository *repositories.TaskRepository
}

func (WriteService) New(repository *repositories.TaskRepository) *WriteService {
	return &WriteService{
		taskRepository: repository,
	}
}

func (us *WriteService) Create(taskRequest dto.TaskCreateDTO) (task *data_model.Task, err error) {

	newTask, err := domain.Create(
		taskRequest.Summary,
		taskRequest.OwnerId,
	)

	task, err = us.taskRepository.Save(&newTask)

	return task, err
}

func (us *WriteService) Update(taskRequest dto.TaskUpdateDTO) (task *data_model.Task, err error) {

	cond := data_model.Task{
		ID: uint(taskRequest.Id),
	}

	if taskRequest.UserLoggedRole != "manager" {
		cond.OwnerId = taskRequest.UserLoggedId
	}

	model, err := us.taskRepository.FindOne(&cond)

	if err != nil {
		return task, err
	}

	if taskRequest.UserLoggedRole != "manager" && uint(taskRequest.UserLoggedId) != model.GetOwnerId() {
		cond.OwnerId = taskRequest.UserLoggedId
	}

	if cond.OwnerId != int(model.GetOwnerId()) {
		return task, errors.New("Task not found.")
	}

	err = model.Update(
		taskRequest.Summary,
		taskRequest.Status,
		taskRequest.UserLoggedId,
	)

	if err != nil {
		fmt.Println("ERROR: ", err)
		return task, err
	}

	task, err = us.taskRepository.Save(model)

	return task, err
}
