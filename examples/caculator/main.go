package main

import (
	cupb "github.com/Yamon955/Learn/examples/caculator/pb"
	"trpc.group/trpc-go/trpc-go"
	_ "trpc.group/trpc-go/trpc-naming-consul"
)

func main() {
	s := trpc.NewServer()
	cupb.RegisterCaculatorService(s, &caculatorImpl{})
	if err := s.Serve(); err != nil {
		panic(err)
	}
}
