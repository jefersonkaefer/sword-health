package services

import (
	"sword-health/task/application/data_model"
	"sword-health/task/application/dto"
	"sword-health/task/application/repositories"
	"sword-health/task/domain"
	grpc_user "sword-health/task/infra/grpc/client/user"

	"github.com/go-redis/redis"
)

type WriteService struct {
	redis          *redis.Client
	userClient     *grpc_user.UserClient
	taskRepository *repositories.TaskRepository
}

func (WriteService) New(
	repository *repositories.TaskRepository,
	redis *redis.Client,
	userClient *grpc_user.UserClient,
) *WriteService {
	return &WriteService{
		taskRepository: repository,
		redis:          redis,
		userClient:     userClient,
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
	)

	userLogged, err := us.userClient.Get(taskRequest.UserLoggedId)

	user := (domain.User{}).Load(int(userLogged.GetId()), userLogged.GetRole())

	if err != nil {
		return task, err
	}

	err = model.Update(
		user.GetId(),
		user.IsManager(),
		taskRequest.Summary,
		taskRequest.Status,
	)

	if err != nil {
		return task, err
	}

	task, err = us.taskRepository.Save(model)

	return task, err
}

func (us *WriteService) Delete(taskRequest dto.TaskUpdateDTO) (err error) {

	userLogged, err := us.userClient.Get(taskRequest.UserLoggedId)

	user := (domain.User{}).Load(int(userLogged.GetId()), userLogged.GetRole())

	model, err := us.taskRepository.FindOne(
		taskRequest.Id,
	)

	if err != nil {
		return err
	}

	err = model.Delete(
		taskRequest.UserLoggedId,
		user.IsManager(),
	)

	if err != nil {
		return err
	}

	return us.taskRepository.Delete(model)
}

func IsManager(UserRole string) bool {
	return UserRole == "manager"
}
