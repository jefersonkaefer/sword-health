package application

import (
	"sword-health/users/application/command"
	"sword-health/users/application/repositories"
	"sword-health/users/application/services"
	"sword-health/users/infra/amqp"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type Container struct {
	redis      *redis.Client
	db         *gorm.DB
	cmd        *command.UserHandler
	repository Repository
	service    Service
	msgBroker  *amqp.Connection
}

type Service struct {
	usersWriter *services.WriteService
	usersRead   *services.ReadService
}

type Repository struct {
	users *repositories.UserRepository
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
	c.repository.users = (repositories.UserRepository{}).
		New(c.redis, c.db)

	c.service.usersWriter = (services.WriteService{}).
		New(c.repository.users)

	c.service.usersRead = (services.ReadService{}).
		New(c.repository.users)

	c.cmd = (command.UserHandler{}).
		New(
			c.service.usersWriter,
			c.service.usersRead,
		)
	return c
}

func (c *Container) GetHandler() *command.UserHandler {
	return c.cmd
}
