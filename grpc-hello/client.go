package main

import (
	"log"
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
	
	req := &hellopb.HelloRequest {
		Hello: &hellopb.Hello {
			Name: "Tomas",
		},
	}
	res, err := c.Hello(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling: %v", err)
	}
	log.Printf("Response from the server: %v", res.Result)
}