package main

import (
	"context"
	"fmt"

	"github.com/Yamon955/Learn/examples/service2/pb"
	"trpc.group/trpc-go/trpc-go/client"
)

func main() {
	opts := []client.Option{
		//client.WithTarget("ip://9.135.140.32:30001"),
		//client.WithTarget("consul://trpc.test.service2.HelloWorldService"),
		client.WithTarget("ip://127.0.0.1:8090"),
	}
	c := pb.NewHelloWorldServiceClientProxy()
	rsp, err := c.Hello(context.Background(), &pb.HelloRequest{Msg: "world"}, opts...)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(rsp.Msg)
}
