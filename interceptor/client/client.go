package main

import (
	"log"

	"github.com/shettyh/grpc-go-examples/interceptor"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:17002", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Failed to connect to server")
	}

	client := test.NewTestServiceClient(conn)

	if err != nil {
		log.Fatalf("Failed to connect to server")
	}

	request := test.HelloRequest{Message: "Hi"}

	resp, err := client.SayHello(context.Background(), &request)

	log.Println("Got response ", resp.Message)

	conn.Close()

}
