package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	greeter "grpctest/server/proto"
)

func main(){
	reg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"127.0.0.1:8500",
		}
	})
	// 初始化服务
		service := micro.NewService(
		micro.Registry(reg),
	)
	service.Init()
		//服务发现
	sayClent := greeter.NewSayService("greeter",service.Client())
	rsp, err := sayClent.Hello(context.Background(), &greeter.SayParam{Msg: "hello server"})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp)
}
