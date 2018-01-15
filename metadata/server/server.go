package main

import (
	"log"
	"net"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"github.com/shettyh/grpc-go-examples/helloworld"
	"google.golang.org/grpc/metadata"
	"github.com/pkg/errors"
)

func main() {
	log.Println("Staring server...")
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
	log.Println("Say Hello called")
	//Get metadata from context
	metaInfo,ok := metadata.FromContext(ctx)

	if !ok {
		return nil, errors.New("not able to read the metadata information")
	}

	value := metaInfo["key1"]
	log.Println("Got the meta value ",value)

	response := &helloworld.HelloResponse{Message: "Hello"}
	return response, nil
}
