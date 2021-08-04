package providers

import (
	"sword-health/api/controllers"
	"sword-health/api/validators"
)

type controller struct {
	Auth controllers.AuthController
	User controllers.UserController
	Task controllers.TaskController
}

func (c *Container) initControllerProviders() {

	validator := validators.JSONValidator{}

	c.Controller = controller{
		Auth: controllers.AuthController{
			Validator:  &validator,
			UserClient: c.Grpc.User,
		},
		User: controllers.UserController{
			Validator:  &validator,
			UserClient: c.Grpc.User,
			AMQ:        c.AMQP,
		},
		Task: controllers.TaskController{
			Validator:  &validator,
			TaskClient: c.Grpc.Task,
			AMQ:        c.AMQP,
		},
	}
}
