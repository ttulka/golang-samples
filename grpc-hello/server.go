package main

import (
	"fmt"
	"net"
	"log"
	"context"
	"google.golang.org/grpc"
	"github.com/ttulka/golang-samples/grpc-hello/hellopb"
)

type server struct {
}

func (*server) Hello(ctx context.Context, req *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {
	fmt.Printf("Serving a request: '%v'\n", req)
	
	name := req.GetHello().GetName()
	result := "Hello, " + name + "!"
	res := &hellopb.HelloResponse {
		Result: result,
	}
	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	
	s := grpc.NewServer()
	hellopb.RegisterHelloServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server: '%v'", err)
	}
}