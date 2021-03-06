package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/grpc_basics_golang/personpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Per(ctx context.Context, req *personpb.PersonRequest) (*personpb.GreetResponse, error) {
	log.Println("Request recieved for Person")
	firstName := req.GetPerson().GetFirstName()
	result := "Hello " + firstName
	res := &personpb.PersonResponse{
		Result: result,
	}
	return res, nil
}

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
	personpb.RegisterFirstServiceServer(s, &server{})
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v ", err)
	}
}
