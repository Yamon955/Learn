package main

import (
	"fmt"
	"context"
	"trpc.group/trpc-go/trpc-go/client"
	"github.com/Yamon955/Learn/examples/service1/pb"
)

func main() {
	opts := []client.Option {
		client.WithTarget("ip://127.0.0.1:8080"),
	}
	c := pb.NewHelloWorldServiceClientProxy()
	rsp, err := c.Hello(context.Background(), &pb.HelloRequest{Msg: "world"}, opts...)
	if err != nil {
		fmt.Println(err.Error())
		return 
	}
	fmt.Println(rsp.Msg)
}
