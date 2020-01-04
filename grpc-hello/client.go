package main

import (
	"fmt"
	"log"
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
	
	fmt.Printf("Client created: %f", c)
}