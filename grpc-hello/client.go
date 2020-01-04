package main

import (
	"log"
	"io"
	"context"
	"google.golang.org/grpc"
	"github.com/ttulka/golang-samples/grpc-hello/hellopb"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()
	
	c := hellopb.NewHelloServiceClient(conn)	
	
	requestUnary(c)	
	requestStreaming(c)
}

func requestUnary(c hellopb.HelloServiceClient) {
	req := &hellopb.HelloRequest {
		Hello: &hellopb.Hello {
			Name: "Tomas",
		},
	}
	res, err := c.Hello(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling: %v", err)
	}
	log.Printf("Response from the server: '%v'", res.Result)
}

func requestStreaming(c hellopb.HelloServiceClient) {
	req := &hellopb.HelloStreamingRequest {
		Hello: &hellopb.Hello {
			Name: "Tomas",
		},
	}
	resStream, err := c.HelloStreaming(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling: %v", err)
	}
	for {
		res, err := resStream.Recv()
		if err == io.EOF {	
			// we've reached the end of the stream
			break
		}
		if err != nil {
			log.Fatalf("Error while reading a stream: %v", err)
		}
		log.Printf("Response from the server: '%v'", res.Result)
	}
}