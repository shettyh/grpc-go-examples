package main

import (
	"errors"
	"log"
	"net"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"github.com/shettyh/grpc-go-examples/interceptor"
)

func main() {
	log.Println("Staring server...")
	srv := TestServiceImpl{}
	lis, err := net.Listen("tcp", ":17002")

	if err != nil {
		log.Fatalf("Failed to bind to port %v", err)
	}

	si := &serverinterceptor{}
	s := grpc.NewServer(grpc.UnaryInterceptor(si.intercept))

	test.RegisterTestServiceServer(s, &srv)

	log.Println("Starting the server ...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to start server")
	}
}

type TestServiceImpl struct {
}

func (*TestServiceImpl) SayHello(ctx context.Context, in *test.HelloRequest) (*test.HelloResponse, error) {
	log.Println("Say Hello called")
	response := &test.HelloResponse{Message: "Hello"}
	return response, nil
}

type serverinterceptor struct {
}

func (*serverinterceptor) intercept(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Println("Inside interceptor")
	log.Println("Method called", info.FullMethod)
	if info == nil {
		return nil, errors.New("passed nil *grpc.UnaryServerInfo")
	}

	log.Println("Calling the RPC")
	resp, err := handler(ctx, req)
	log.Println("Called rpc and response is ",resp)
	return resp, err
}
