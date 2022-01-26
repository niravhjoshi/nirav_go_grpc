package main

import (
	"context"
	"fmt"
	"log"

	"github.com/niravhjoshi/GOgrpcProj/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I am client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect to server %v", err)
	}
	defer cc.Close()
	c := greetpb.NewGreetServiceClient(cc)
	DoUnaryCall(c)
	//fmt.Printf("Created Client %f ", c)

}

func DoUnaryCall(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to Unary API")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Nirav",
			LastName:  "Joshi",
		},
	}
	resp, err := c.Greet(context.Background(), req)

	if err != nil {
		log.Fatalf("Error While Getting Response from API Server %v", err)
	}
	log.Printf("Response from Unary API %v", resp.Result)
}
