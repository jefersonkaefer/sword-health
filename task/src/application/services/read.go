package services

import (
	"encoding/json"
	"fmt"
	"sword-health/task/application/data_model"
	"sword-health/task/application/repositories"
	"sword-health/task/domain"
	grpc_user "sword-health/task/infra/grpc/client/user"
	"time"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type ReadService struct {
	redis          *redis.Client
	userClient     *grpc_user.UserClient
	taskRepository *repositories.TaskRepository
}

func (ReadService) New(
	repository *repositories.TaskRepository,
	redis *redis.Client,
	userClient *grpc_user.UserClient,
) *ReadService {
	return &ReadService{
		taskRepository: repository,
		redis:          redis,
		userClient:     userClient,
	}
}

func (us *ReadService) FindOne(userLoggedId int, taskId int) (task *data_model.Task, err error) {

	userLogged, err := us.userClient.Get(userLoggedId)

	user := (domain.User{}).Load(int(userLogged.GetId()), userLogged.GetRole())

	key := fmt.Sprintf("task.%d", taskId)

	data, err := us.redis.Get(key).Bytes()

	if len(data) == 0 {
		model, err := us.taskRepository.FindOne(
			taskId,
		)

		if err != nil {
			return task, err
		}

		user, err := us.userClient.Get(int(model.GetOwnerId()))

		task := model.GetDataModel()

		if err != nil {
			return task, err
		}

		task.OwnerFirstName = user.GetFirstName()
		task.OwnerLastName = user.GetLastName()
		task.OwnerEmail = user.GetEmail()

		data, err = json.Marshal(&task)

		if err != nil {
			return task, gorm.ErrInvalidData
		}

		ttl, _ := time.ParseDuration("300s")

		us.redis.Set(key, data, ttl)
	}

	err = json.Unmarshal(data, &task)

	model := (domain.TaskModel{}).Load(task)

	if !model.IsOwner(user.GetId()) && !user.IsManager() {
		return task, gorm.ErrRecordNotFound
	}

	return model.GetDataModel(), err
}

func (us *ReadService) ListTasks(userLoggedId int, ownerId int, limit int) (tasks []*data_model.Task, err error) {
	userLogged, err := us.userClient.Get(userLoggedId)

	user := (domain.User{}).Load(int(userLogged.GetId()), userLogged.GetRole())

	if err != nil {
		return tasks, err
	}

	if !user.IsManager() && ownerId != int(userLogged.GetId()) {
		ownerId = int(userLogged.GetId())
	}

	key := fmt.Sprintf("task.list.user.%d", userLoggedId)

	data, err := us.redis.Get(key).Bytes()

	if len(data) == 0 {
		var taskModels []*domain.TaskModel

		taskModels, err = us.taskRepository.ListTasks(ownerId, limit)

		users := make(map[int]*grpc_user.User)

		for _, task := range taskModels {

			user, ok := users[int(task.GetOwnerId())]

			if !ok {
				user, err = us.userClient.Get(int(task.GetOwnerId()))
				users[int(user.GetId())] = user
				if err != nil {
					return tasks, err
				}
			}

			taskDataModel := task.GetDataModel()

			taskDataModel.OwnerFirstName = user.GetFirstName()
			taskDataModel.OwnerLastName = user.GetLastName()
			taskDataModel.OwnerEmail = user.GetEmail()

			tasks = append(tasks, taskDataModel)

		}

		data, err = json.Marshal(&tasks)

		if err != nil {
			return tasks, gorm.ErrInvalidData
		}

		ttl, _ := time.ParseDuration("300s")

		us.redis.Set(key, data, ttl)

	}

	err = json.Unmarshal(data, &tasks)

	return tasks, err
}
