package main

import (
	"context"
	"fmt"
	//	_ "github.com/micro/go-plugins/registry/consul"

	micro "github.com/micro/go-micro/v2"
	proto "github.com/qiuzhanghua/go-with-micro/proto"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}

func main() {
	//	registry := consul.NewRegistry()
	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("greeter"),
		//		micro.Registry(registry),
	)

	// Init will parse the command line flags.
	service.Init()

	// Register handler
	proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
