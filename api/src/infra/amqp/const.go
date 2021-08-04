package amqp

const (
	ExchangeUser   = "ex-user"
	ExchangeTask   = "ex-task"
	ExchangeNotify = "ex-notify"

	QueueUser   = "user"
	QueueTask   = "task"
	QueueNotify = "notify"

	RouteKeyUserCreate = "user.create"
	RouteKeyUserUpdate = "user.update"
	RouteKeyUserDelete = "user.delete"

	RouteKeyTaskCreate = "task.create"
	RouteKeyTaskUpdate = "task.update"
	RouteKeyTaskDelete = "task.delete"
)
