package grpc_notification

import (
	context "context"
	"sword-health/notification/application/command"
)

type GrpcServer interface {
	Start(cmdService *command.NotificationHandler, port string)
	Get(ctx context.Context, in *NotificationRequest) (notification *Notification, err error)
	List(ctx context.Context, in *NotificationRequest) (List *ListNotification, err error)
}
