// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: greet.proto

package greet

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GreetClient is the client API for Greet service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreetClient interface {
	Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type greetClient struct {
	cc grpc.ClientConnInterface
}

func NewGreetClient(cc grpc.ClientConnInterface) GreetClient {
	return &greetClient{cc}
}

func (c *greetClient) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/greet.Greet/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreetServer is the server API for Greet service.
// All implementations must embed UnimplementedGreetServer
// for forward compatibility
type GreetServer interface {
	Ping(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedGreetServer()
}

// UnimplementedGreetServer must be embedded to have forward compatible implementations.
type UnimplementedGreetServer struct {
}

func (UnimplementedGreetServer) Ping(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedGreetServer) mustEmbedUnimplementedGreetServer() {}

// UnsafeGreetServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreetServer will
// result in compilation errors.
type UnsafeGreetServer interface {
	mustEmbedUnimplementedGreetServer()
}

func RegisterGreetServer(s grpc.ServiceRegistrar, srv GreetServer) {
	s.RegisterService(&Greet_ServiceDesc, srv)
}

func _Greet_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greet.Greet/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetServer).Ping(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Greet_ServiceDesc is the grpc.ServiceDesc for Greet service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Greet_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "greet.Greet",
	HandlerType: (*GreetServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Greet_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "greet.proto",
}
