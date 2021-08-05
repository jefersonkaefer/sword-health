package application

import (
	"sword-health/task/application/command"
	"sword-health/task/application/repositories"
	"sword-health/task/application/services"
	"sword-health/task/infra/amqp"
	grpc_user "sword-health/task/infra/grpc/client/user"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type Container struct {
	redis      *redis.Client
	db         *gorm.DB
	cmd        *command.TaskHandler
	repository Repository
	service    Service
	msgBroker  *amqp.Connection
	grpc       *GrpcClient
}

type GrpcClient struct {
	User *grpc_user.UserClient
}

type Service struct {
	tasksWriter *services.WriteService
	tasksRead   *services.ReadService
}

type Repository struct {
	tasks *repositories.TaskRepository
}

func (Container) New(
	redis *redis.Client,
	db *gorm.DB,
	msgBroker *amqp.Connection,
	grpcClient *GrpcClient,
) *Container {
	c := &Container{
		redis:     redis,
		db:        db,
		msgBroker: msgBroker,
		grpc:      grpcClient,
	}
	return c.init()
}

func (c *Container) init() *Container {
	c.repository.tasks = (repositories.TaskRepository{}).
		New(c.redis, c.db, c.grpc.User)

	c.service.tasksWriter = (services.WriteService{}).
		New(
			c.repository.tasks,
			c.redis,
			c.grpc.User,
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

func (c *Container) GetHandler() *command.TaskHandler {
	return c.cmd
}
