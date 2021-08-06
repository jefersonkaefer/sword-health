package providers

import (
	"sword-health/api/validators"
)

type Container struct {
	Grpc       grpcClient
	Validator  validators.JSONValidator
	Controller controller
}

func (c *Container) Run() *Container {
	c.initControllerProviders()
	return c
}