package main

import (
	"context"
	"fmt"

	pb "github.com/Yamon955/Learn/examples/caculator/pb"
	"google.golang.org/protobuf/proto"
	"trpc.group/trpc-go/trpc-go/client"
	"trpc.group/trpc-go/trpc-go/log"
)


func main() {
	opts := []client.Option{
		client.WithTarget("ip://192.168.58.2:30001"),
	}
	proxy := pb.NewCaculatorClientProxy()
	a := 1.0
	b := 1.0
	req2 := &pb.CaculateReq{
		A: proto.Float64(a),
		B: proto.Float64(b),
		Op: pb.Operators_ADD.Enum(),
	}
	ctx := context.Background()
	cuRsp, err := proxy.Caculate(ctx, req2, opts...)
	if err != nil {
		log.ErrorContextf(ctx, "Caculate fail, err:%v", err)
	}
	msg := fmt.Sprintf("%g + %g = %g", a, b, cuRsp.GetAns())
	fmt.Println(msg)
}