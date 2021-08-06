package services

import (
	"encoding/json"
	"fmt"
	"sword-health/task/application/command"
	"sword-health/task/application/data_model"
	"sword-health/task/application/dto"
	"sword-health/task/application/repositories"
	"sword-health/task/domain"
	grpc_user "sword-health/task/infra/grpc/client/user"
	"sword-health/task/infra/message"

	"github.com/go-redis/redis"
)

type WriteService struct {
	redis          *redis.Client
	userClient     *grpc_user.UserClient
	taskRepository repositories.Repository
	broker         message.Broker
}

func (WriteService) New(
	repository repositories.Repository,
	redis *redis.Client,
	userClient *grpc_user.UserClient,
	broker message.Broker,
) command.Write {
	return &WriteService{
		taskRepository: repository,
		redis:          redis,
		userClient:     userClient,
		broker:         broker,
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

	user, err := us.userClient.Get(taskRequest.UserLoggedId)

	if err != nil {
		return task, err
	}

	err = model.Update(
		int(user.GetId()),
		user.GetIsManager(),
		taskRequest.Summary,
	)

	if err != nil {
		return task, err
	}

	task, err = us.taskRepository.Save(model)

	if err != nil {
		return task, err
	}

	go us.ClearCache(int(user.GetId()), int(model.GetId()))

	return task, err
}

func (us *WriteService) Close(userId int, id int) (err error) {

	model, err := us.taskRepository.FindOne(
		id,
	)

	if err != nil {
		return err
	}

	user, err := us.userClient.Get(userId)

	if err != nil {
		return err
	}

	if err != nil {
		return err
	}

	err = model.Close(
		int(user.GetId()),
		user.GetIsManager(),
	)
	fmt.Sprintf(
		"The task %d has been closed by %s.",
		model.GetId(),
		user.GetFullName(),
	)
	if err != nil {
		return err
	}

	task, err := us.taskRepository.Save(model)

	if err != nil {
		return err
	}

	go us.notify(
		int(model.GetId()),
		int(user.GetId()),
		fmt.Sprintf(
			"The task %d has been closed by %s.",
			model.GetId(),
			user.GetFullName(),
		),
	)

	go us.ClearCache(int(user.GetId()), int(task.ID))

	return err
}

func (us *WriteService) Delete(userId int, id int) (err error) {

	user, err := us.userClient.Get(userId)

	model, err := us.taskRepository.FindOne(
		id,
	)

	if err != nil {
		return err
	}

	err = model.Delete(
		user.GetIsManager(),
	)

	if err != nil {
		return err
	}

	err = us.taskRepository.Delete(model)

	if err != nil {
		return err
	}

	go us.ClearCache(int(user.GetId()), int(model.GetId()))

	return err
}

func (us *WriteService) ClearCache(userId int, taskId int) {

	keys := []string{
		fmt.Sprintf(CacheKeyTask, taskId),
		fmt.Sprintf(CacheKeyTaskListUser, userId),
	}

	us.redis.Del(keys...)
}

func (us *WriteService) notify(taskId int, fromId int, content string) {
	notification := dto.CreateNotificationDTO{
		NotificationType: "task",
		Content:          content,
		FromId:           fromId,
	}

	body, err := json.Marshal(notification)

	if err == nil {
		us.broker.Dispatch(
			message.NotificationExchange,
			message.NotificationRouteKeyCreate,
			body,
		)
	}
}
