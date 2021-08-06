package main

import (
	"log"
	"os"
	"sword-health/user/application"
	"sword-health/user/application/data_model"
	grpc_user "sword-health/user/infra/grpc"

	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	redisCli := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: "",
		DB:       0,
	})

	grpc := grpc_user.Server{}

	db, err := gorm.Open(
		mysql.Open(os.Getenv("DB_USERS_DSN")), 
		&gorm.Config{},
	)

	err = db.AutoMigrate(&data_model.User{})

	if err != nil {
		log.Fatalln("Error: ", err)
	}

	container := (application.Container{}).New(redisCli, db)

	grpc.Start(container.GetHandler(), os.Getenv("GRPC_SERVER_PORT"))
}
