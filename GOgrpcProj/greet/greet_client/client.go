package main

import (
	"context"
	"fmt"
	"io"
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
	//Do unary call
	DoUnaryCall(c)
	//Do Streaming Call function
	ServerStreamingCall(c)
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

func ServerStreamingCall(c greetpb.GreetServiceClient) {
	fmt.Println("Starting Streaming API for Server Streaming API Data sending")
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Darth",
			LastName:  "Vader",
		},
	}
	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error While Sending request to server via grpc : %v\n", err)
	}

	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			// we have reached end of stream
			break
		}

		if err != nil {
			log.Fatalf("Error While reading stream %v\n", err)
		}
		log.Printf("Response from Greet ManyTime API Server Streaming op %v\n", msg.GetResult())
	}

}
