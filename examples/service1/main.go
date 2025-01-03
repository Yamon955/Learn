package main

import (
	"context"

	trpc "trpc.group/trpc-go/trpc-go"
	"github.com/Yamon955/Learn/examples/service1/pb"
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
	log.Infof("got request: %s", req.Msg)
	return &pb.HelloResponse{Msg: "I'm service1 !"}, nil
}