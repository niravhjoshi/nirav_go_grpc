package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/niravhjoshi/GOgrpcProj/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct {
}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumReq) (*calculatorpb.SumResp, error) {
	fmt.Printf("Received Sum RPC %v\n", req)
	firstNum := req.FirstNum
	secondNum := req.SecondNum

	addsum := firstNum + secondNum
	result := &calculatorpb.SumResp{
		SumRest: addsum,
	}
	return result, nil
}

func main() {

	fmt.Println("Calculator server ")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("failed to listen %v ", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to start Server: %v", err)
	}
}
