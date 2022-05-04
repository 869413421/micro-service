package main

import (
	"github.com/869413421/micro-service/common/pkg/container"
	"github.com/869413421/micro-service/common/pkg/db"
	"github.com/869413421/micro-service/user/handler"
	"github.com/869413421/micro-service/user/pkg/model"
	"github.com/869413421/micro-service/user/subscriber"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"

	proto "github.com/869413421/micro-service/user/proto/user"
)

func main() {

	// 1.准备数据库连接，并且执行数据库迁移
	db := db.GetDB()
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.PasswordReset{})

	// 2.创建服务
	service := micro.NewService(
		micro.Name("micro.service.user"),
		micro.Version("v1"),
	)

	// 3.初始化服务,全局化service对象
	service.Init()
	container.SetService(service)

	// 4.初始化borker
	brk := service.Options().Broker
	defer brk.Disconnect()
	err := brk.Init()
	if err != nil {
		log.Fatal(err)
		return
	}
	err = brk.Connect()
	if err != nil {
		log.Fatal("connection broker error:", err)
		return
	}

	// 5.订阅事件
	eventSubscriber := subscriber.NewEventSubscriber(brk)
	err = eventSubscriber.Subscriber()
	if err != nil {
		log.Fatal("subscriber broker error:", err)
		return
	}

	// 6.注册服务处理器
	proto.RegisterUserServiceHandler(service.Server(), handler.NewUserServiceHandler())

	// 7.运行服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
