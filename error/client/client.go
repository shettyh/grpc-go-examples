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


	var trailer metadata.MD
	client := errorservice.NewErrorServiceClient(conn)
	ctx := context.Background()
	_,err = client.TestError(ctx,&errorservice.Request{},grpc.Trailer(&trailer))

	errMsg := errorservice.UnmarshalError(err,trailer)

	log.Println(errMsg.Message)
}