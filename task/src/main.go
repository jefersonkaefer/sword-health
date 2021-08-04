package main

import (
	"log"
	"sword-health/task/application"
	"sword-health/task/application/data_model"
	"sword-health/task/infra/amqp"
	grpc_task "sword-health/task/infra/grpc"
	grpc_user "sword-health/task/infra/grpc/client/user"

	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var container *application.Container

func main() {

	redisCli := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	amqp := (amqp.Connection{}).New("guest", "guest", "rabbitmq", 5672)

	grpc := grpc_task.Server{}

	dsn := "root:swt4sks@tcp(mysql:3306)/sw_tasks?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	grpcClients := grpcConnection()

	container = (application.Container{}).New(redisCli, db, amqp, grpcClients)

	err = db.AutoMigrate(&data_model.Task{})

	if err != nil {
		log.Fatalln("Error: ", err)
	}

	go amqp.Consume("task", container.GetHandler(), "task")

	grpc.Start(container.GetHandler(), 5000)
}

func grpcConnection() *application.GrpcClient {

	grpc := application.GrpcClient{
		User: &grpc_user.UserClient{},
	}

	grpc.User.CreateConnection("users", 5000)
	grpc.User.Start()

	return &grpc
}
