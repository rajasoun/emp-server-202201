// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.2
// source: api/service.proto

package api

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

// SimpleClient is the client API for Simple service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SimpleClient interface {
	Greeting(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (*GreetingResponse, error)
}

type simpleClient struct {
	cc grpc.ClientConnInterface
}

func NewSimpleClient(cc grpc.ClientConnInterface) SimpleClient {
	return &simpleClient{cc}
}

func (c *simpleClient) Greeting(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (*GreetingResponse, error) {
	out := new(GreetingResponse)
	err := c.cc.Invoke(ctx, "/Simple/Greeting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SimpleServer is the server API for Simple service.
// All implementations must embed UnimplementedSimpleServer
// for forward compatibility
type SimpleServer interface {
	Greeting(context.Context, *HelloReq) (*GreetingResponse, error)
	mustEmbedUnimplementedSimpleServer()
}

// UnimplementedSimpleServer must be embedded to have forward compatible implementations.
type UnimplementedSimpleServer struct {
}

func (UnimplementedSimpleServer) Greeting(context.Context, *HelloReq) (*GreetingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Greeting not implemented")
}
func (UnimplementedSimpleServer) mustEmbedUnimplementedSimpleServer() {}

// UnsafeSimpleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SimpleServer will
// result in compilation errors.
type UnsafeSimpleServer interface {
	mustEmbedUnimplementedSimpleServer()
}

func RegisterSimpleServer(s grpc.ServiceRegistrar, srv SimpleServer) {
	s.RegisterService(&Simple_ServiceDesc, srv)
}

func _Simple_Greeting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimpleServer).Greeting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Simple/Greeting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimpleServer).Greeting(ctx, req.(*HelloReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Simple_ServiceDesc is the grpc.ServiceDesc for Simple service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Simple_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Simple",
	HandlerType: (*SimpleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Greeting",
			Handler:    _Simple_Greeting_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/service.proto",
}