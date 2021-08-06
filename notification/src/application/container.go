package application

import (
	"sword-health/notification/application/command"
	"sword-health/notification/application/repositories"
	"sword-health/notification/application/services"
	"sword-health/notification/infra/message"

	grpc_user "sword-health/notification/infra/grpc/client/user"

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

type Service struct {
	notificationWriter command.Write
	notificationRead   command.Read
}

type Repository struct {
	notification repositories.Repository
}

type GrpcClient struct {
	User *grpc_user.UserClient
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
	c.repository.notification = (repositories.NotificationRepository{}).
		New(c.db)

	c.service.notificationWriter = (services.WriteService{}).
		New(c.redis, c.grpc.User, c.repository.notification)

	c.service.notificationRead = (services.ReadService{}).
		New(c.redis, c.grpc.User, c.repository.notification)

	c.cmd = (command.NotificationHandler{}).
		New(
			c.service.notificationWriter,
			c.service.notificationRead,
		)
	return c
}

func (c *Container) GetHandler() command.Handler {
	return c.cmd
}
