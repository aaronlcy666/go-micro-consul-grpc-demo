package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul" // 注意地址 已经修改到 go-plugins
	"grpctest/server/model"
	greeter "grpctest/server/proto"
)

func main() {
	//consul 注册服务
	reg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"127.0.0.1:8500",
		}
	})
	service := micro.NewService(
		micro.Registry(reg),   // 把consul注册到micro
		micro.Name("greeter"), // name很重要 服务发现就靠它
	)
	service.Init() // init 固定步骤

	//调用注册方法 类似于http里面的注册路由
	greeter.RegisterSayHandler(service.Server(), new(model.Say)) //注册服务端
	// run server
	if err := service.Run(); err != nil { //启动
		panic(err)
	}
}
