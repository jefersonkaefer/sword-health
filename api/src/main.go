package main

import (
	grpc_notification "sword-health/api/grpc/notification"
	grpc_user "sword-health/api/grpc/user"

	grpc_task "sword-health/api/grpc/task"
	middleware "sword-health/api/http"
	"sword-health/api/providers"

	"github.com/gin-gonic/gin"
)

var container providers.Container

func init() {

	container = providers.Container{}
	initGrpcConnection(&container)
}

func main() {

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
		PATCH("/task/:id/close", func(c *gin.Context) {
			container.Controller.Task.Close(c)
		})

	task := router.
		Group("/task").
		Use(middleware.VerifyToken())
	{
		task.
			GET("", func(c *gin.Context) {
				container.Controller.Task.List(c)
			}).
			GET("/:id", func(c *gin.Context) {
				container.Controller.Task.Get(c)
			}).
			Use(middleware.IsManager()).
			DELETE("/:id", func(c *gin.Context) {
				container.Controller.Task.Delete(c)
			})
	}
	notification := router.
		Group("/notification").
		Use(middleware.VerifyToken()).
		Use(middleware.IsManager())
	{
		notification.
			Use(middleware.VerifyToken()).
			Use(middleware.IsManager()).
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
