package main

import (
	"github.com/869413421/micro-service/common/pkg/container"
	"github.com/869413421/micro-service/user-api/bootstarp"
	pb "github.com/869413421/micro-service/user/proto/user"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/web"
	"log"
	"time"
)

func main() {
	// 1.初始化web客户端
	g := bootstarp.SetupRoute()
	var serviceName = "micro.api.user"
	service := web.NewService(
		web.Name(serviceName),
		web.Address(":81"),
		web.Handler(g),
		// 指定服务注册信息在注册中心的有效期。 默认为一分种
		web.RegisterTTL(time.Minute*2),
		// 指定服务主动向注册中心报告健康状态的时间间隔,默认为 30 秒。
		web.RegisterInterval(time.Minute*1),
	)


	// 2.初始化用户服务客户端
	clientService := micro.NewService(
		micro.Name("pg.api.user.cli"),
	)
	client := pb.NewUserService("micro.service.user", clientService.Client())
	container.SetUserServiceClient(client)

	// 3.启动web客户端
	err := service.Init()
	if err != nil {
		log.Fatal("Init api error:", err)
	}
	err = service.Run()
	if err != nil {
		log.Fatal("start api error:", err)
		return
	}
}
