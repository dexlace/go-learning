package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	hello "hello/proto/dexlace"
)

func main() {
	// 实例化
	service := micro.NewService(
		micro.Name("dexlace.hello.client"),
	)

	// 初始化
	service.Init()

	helloService := hello.NewHelloService("dexlace.hello.server", service.Client())

	res, err := helloService.SayHello(context.TODO(), &hello.SayRequest{Message: "美丽你好"})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res.Answer)

}
