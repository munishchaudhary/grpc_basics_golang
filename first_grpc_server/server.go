package main

import (
	"fmt"
	"log"
	"net"

	"github.com/grpc_basics_golang/firstpb"
	"google.golang.org/grpc"
)

type server struct{}

func main() {
	fmt.Println("Starting grpc server ...")
	// bind host ip address and port with the server
	l, e := net.Listen("tcp", "localhost:50051")
	if e != nil {
		log.Fatalf("Error while listing: %v ", e)
	}
	// create a new grpc server object
	s := grpc.NewServer()
	// register grpc server with the proto file
	firstpb.RegisterFirstServiceServer(s, &server{})
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v ", err)
	}
}
