// Code generated by protoc-gen-go. DO NOT EDIT.
// source: helloworld.proto

/*
Package helloworld is a generated protocol buffer package.

It is generated from these files:
	helloworld.proto

It has these top-level messages:
	HelloRequest
	HelloResponse
*/
package helloworld

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type HelloRequest struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *HelloRequest) Reset()                    { *m = HelloRequest{} }
func (m *HelloRequest) String() string            { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()               {}
func (*HelloRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HelloRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type HelloResponse struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *HelloResponse) Reset()                    { *m = HelloResponse{} }
func (m *HelloResponse) String() string            { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()               {}
func (*HelloResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HelloResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "helloworld.HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "helloworld.HelloResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for HelloWorldService service

type HelloWorldServiceClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
}

type helloWorldServiceClient struct {
	cc *grpc.ClientConn
}

func NewHelloWorldServiceClient(cc *grpc.ClientConn) HelloWorldServiceClient {
	return &helloWorldServiceClient{cc}
}

func (c *helloWorldServiceClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := grpc.Invoke(ctx, "/helloworld.HelloWorldService/SayHello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for HelloWorldService service

type HelloWorldServiceServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloResponse, error)
}

func RegisterHelloWorldServiceServer(s *grpc.Server, srv HelloWorldServiceServer) {
	s.RegisterService(&_HelloWorldService_serviceDesc, srv)
}

func _HelloWorldService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloWorldServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.HelloWorldService/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloWorldServiceServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HelloWorldService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.HelloWorldService",
	HandlerType: (*HelloWorldServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _HelloWorldService_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "helloworld.proto",
}

func init() { proto.RegisterFile("helloworld.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 138 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc8, 0x48, 0xcd, 0xc9,
	0xc9, 0x2f, 0xcf, 0x2f, 0xca, 0x49, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x42, 0x88,
	0x28, 0x69, 0x70, 0xf1, 0x78, 0x80, 0x78, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x12,
	0x5c, 0xec, 0xb9, 0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41,
	0x30, 0xae, 0x92, 0x26, 0x17, 0x2f, 0x54, 0x65, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0x2a, 0x6e, 0xa5,
	0x46, 0x21, 0x5c, 0x82, 0x60, 0xa5, 0xe1, 0x20, 0x2b, 0x82, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53,
	0x85, 0xec, 0xb9, 0x38, 0x82, 0x13, 0x2b, 0xc1, 0xe2, 0x42, 0x12, 0x7a, 0x48, 0x8e, 0x42, 0xb6,
	0x5f, 0x4a, 0x12, 0x8b, 0x0c, 0xc4, 0xbe, 0x24, 0x36, 0xb0, 0xeb, 0x8d, 0x01, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x1f, 0xb6, 0x71, 0xca, 0xd1, 0x00, 0x00, 0x00,
}
