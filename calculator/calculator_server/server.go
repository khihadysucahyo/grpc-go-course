package main

import (
	"context"
	"fmt"
	"io"
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

func (*server) FindMaximum(stream calculatorpb.CalculatorService_FindMaximumServer) error {
	log.Printf("FindMaximum function was invoked")
	maximum := int32(0)

	for {
		req, err := stream.Recv()
		if err != io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}

		number := req.GetNumber()
		if number > maximum {
			maximum = number
			sendErr := stream.Send(&calculatorpb.FindMaximumResponse{
				Maximum: maximum,
			})

			if sendErr != nil {
				log.Fatalf("Error while sending data to client stream: %v", sendErr)
				return sendErr
			}
		}
	}
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
