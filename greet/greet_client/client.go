package main

import (
	"context"
	"fmt"
	"log"

	"github.com/khihadysucahyo/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I'm a client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect %v", err)
	}

	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	// fmt.Printf("Created client: %v", c)

	doUnary(c)

}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Unary RPC...")

	in := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Khihady",
			LastName:  "Sucahyo",
		},
	}

	res, err := c.Greet(context.Background(), in)

	if err != nil {
		log.Fatalf("Error while calling Greet RPC with %v", err)
	}

	log.Printf("Response from Greet: %v", res.Result)
}
