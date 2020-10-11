package main

import (
	"log"

	"github.com/grpc_basics_golang/firstpb"
	"google.golang.org/grpc"
)

func main() {
	// initiate a connection request to the server address and port
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("falied to start client %v ", err)
	}
	defer cc.Close()
	// register client to the proto file object
	c := firstpb.NewFirstServiceClient(cc)
	log.Println("new client connection is ", c)
}
