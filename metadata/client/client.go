package main

import (
	"log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"time"
	"github.com/shettyh/grpc-go-examples/helloworld"
	"google.golang.org/grpc/metadata"
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
	//Setting the client timeout
	ctx,_ := context.WithTimeout(context.Background(),time.Second*10)

	//Adding metadata to the client call
	//Create key-value pairs and create metadata ctx from it
	metaMap := make(map[string]string)
	metaMap["key1"] = "value1"
	metaInfo := metadata.New(metaMap)
	metaCtx := metadata.NewContext(ctx,metaInfo)

	resp, err := client.SayHello(metaCtx, &request)
	log.Println("Got response ", resp.Message)
	conn.Close()

}
