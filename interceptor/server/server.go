package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"github.com/shettyh/grpc-go-examples/interceptor"
)

func main() {
	fmt.Println("Staring server...")
	srv := TestServiceImpl{}
	lis, err := net.Listen("tcp", ":17002")

	if err != nil {
		log.Fatalf("Failed to bind to port %v", err)
	}

	s := grpc.NewServer(ServerInterceptor())

	interceptor.RegisterTestServiceServer(s, &srv)

	log.Println("Starting the server ...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to start server")
	}
}

type TestServiceImpl struct {}

func (*TestServiceImpl) SayHello(ctx context.Context, in *interceptor.HelloRequest) (*interceptor.HelloResponse, error) {
	log.Println("RPC called SayHello")
	response := &interceptor.HelloResponse{Message: "Hello"}
	return response, nil
}

// ServerInterceptor will return the serverInterceptor as ServerOption
func ServerInterceptor() grpc.ServerOption {
	return grpc.UnaryInterceptor(serverInterceptor)
}

// serverInterceptor will intercept the all the grpc unary calls
func serverInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Println("Intercepting the call ", info.FullMethod)
	if info == nil {
		return nil, errors.New("passed nil *grpc.UnaryServerInfo")
	}
	resp, err := handler(ctx, req)
	return resp, err
}
