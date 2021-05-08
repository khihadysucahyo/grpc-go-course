package main

import (
	"context"
	"fmt"
	"log"

	"github.com/khihadysucahyo/grpc-go-course/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Starting Calculator Client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect %v", err)
	}

	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)

	doUnary(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	in := &calculatorpb.SumRequest{
		Sum: &calculatorpb.Sum{
			FirstNumber:  10,
			SecondNumber: 3,
		},
	}

	res, err := c.Sum(context.Background(), in)

	if err != nil {
		log.Fatalf("Error while calling Greet RPC with %v", err)
	}

	log.Printf("Response from Greet: %v", res.Result)
}
