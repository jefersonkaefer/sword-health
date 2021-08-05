package providers

import (
	"sword-health/api/controllers"
	"sword-health/api/validators"
)

type controller struct {
	Auth         controllers.AuthController
	User         controllers.UserController
	Task         controllers.TaskController
	Notification controllers.NotificationController
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
		},
		Task: controllers.TaskController{
			Validator:  &validator,
			TaskClient: c.Grpc.Task,
		},
		Notification: controllers.NotificationController{
			Validator:          &validator,
			NotificationClient: c.Grpc.Notification,
		},
	}
}
