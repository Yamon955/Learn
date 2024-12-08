package main

import (
	"trpc.group/trpc-go/trpc-go"
	cupb "github.com/Yamon955/Learn/examples/caculator/pb"
)

func main() {
	s := trpc.NewServer()
	cupb.RegisterCaculatorService(s, &caculatorImpl{})
	if err := s.Serve(); err != nil {
		panic(err)
	}
}