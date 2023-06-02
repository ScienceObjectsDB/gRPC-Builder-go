// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: aruna/api/internal/v1/authorize.proto

package v1

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

// InternalAuthorizeServiceClient is the client API for InternalAuthorizeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InternalAuthorizeServiceClient interface {
	Authorize(ctx context.Context, in *AuthorizeRequest, opts ...grpc.CallOption) (*AuthorizeResponse, error)
	GetSecret(ctx context.Context, in *GetSecretRequest, opts ...grpc.CallOption) (*GetSecretResponse, error)
	GetTokenFromSecret(ctx context.Context, in *GetTokenFromSecretRequest, opts ...grpc.CallOption) (*GetTokenFromSecretResponse, error)
}

type internalAuthorizeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInternalAuthorizeServiceClient(cc grpc.ClientConnInterface) InternalAuthorizeServiceClient {
	return &internalAuthorizeServiceClient{cc}
}

func (c *internalAuthorizeServiceClient) Authorize(ctx context.Context, in *AuthorizeRequest, opts ...grpc.CallOption) (*AuthorizeResponse, error) {
	out := new(AuthorizeResponse)
	err := c.cc.Invoke(ctx, "/aruna.api.internal.v1.InternalAuthorizeService/Authorize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalAuthorizeServiceClient) GetSecret(ctx context.Context, in *GetSecretRequest, opts ...grpc.CallOption) (*GetSecretResponse, error) {
	out := new(GetSecretResponse)
	err := c.cc.Invoke(ctx, "/aruna.api.internal.v1.InternalAuthorizeService/GetSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalAuthorizeServiceClient) GetTokenFromSecret(ctx context.Context, in *GetTokenFromSecretRequest, opts ...grpc.CallOption) (*GetTokenFromSecretResponse, error) {
	out := new(GetTokenFromSecretResponse)
	err := c.cc.Invoke(ctx, "/aruna.api.internal.v1.InternalAuthorizeService/GetTokenFromSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InternalAuthorizeServiceServer is the server API for InternalAuthorizeService service.
// All implementations should embed UnimplementedInternalAuthorizeServiceServer
// for forward compatibility
type InternalAuthorizeServiceServer interface {
	Authorize(context.Context, *AuthorizeRequest) (*AuthorizeResponse, error)
	GetSecret(context.Context, *GetSecretRequest) (*GetSecretResponse, error)
	GetTokenFromSecret(context.Context, *GetTokenFromSecretRequest) (*GetTokenFromSecretResponse, error)
}

// UnimplementedInternalAuthorizeServiceServer should be embedded to have forward compatible implementations.
type UnimplementedInternalAuthorizeServiceServer struct {
}

func (UnimplementedInternalAuthorizeServiceServer) Authorize(context.Context, *AuthorizeRequest) (*AuthorizeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authorize not implemented")
}
func (UnimplementedInternalAuthorizeServiceServer) GetSecret(context.Context, *GetSecretRequest) (*GetSecretResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSecret not implemented")
}
func (UnimplementedInternalAuthorizeServiceServer) GetTokenFromSecret(context.Context, *GetTokenFromSecretRequest) (*GetTokenFromSecretResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTokenFromSecret not implemented")
}

// UnsafeInternalAuthorizeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InternalAuthorizeServiceServer will
// result in compilation errors.
type UnsafeInternalAuthorizeServiceServer interface {
	mustEmbedUnimplementedInternalAuthorizeServiceServer()
}

func RegisterInternalAuthorizeServiceServer(s grpc.ServiceRegistrar, srv InternalAuthorizeServiceServer) {
	s.RegisterService(&InternalAuthorizeService_ServiceDesc, srv)
}

func _InternalAuthorizeService_Authorize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalAuthorizeServiceServer).Authorize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aruna.api.internal.v1.InternalAuthorizeService/Authorize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalAuthorizeServiceServer).Authorize(ctx, req.(*AuthorizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalAuthorizeService_GetSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalAuthorizeServiceServer).GetSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aruna.api.internal.v1.InternalAuthorizeService/GetSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalAuthorizeServiceServer).GetSecret(ctx, req.(*GetSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalAuthorizeService_GetTokenFromSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTokenFromSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalAuthorizeServiceServer).GetTokenFromSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aruna.api.internal.v1.InternalAuthorizeService/GetTokenFromSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalAuthorizeServiceServer).GetTokenFromSecret(ctx, req.(*GetTokenFromSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InternalAuthorizeService_ServiceDesc is the grpc.ServiceDesc for InternalAuthorizeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InternalAuthorizeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "aruna.api.internal.v1.InternalAuthorizeService",
	HandlerType: (*InternalAuthorizeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authorize",
			Handler:    _InternalAuthorizeService_Authorize_Handler,
		},
		{
			MethodName: "GetSecret",
			Handler:    _InternalAuthorizeService_GetSecret_Handler,
		},
		{
			MethodName: "GetTokenFromSecret",
			Handler:    _InternalAuthorizeService_GetTokenFromSecret_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "aruna/api/internal/v1/authorize.proto",
}
