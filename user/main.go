package main

import (
	"github.com/869413421/micro-service/common/pkg/db"
	"github.com/869413421/micro-service/user/handler"
	"github.com/869413421/micro-service/user/pkg/model"
	"github.com/869413421/micro-service/user/subscriber"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"

	proto "github.com/869413421/micro-service/user/proto/user"
)

func main() {

	//1.准备数据库连接，并且执行数据库迁移
	db := db.GetDB()
	db.AutoMigrate(&model.User{})

	// New Service
	service := micro.NewService(
		micro.Name("micro.service.user"),
		micro.Version("v1"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	proto.RegisterUserHandler(service.Server(), new(handler.User))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("micro.service.user", service.Server(), new(subscriber.User))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
