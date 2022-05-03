package main

import (
	"github.com/869413421/micro-service/common/pkg/db"
	"github.com/869413421/micro-service/user/handler"
	"github.com/869413421/micro-service/user/pkg/model"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"

	proto "github.com/869413421/micro-service/user/proto/user"
)

func main() {

	// 1.准备数据库连接，并且执行数据库迁移
	db := db.GetDB()
	db.AutoMigrate(&model.User{})

	// 2.创建服务
	service := micro.NewService(
		micro.Name("micro.service.user"),
		micro.Version("v1"),
	)

	// 3.初始化服务
	service.Init()

	// 4.注册服务处理器
	proto.RegisterUserServiceHandler(service.Server(),handler.NewUserServiceHandler())

	// 5.运行服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
