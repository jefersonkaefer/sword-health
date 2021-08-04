package main

import (
	"log"
	"sword-health/users/application"
	"sword-health/users/application/data_model"
	"sword-health/users/infra/amqp"
	grpc_user "sword-health/users/infra/grpc"

	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	redisCli := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	amqp := (amqp.Connection{}).New("guest", "guest", "rabbitmq", 5672)

	grpc := grpc_user.Server{}

	dsn := "root:swt4sks@tcp(mysql:3306)/sw_users?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	err = db.AutoMigrate(&data_model.User{})

	if err != nil {
		log.Fatalln("Error: ", err)
	}

	container = (application.Container{}).New(redisCli, db, amqp, grpc)

	container.Run(db)

	go amqp.Consume("user", container.GetHandler(), "user")

	grpc.Start(container.GetHandler().Read(), 5000)
}
