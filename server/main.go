package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"grpctest/server/model"
	greeter "grpctest/server/proto"
)


func main(){
	reg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"127.0.0.1:8500",
		}
	})
	service :=micro.NewService(
		micro.Registry(reg),
		micro.Name("greeter"),
		)
	service.Init()
	greeter.RegisterSayHandler(service.Server(),new(model.Say))
	// run server
	if err := service.Run(); err != nil {
		panic(err)
	}
}
