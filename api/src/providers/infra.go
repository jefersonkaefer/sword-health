package providers

import (
	grpc_notification "sword-health/api/grpc/notification"
	grpc_task "sword-health/api/grpc/task"
	grpc_user "sword-health/api/grpc/user"
	"sword-health/api/infra/amqp"
)

type grpcClient struct {
	User         *grpc_user.UserClient
	Task         *grpc_task.TaskClient
	Notification *grpc_notification.NotificationClient
}
type amqpClient struct {
	User amqp.Connection
}
