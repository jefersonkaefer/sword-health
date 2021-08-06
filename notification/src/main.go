package main

import (
	"log"
	"os"
	"sword-health/notification/application"
	"sword-health/notification/application/data_model"
	grpc_notification "sword-health/notification/infra/grpc"
	"sword-health/notification/infra/message"

	grpc_user "sword-health/notification/infra/grpc/client/user"

	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var container *application.Container

func main() {

	redisCli := redis.NewClient(&redis.Options{
		Addr:      os.Getenv("REDIS_ADDR"),
		Password: "",
		DB:       0,
	})

	broker := (message.AMQP{}).New(os.Getenv("BROKER_DSN"))

	grpc := grpc_notification.Server{}


	db, err := gorm.Open(
		mysql.Open(os.Getenv("DB_USERS_DSN")), 
		&gorm.Config{},
	)

	grpcClients := grpcConnection()

	container = (application.Container{}).New(redisCli, db, broker, grpcClients)

	err = db.AutoMigrate(&data_model.Notification{})

	if err != nil {
		log.Fatalln("Error: ", err)
	}

	go broker.
		QueueDeclare(
			message.NotificationExchange,
			message.NotificationQueue,
			message.NotificationRouteKeyCreate,
		).
		Consume(
			container.GetHandler(),
			"notification",
			"notification",
		)

	grpc.Start(container.GetHandler(), os.Getenv("GRPC_SERVER_PORT"))
}

func grpcConnection() *application.GrpcClient {

	grpc := application.GrpcClient{
		User: &grpc_user.UserClient{},
	}

	grpc.User.CreateConnection(
		os.Getenv("GRPC_USER_CLIENT_HOSTNAME"), 
		os.Getenv("GRPC_USER_CLIENT_PORT"),
	)
	grpc.User.Start()

	return &grpc
}
