package application

import (
	"sword-health/task/application/command"
	"sword-health/task/application/repositories"
	"sword-health/task/application/services"
	"sword-health/task/infra/amqp"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type Container struct {
	redis      *redis.Client
	db         *gorm.DB
	cmd        *command.NotificationHandler
	repository Repository
	service    Service
	msgBroker  *amqp.Connection
}

type Service struct {
	notificationWriter *services.WriteService
	notificationRead   *services.ReadService
}

type Repository struct {
	notification *repositories.NotificationRepository
}

func (Container) New(
	redis *redis.Client,
	db *gorm.DB,
	msgBroker *amqp.Connection,
) *Container {
	c := &Container{
		redis:     redis,
		db:        db,
		msgBroker: msgBroker,
	}
	return c.init()
}

func (c *Container) init() *Container {
	c.repository.notification = (repositories.NotificationRepository{}).
		New(c.redis, c.db)

	c.service.notificationWriter = (services.WriteService{}).
		New(c.repository.notification)

	c.service.notificationRead = (services.ReadService{}).
		New(c.repository.notification)

	c.cmd = (command.NotificationHandler{}).
		New(
			c.service.notificationWriter,
			c.service.notificationRead,
		)
	return c
}

func (c *Container) GetHandler() *command.NotificationHandler {
	return c.cmd
}
