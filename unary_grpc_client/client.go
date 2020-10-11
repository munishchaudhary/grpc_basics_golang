package main

import (
	"context"
	"log"

	personpb "github.com/grpc_basics_golang/unarypb"

	"google.golang.org/grpc"
)

func main() {
	// initiate a connection request to the server address and port
	log.Println("starting grpc client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("falied to start client %v ", err)
	}
	defer cc.Close()
	// register client to the proto file object
	c := personpb.NewPersonServiceClient(cc)

	// creating Request message for the GetFirstName rpc
	req := &personpb.PersonRequest{
		Person: &personpb.Person{
			FirstName: "Munish",
			LastName:  "Chaudhary",
		},
	}
	// sending GetFirstName request with request message
	res, err := c.GetFirstName(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while sending GetFirstName Request %v ", err)
	}
	log.Printf("GetFirstName response is %v ", res.Result)
}
