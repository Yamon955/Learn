package main

import (
	"fmt"
	"context"
	"trpc.group/trpc-go/trpc-go/client"
	"github.com/Yamon955/Learn/examples/service2/pb"
)

func main() {
	opts := []client.Option {
		client.WithTarget("ip://9.135.140.32:30001"),
	}
	c := pb.NewHelloWorldServiceClientProxy()
	rsp, err := c.Hello(context.Background(), &pb.HelloRequest{Msg: "world"}, opts...)
	if err != nil {
		fmt.Println(err.Error())
		return 
	}
	fmt.Println(rsp.Msg)
}