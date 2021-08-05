package main

import (
	grpc_notification "sword-health/api/grpc/notification"
	grpc_user "sword-health/api/grpc/user"

	grpc_task "sword-health/api/grpc/task"
	middleware "sword-health/api/http"
	"sword-health/api/infra/amqp"
	"sword-health/api/providers"

	"github.com/gin-gonic/gin"
)

var container providers.Container

func init() {

	container = providers.Container{}
	initAmqpConnection(&container)
	initGrpcConnection(&container)
}

func main() {
	defer container.Clear()

	router := gin.Default()
	router.
		Use(middleware.JSONMiddleware()).
		POST("/auth", func(c *gin.Context) {
			container.Controller.Auth.Login(c)
		}).
		POST("/user", func(c *gin.Context) {
			container.Controller.User.Create(c)
		}).
		Use(middleware.VerifyToken()).
		POST("/task", func(c *gin.Context) {
			container.Controller.Task.Create(c)
		}).
		PATCH("/task/:id", func(c *gin.Context) {
			container.Controller.Task.Update(c)
		})

	tasks := router.
		Group("/task").
		Use(middleware.VerifyToken()).
		Use(middleware.IsManager())
	{
		tasks.
			GET("", func(c *gin.Context) {
				container.Controller.Task.List(c)
			}).
			GET("/:id", func(c *gin.Context) {
				container.Controller.Task.Get(c)
			})
	}
	notification := router.
		Group("/notification").
		Use(middleware.VerifyToken()).
		Use(middleware.IsManager())
	{
		notification.
			GET("", func(c *gin.Context) {
				container.Controller.Notification.List(c)
			}).
			GET("/:id", func(c *gin.Context) {
				container.Controller.Notification.Get(c)
			})
	}

	container.Run()
	router.Run(":8000")

}

func initGrpcConnection(c *providers.Container) {

	c.Grpc.User = &grpc_user.UserClient{}
	c.Grpc.User.CreateConnection("user", 5000)
	c.Grpc.User.Start()

	c.Grpc.Task = &grpc_task.TaskClient{}
	c.Grpc.Task.CreateConnection("task", 5000)
	c.Grpc.Task.Start()

	c.Grpc.Notification = &grpc_notification.NotificationClient{}
	c.Grpc.Notification.CreateConnection("notification", 5000)
	c.Grpc.Notification.Start()
}

func initAmqpConnection(c *providers.Container) {
	c.AMQP = amqp.NewConnection("guest", "guest", "rabbitmq", 5672).
		DeclareExchange(amqp.ExchangeUser).
		DeclareExchange(amqp.ExchangeTask).
		QueueDeclare(
			amqp.ExchangeUser,
			amqp.QueueUser,
			amqp.RouteKeyUserCreate,
		).
		QueueDeclare(
			amqp.ExchangeTask,
			amqp.QueueTask,
			amqp.RouteKeyTaskCreate,
		)
}
