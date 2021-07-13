package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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

	// doUnary(c)
	doBiDiStreaming(c)
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

func doBiDiStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do FindMaximum BiDi Streaming RPC...")

	stream, err := c.FindMaximum(context.Background())

	if err != nil {
		log.Fatalf("Error while opening stream and calling FindMaximum: %v", err)
	}

	waitc := make(chan struct{})

	// send go routine
	go func() {
		numbers := []int32{4, 7, 2, 19, 4, 6, 32}
		for _, number := range numbers {
			fmt.Printf("Sending number: %v", number)
			stream.Send(&calculatorpb.FindMaximumRequest{
				Number: number,
			})
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	// receive go routine
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Problem while reading server stream: %v", err)
				break
			}

			maximum := res.GetMaximum()
			fmt.Printf("Received a new maximum of...: %v", maximum)
		}
		close(waitc)
	}()
	<-waitc
}
