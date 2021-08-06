package main

import (
	"log"
	"os"
	"sword-health/task/application"
	"sword-health/task/application/data_model"
	grpc_task "sword-health/task/infra/grpc"
	grpc_user "sword-health/task/infra/grpc/client/user"
	"sword-health/task/infra/message"

	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var container *application.Container

func main() {

	redisCli := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: "",
		DB:       0,
	})

	broker := (message.AMQP{}).New(os.Getenv("BROKER_DSN"))

	grpc := grpc_task.Server{}

	dsn := os.Getenv("DB_TASKS_DSN")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	grpcClients := grpcConnection()

	container = (application.Container{}).New(redisCli, db, broker, grpcClients)

	err = db.AutoMigrate(&data_model.Task{})

	if err != nil {
		log.Fatalln("Error: ", err)
	}

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
