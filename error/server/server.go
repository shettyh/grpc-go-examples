package main

import (
	"errors"
	"github.com/shettyh/grpc-go-examples/error"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"log"
	"net"
)

func main() {
	srvc := Service{}

	srv := grpc.NewServer(ServerInterceptor())
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
	actualError := errors.New("invalid request")
	err := errorservice.Errorf(codes.Internal, 500, false, "RPC failed", actualError)
	return nil, err
}

// ServerInterceptor will return the serverInterceptor as ServerOption
func ServerInterceptor() grpc.ServerOption {
	return grpc.UnaryInterceptor(serverInterceptor)
}

// serverInterceptor will intercept the all the grpc unary calls and add the error to grpc trailer
func serverInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if info == nil {
		return nil, errors.New("passed nil *grpc.UnaryServerInfo")
	}
	resp, err := handler(ctx, req)
	// Marshall error and add to ctx
	err = errorservice.MarshalError(err, ctx)
	return resp, err
}
