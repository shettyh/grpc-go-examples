package main

import (
	"google.golang.org/grpc"
	"log"
	"github.com/shettyh/grpc-go-examples/error"
	"context"
	"google.golang.org/grpc/metadata"
)

func main(){
	conn,err := grpc.Dial("localhost:10001",grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
	}

	client := errorservice.NewErrorServiceClient(conn)
	ctx := metadata.NewOutgoingContext(context.Background(),metadata.Pairs("",""))
	_,err = client.TestError(ctx,&errorservice.Request{})
	md,_ := metadata.FromIncomingContext(ctx)
	errMsg := errorservice.UnmarshalError(err,md)

	log.Println(errMsg.Message)
}