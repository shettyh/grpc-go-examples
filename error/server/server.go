package main

import (
	"golang.org/x/net/context"
	"github.com/shettyh/grpc-go-examples/error"
	"google.golang.org/grpc/codes"
	"errors"
	"google.golang.org/grpc"
	"net"
	"log"
)

func main() {
	srvc := Service{}

	srv := grpc.NewServer()
	errorservice.RegisterErrorServiceServer(srv, srvc)

	lis, err := net.Listen("tcp", ":10001")

	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(srv.Serve(lis))
}

type Service struct{}

func (Service) TestError(ctx context.Context, request *errorservice.Request) (*errorservice.Response, error) {
	//Add actual logic here

	// On Error
	err := errorservice.Errorf(codes.Internal, 500, false, "RPC failed")
	return nil, err
}

// ServerInterceptor will return the serverInterceptor as ServerOption
func ServerInterceptor() grpc.ServerOption {
	return grpc.UnaryInterceptor(serverInterceptor)
}

// serverInterceptor will intercept the all the grpc unary calls
func serverInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if info == nil {
		return nil, errors.New("passed nil *grpc.UnaryServerInfo")
	}
	resp, err := handler(ctx, req)
	// Marshall error and add to ctx
	err = errorservice.MarshalError(err, ctx)
	return resp, err
}
