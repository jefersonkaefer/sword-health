package providers

import (
	grpc_task "sword-health/api/grpc/task"
	grpc_user "sword-health/api/grpc/user"
	"sword-health/api/infra/amqp"
)

type grpcClient struct {
	User *grpc_user.UserClient
	Task *grpc_task.TaskClient
}
type amqpClient struct {
	User amqp.Connection
}
