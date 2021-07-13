package main

import (
	"context"
	"fmt"
	"log"

	"github.com/khihadysucahyo/grpc-go-course/blog/blogpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Blog Client")

	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("could not connect %v", err)
	}

	defer cc.Close()

	c := blogpb.NewBlogServiceClient(cc)

	// create blog
	fmt.Println("Creating Blog")
	blog := &blogpb.Blog{
		AuthorId: "Stephane",
		Title:    "My first Blog",
		Content:  "Content of the first blog",
	}

	createBlogRes, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{
		Blog: blog,
	})

	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}

	fmt.Printf("Blog has been created: %v", createBlogRes)
	blogId := createBlogRes.GetBlog().GetId()

	// read blog
	fmt.Println("Reading Blog")
	readBlogReq := &blogpb.ReadBlogRequest{BlogId: blogId}
	readBlogRes, readBlogErr := c.ReadBlog(context.Background(), readBlogReq)
	if readBlogErr != nil {
		log.Fatalf("Error happened while reading: %v", err)
	}
	fmt.Printf("Blog has been fetched: %v", readBlogRes)
}
