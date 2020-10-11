package main

import (
	"context"
	"log"
	"net"

	personpb "github.com/grpc_basics_golang/unarypb"
	"google.golang.org/grpc"
)

type server struct{}

/**
** GetFirstName rpc request reciever function
**/
func (*server) GetFirstName(ctx context.Context, req *personpb.PersonRequest) (*personpb.PersonResponse, error) {
	log.Println("Request recieved for Person")
	firstName := req.GetPerson().GetFirstName()
	result := "First Name: " + firstName
	res := &personpb.PersonResponse{
		Result: result,
	}
	return res, nil

}
func main() {
	log.Println("Starting grpc server ...")
	// bind host ip address and port with the server
	l, e := net.Listen("tcp", "localhost:50051")
	if e != nil {
		log.Fatalf("Error while listing: %v ", e)
	}
	// create a new grpc server object
	s := grpc.NewServer()
	// register grpc server with the proto file
	personpb.RegisterPersonServiceServer(s, &server{})
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v ", err)
	}
}
