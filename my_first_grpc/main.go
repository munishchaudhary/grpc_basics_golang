package main

import (
	"fmt"
	"net"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Starting grpc server ...")
	l, e := net.Listen("tcp", "0.0.0.0:50051")
	if e != nil {
		fmt.Println("Error while listing ", e)
	}
	s := grpc.NewServer()
	firstpb.RegisterFirstServer()
}
