package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/niravhjoshi/GOgrpcProj/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct {
}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet Function was invoked with req %v ", req)
	firstname := req.GetGreeting().GetFirstName()
	result := "Hello" + firstname
	res := &greetpb.GreetResponse{
		Result: result,
	}
	return res, nil
}

func main() {

	fmt.Println("Hello World")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("failed to listen %v ", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to start Server: %v", err)
	}
}
