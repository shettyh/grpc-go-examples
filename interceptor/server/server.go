package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"github.com/shettyh/grpc-go-examples/interceptor"
	"time"
)

func main() {
	fmt.Println("Staring server...")
	srv := TestServiceImpl{}
	lis, err := net.Listen("tcp", ":17002")

	if err != nil {
		log.Fatalf("Failed to bind to port %v", err)
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(serverInterceptor))

	interceptor.RegisterTestServiceServer(s, &srv)

	log.Println("Starting the server ...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to start server")
	}
}

type TestServiceImpl struct {
}

func (*TestServiceImpl) SayHello(ctx context.Context, in *interceptor.HelloRequest) (*interceptor.HelloResponse, error) {
	fmt.Println("Say Hello called")
	response := &interceptor.HelloResponse{Message: "Hello"}
	return response, nil
}

func serverInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Println("Inside interceptor")
	log.Println("Method called %s", info.FullMethod)
	if info == nil {
		return nil, errors.New("passed nil *grpc.UnaryServerInfo")
	}

	resp, err := handler(ctx, req)
	return resp, err
}
