package main

import (
	"fmt"
	"strconv"
	"time"
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

func (*server) HelloStreaming(req *hellopb.HelloStreamingRequest, stream hellopb.HelloService_HelloStreamingServer) error {
	fmt.Printf("Streaming a request: '%v'\n", req)
	
	name := req.GetHello().GetName()
	for i := 0; i < 10; i++ {
		result := "Hello, " + name + " for " + strconv.Itoa(i + 1) + ". time!"
		res := &hellopb.HelloStreamingResponse {
			Result: result,
		}
		stream.Send(res)
		time.Sleep(time.Second)
	}	
	return nil
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