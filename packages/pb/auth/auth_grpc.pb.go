// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.2
// source: packages/proto/auth.proto

package auth

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

// IsAuthClient is the client API for IsAuth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IsAuthClient interface {
	IsAuth(ctx context.Context, in *IsAuthRequest, opts ...grpc.CallOption) (*IsAuthResponse, error)
}

type isAuthClient struct {
	cc grpc.ClientConnInterface
}

func NewIsAuthClient(cc grpc.ClientConnInterface) IsAuthClient {
	return &isAuthClient{cc}
}

func (c *isAuthClient) IsAuth(ctx context.Context, in *IsAuthRequest, opts ...grpc.CallOption) (*IsAuthResponse, error) {
	out := new(IsAuthResponse)
	err := c.cc.Invoke(ctx, "/IsAuth/IsAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IsAuthServer is the server API for IsAuth service.
// All implementations must embed UnimplementedIsAuthServer
// for forward compatibility
type IsAuthServer interface {
	IsAuth(context.Context, *IsAuthRequest) (*IsAuthResponse, error)
	mustEmbedUnimplementedIsAuthServer()
}

// UnimplementedIsAuthServer must be embedded to have forward compatible implementations.
type UnimplementedIsAuthServer struct {
}

func (UnimplementedIsAuthServer) IsAuth(context.Context, *IsAuthRequest) (*IsAuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsAuth not implemented")
}
func (UnimplementedIsAuthServer) mustEmbedUnimplementedIsAuthServer() {}

// UnsafeIsAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IsAuthServer will
// result in compilation errors.
type UnsafeIsAuthServer interface {
	mustEmbedUnimplementedIsAuthServer()
}

func RegisterIsAuthServer(s grpc.ServiceRegistrar, srv IsAuthServer) {
	s.RegisterService(&IsAuth_ServiceDesc, srv)
}

func _IsAuth_IsAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IsAuthServer).IsAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/IsAuth/IsAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IsAuthServer).IsAuth(ctx, req.(*IsAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// IsAuth_ServiceDesc is the grpc.ServiceDesc for IsAuth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IsAuth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "IsAuth",
	HandlerType: (*IsAuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsAuth",
			Handler:    _IsAuth_IsAuth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "packages/proto/auth.proto",
}
