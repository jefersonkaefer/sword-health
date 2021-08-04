package services

import (
	"sword-health/task/application/data_model"
	"sword-health/task/application/dto"
	"sword-health/task/application/repositories"
)

type ReadService struct {
	taskRepository *repositories.TaskRepository
}

func (ReadService) New(repository *repositories.TaskRepository) *ReadService {
	return &ReadService{
		taskRepository: repository,
	}
}

func (us *ReadService) ListTasks(params *dto.FindTask) []*data_model.Task {

	taskDataModel := &data_model.Task{
		OwnerId: params.OwnerId,
	}

	if params.Role != "manager" {
		taskDataModel.OwnerId = params.UserLoggedId
	}

	return us.taskRepository.ListTasks(taskDataModel, params.Limit)
}

func (us *ReadService) FindOne(params *dto.FindTask) (task *data_model.Task, err error) {

	taskDataModel := &data_model.Task{
		ID: uint(params.Id),
	}

	if params.Role != "manager" || params.OwnerId == 0 {
		taskDataModel.OwnerId = params.UserLoggedId
	}

	model, err := us.taskRepository.FindOne(taskDataModel)

	return model.GetDataModel(), err
}
