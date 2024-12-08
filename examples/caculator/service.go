package main

import (
	"context"

	cupb "github.com/Yamon955/Learn/examples/caculator/pb"
	"google.golang.org/protobuf/proto"
	"trpc.group/trpc-go/trpc-go/errs"
)



type caculatorImpl struct {}

func (c *caculatorImpl) Caculate(ctx context.Context, req *cupb.CaculateReq) (*cupb.CaculateRsp, error){
	rsp := &cupb.CaculateRsp{}
	switch req.GetOp() {
	case cupb.Operators_ADD:
		rsp.Ans = proto.Float64(req.GetA() + req.GetB())
	case cupb.Operators_SUB:
		rsp.Ans = proto.Float64(req.GetA() - req.GetB())
	case cupb.Operators_MUL:
		rsp.Ans = proto.Float64(req.GetA() * req.GetB())
	case cupb.Operators_DIV:
		if req.GetB() == 0 {
			return nil, errs.New(10001, "divisor isn't zero")
		}
		rsp.Ans = proto.Float64(req.GetA() / req.GetB())
	default:
		return nil, errs.New(10002, "invalid req operator")
	}
	return rsp, nil
}