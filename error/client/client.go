package main

import (
	"context"
	"github.com/shettyh/grpc-go-examples/error"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:10001", grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
	}

	client := errorservice.NewErrorServiceClient(conn)

	// Fetch trailer from request
	var trailer metadata.MD
	ctx := context.Background()
	_, err = client.TestError(ctx, &errorservice.Request{}, grpc.Trailer(&trailer))

	// Unmarshal error message
	errMsg := errorservice.UnmarshalError(err, trailer)

	log.Println(errMsg.Message)
}
