package main

import (
	"context"
	"fmt"
	"log"

	"github.com/niravhjoshi/GOgrpcProj/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Calculator client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect to server %v", err)
	}
	defer cc.Close()
	c := calculatorpb.NewCalculatorServiceClient(cc)
	DoUnaryCall(c)
	//fmt.Printf("Created Client %f ", c)

}

func DoUnaryCall(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to Unary API for Sum Example")
	req := &calculatorpb.SumReq{
		FirstNum:  40,
		SecondNum: 40,
	}

	resp, err := c.Sum(context.Background(), req)

	if err != nil {
		log.Fatalf("Error While Getting Response from API Server %v", err)
	}
	log.Printf("Response from Unary API %v", resp.SumRest)
}
