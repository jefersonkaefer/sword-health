package providers

import (
	"sword-health/api/infra/amqp"
	"sword-health/api/validators"
)

type Container struct {
	Grpc       grpcClient
	AMQP       *amqp.Connection
	Validator  validators.JSONValidator
	Controller controller
}

func (c *Container) Run() *Container {
	c.initControllerProviders()
	return c
}

func (c *Container) Clear() {
	c.Grpc.User.CloseConnect()
}
