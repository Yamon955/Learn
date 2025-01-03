package main

import (
	"context"
	"fmt"

	cupb "github.com/Yamon955/Learn/examples/caculator/pb"
	"github.com/Yamon955/Learn/examples/service2/pb"
	"google.golang.org/protobuf/proto"
	trpc "trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/client"
	"trpc.group/trpc-go/trpc-go/log"
)

func main() {
	s := trpc.NewServer()
	pb.RegisterHelloWorldServiceService(s, &Greeter{})
	if err := s.Serve(); err != nil {
		log.Error(err)
	}
}

type Greeter struct{}

func (g Greeter) Hello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	opts := []client.Option {
		client.WithServiceName("trpc.test.caculator.Caculator"),
		//client.WithTarget("ip://127.0.0.1:9000"),
	}
	proxy := cupb.NewCaculatorClientProxy()
	a := 1.0
	b := 1.0
	req2 := &cupb.CaculateReq{
		A: proto.Float64(a),
		B: proto.Float64(b),
		Op: cupb.Operators_ADD.Enum(),
	}
	cuRsp, err := proxy.Caculate(ctx, req2, opts...)
	if err != nil {
		log.ErrorContextf(ctx, "Caculate fail, err:%v", err)
		return &pb.HelloResponse{Msg: "Caculate fail !"}, nil
	}
	msg := fmt.Sprintf("%g + %g = %g", a, b, cuRsp.GetAns())
	return &pb.HelloResponse{Msg: msg}, nil
}