package application

import (
	"sword-health/task/application/command"
	"sword-health/task/application/repositories"
	"sword-health/task/application/services"
	grpc_user "sword-health/task/infra/grpc/client/user"
	"sword-health/task/infra/message"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type Container struct {
	redis      *redis.Client
	db         *gorm.DB
	cmd        command.Handler
	repository Repository
	service    Service
	broker     message.Broker
	grpc       *GrpcClient
}

type GrpcClient struct {
	User *grpc_user.UserClient
}

type Service struct {
	tasksWriter command.Write
	tasksRead   command.Read
}

type Repository struct {
	tasks repositories.Repository
}

func (Container) New(
	redis *redis.Client,
	db *gorm.DB,
	broker message.Broker,
	grpcClient *GrpcClient,
) *Container {
	c := &Container{
		redis:  redis,
		db:     db,
		broker: broker,
		grpc:   grpcClient,
	}
	return c.init()
}

func (c *Container) init() *Container {
	c.repository.tasks = (repositories.TaskRepository{}).
		New(
			c.redis,
			c.db,
			c.grpc.User,
		)

	c.service.tasksWriter = (services.WriteService{}).
		New(
			c.repository.tasks,
			c.redis,
			c.grpc.User,
			c.broker,
		)

	c.service.tasksRead = (services.ReadService{}).
		New(
			c.repository.tasks,
			c.redis,
			c.grpc.User,
		)

	c.cmd = (command.TaskHandler{}).
		New(
			c.service.tasksWriter,
			c.service.tasksRead,
		)
	return c
}

func (c *Container) GetHandler() command.Handler {
	return c.cmd
}
