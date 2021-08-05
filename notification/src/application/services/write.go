package services

import (
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

	model, err := us.taskRepository.FindOne(
		taskRequest.Id,
		taskRequest.UserLoggedId,
		taskRequest.UserLoggedRole == "manager",
	)

	if err != nil {
		return task, err
	}

	err = model.Update(
		taskRequest.Summary,
		taskRequest.Status,
		taskRequest.UserLoggedId,
		taskRequest.UserLoggedRole == "manager",
	)

	if err != nil {
		return task, err
	}

	task, err = us.taskRepository.Save(model)

	return task, err
}

func (us *WriteService) Delete(taskRequest dto.TaskUpdateDTO) (err error) {

	isManager := taskRequest.UserLoggedRole == "manager"

	model, err := us.taskRepository.FindOne(
		taskRequest.Id,
		taskRequest.UserLoggedId,
		isManager,
	)

	if err != nil {
		return err
	}

	err = model.Delete(
		taskRequest.UserLoggedId,
		isManager,
	)

	if err != nil {
		return err
	}

	return us.taskRepository.Delete(model)
}
