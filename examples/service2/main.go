package main

import (
	"context"
	"github.com/Yamon955/Learn/examples/service2/pb"
	trpc "trpc.group/trpc-go/trpc-go"
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
	log.Infof("got hello request: %s", req.Msg)
	return &pb.HelloResponse{Msg: "Hello " + req.Msg + "!"}, nil
}