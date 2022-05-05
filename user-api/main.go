package main

import (
	"github.com/micro/go-micro/v2/web"
	"log"
	"time"
)

func main() {
	var serviceName = "micro.api.user"
	service := web.NewService(
		web.Name(serviceName),
		web.Address(":81"),
		// 指定服务注册信息在注册中心的有效期。 默认为一分种
		web.RegisterTTL(time.Minute*2),
		// 指定服务主动向注册中心报告健康状态的时间间隔,默认为 30 秒。
		web.RegisterInterval(time.Minute*1),
	)

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
