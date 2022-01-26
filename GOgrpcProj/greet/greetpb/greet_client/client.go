package main

import (
	"fmt"
	"log"

	"github.com/niravhjoshi/GOgrpcProj/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I am client")
	cc, err := grpc.Dial("ocalhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect to server %v", err)
	}
	defer cc.Close()
	c := greetpb.NewGreetServiceClient(cc)
	fmt.Printf("Created Client %f ", c)

}
