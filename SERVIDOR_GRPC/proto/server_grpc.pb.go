// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.3
// source: proto/server.proto

package proto

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

const (
	Json_Connection_FullMethodName = "/proto.Json/Connection"
)

// JsonClient is the client API for Json service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JsonClient interface {
	Connection(ctx context.Context, in *JsonRequest, opts ...grpc.CallOption) (*JsonResponse, error)
}

type jsonClient struct {
	cc grpc.ClientConnInterface
}

func NewJsonClient(cc grpc.ClientConnInterface) JsonClient {
	return &jsonClient{cc}
}

func (c *jsonClient) Connection(ctx context.Context, in *JsonRequest, opts ...grpc.CallOption) (*JsonResponse, error) {
	out := new(JsonResponse)
	err := c.cc.Invoke(ctx, Json_Connection_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JsonServer is the server API for Json service.
// All implementations must embed UnimplementedJsonServer
// for forward compatibility
type JsonServer interface {
	Connection(context.Context, *JsonRequest) (*JsonResponse, error)
	mustEmbedUnimplementedJsonServer()
}

// UnimplementedJsonServer must be embedded to have forward compatible implementations.
type UnimplementedJsonServer struct {
}

func (UnimplementedJsonServer) Connection(context.Context, *JsonRequest) (*JsonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Connection not implemented")
}
func (UnimplementedJsonServer) mustEmbedUnimplementedJsonServer() {}

// UnsafeJsonServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JsonServer will
// result in compilation errors.
type UnsafeJsonServer interface {
	mustEmbedUnimplementedJsonServer()
}

func RegisterJsonServer(s grpc.ServiceRegistrar, srv JsonServer) {
	s.RegisterService(&Json_ServiceDesc, srv)
}

func _Json_Connection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JsonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JsonServer).Connection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Json_Connection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JsonServer).Connection(ctx, req.(*JsonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Json_ServiceDesc is the grpc.ServiceDesc for Json service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Json_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Json",
	HandlerType: (*JsonServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Connection",
			Handler:    _Json_Connection_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/server.proto",
}
