package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	greeter "grpctest/server/proto"
)

func main() {
	reg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"127.0.0.1:8500",
		}
	})
	// 初始化服务 同一个consul
	service := micro.NewService(
		micro.Registry(reg),
	)
	service.Init()
	//服务发现   凭借我们写的name 使用NewSayService

	//这边会返回一个name为greeter并且 client.Client 为 service.Client() 的 接口 接口依然是 proto帮我们实现的
	//	当前代码下 接口下有一个方法 就是 Hello方法 并且有一个name为 greeter  如果不传name 那么它默认会是我们声明的服务接口名Say
	sayClient := greeter.NewSayService("greeter", service.Client())
	rsp, err := sayClient.Hello(context.Background(), &greeter.SayParam{Msg: "hello server"})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp)
}
