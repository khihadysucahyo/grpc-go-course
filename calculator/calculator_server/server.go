package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/khihadysucahyo/grpc-go-course/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, in *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	log.Printf("Sum function was invoked with %v\n", in)

	firstNumber := in.Sum.GetFirstNumber()
	secondNumber := in.Sum.GetSecondNumber()

	result := firstNumber + secondNumber

	res := &calculatorpb.SumResponse{
		Result: result,
	}

	return res, nil
}

func main() {
	fmt.Println("Starting Calculator server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}
