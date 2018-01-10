package main

import (
	"log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"github.com/shettyh/grpc-go-examples/helloworld"
)

func main() {
	conn, err := grpc.Dial("localhost:17002", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Failed to connect to server")
	}

	client := helloworld.NewHelloWorldServiceClient(conn)

	if err != nil {
		log.Fatalf("Failed to connect to server")
	}

	request := helloworld.HelloRequest{Message: "Hi"}

	resp, err := client.SayHello(context.Background(), &request)

	log.Println("Got response ", resp.Message)

	conn.Close()

}
