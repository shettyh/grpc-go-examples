package main

import (
	"fmt"
	"log"
	"net"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"github.com/shettyh/grpc-go-examples/helloworld"
)

func main() {
	fmt.Println("Staring server...")
	srv := TestServiceImpl{}
	lis, err := net.Listen("tcp", ":17002")

	if err != nil {
		log.Fatalf("Failed to bind to port %v", err)
	}

	s := grpc.NewServer()

	helloworld.RegisterHelloWorldServiceServer(s, &srv)

	log.Println("Starting the server ...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to start server")
	}
}

type TestServiceImpl struct {
}

func (*TestServiceImpl) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloResponse, error) {
	fmt.Println("Say Hello called")
	response := &helloworld.HelloResponse{Message: "Hello"}
	return response, nil
}
