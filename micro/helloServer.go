package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	hello "hello/proto/dexlace"
)

type HelloServer struct {
}

// 需要实现的方法

func (c *HelloServer) SayHello(ctx context.Context, req *hello.SayRequest,
	res *hello.SayResponse) error {

	res.Answer = "这是第一个micro server" + req.Message
	return nil
}

func main() {

	// 创建新的服务
	service := micro.NewService(
		micro.Name("dexlace.hello.server"),
	)

	// 初始化方法
	service.Init()

	// 注册服务
	hello.RegisterHelloHandler(service.Server(), new(HelloServer))

	// 运行服务
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}

}
