package main

import (
	"log"
	"time"

	micro "github.com/micro/go-micro"

	"cloudstore/service/account/handler"
	proto "cloudstore/service/account/proto"
)

func main() {
	// 创建一个service
	service := micro.NewService(
		micro.Name("go.micro.service.user"),
		micro.RegisterTTL(time.Second*10),
		micro.RegisterInterval(time.Second*5),
	)
	// 这里是为了接收命令行的参数
	service.Init()

	proto.RegisterUserServiceHandler(service.Server(), new(handler.User))
	if err := service.Run(); err != nil {
		log.Println(err)
	}
}
