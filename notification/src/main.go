package main

import (
	"log"
	"sword-health/notification/application"
	"sword-health/notification/infra/amqp"
	grpc_notification "sword-health/notification/infra/grpc"
	"sword-health/task/application/data_model"

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

	grpc := grpc_notification.Server{}

	dsn := "root:swt4sks@tcp(mysql:3306)/sw_notifications?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	container = (application.Container{}).New(redisCli, db, amqp)

	err = db.AutoMigrate(&data_model.Notification{})

	if err != nil {
		log.Fatalln("Error: ", err)
	}

	go amqp.Consume("notification", container.GetHandler(), "notification")

	grpc.Start(container.GetHandler(), 5000)
}
