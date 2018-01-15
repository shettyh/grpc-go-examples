package main

import (
	"github.com/shettyh/grpc-go-examples/interceptor"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:17002", grpc.WithInsecure(), WithClientInterceptor())

	if err != nil {
		log.Fatalf("Failed to connect to server")
	}

	client := interceptor.NewTestServiceClient(conn)

	if err != nil {
		log.Fatalf("Failed to connect to server")
	}

	request := interceptor.HelloRequest{Message: "Hi"}

	resp, err := client.SayHello(context.Background(), &request)

	log.Println("Got response ", resp.Message)

	conn.Close()

}

func WithClientInterceptor() grpc.DialOption {
	return grpc.WithUnaryInterceptor(clientInterceptor)
}

func clientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	log.Println("Intercepting the remote call ",method)
	start := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...) // <==
	log.Printf("invoke remote method=%s duration=%s error=%v", method, time.Since(start), err)
	return err
}
