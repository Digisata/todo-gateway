// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.21.12
// source: social/social_service.proto

package social

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	SocialService_CreateSocial_FullMethodName = "/proto.SocialService/CreateSocial"
	SocialService_GetSocial_FullMethodName    = "/proto.SocialService/GetSocial"
	SocialService_GetAllSocial_FullMethodName = "/proto.SocialService/GetAllSocial"
	SocialService_UpdateSocial_FullMethodName = "/proto.SocialService/UpdateSocial"
	SocialService_DeleteSocial_FullMethodName = "/proto.SocialService/DeleteSocial"
)

// SocialServiceClient is the client API for SocialService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SocialServiceClient interface {
	CreateSocial(ctx context.Context, in *CreateSocialRequest, opts ...grpc.CallOption) (*SocialBaseResponse, error)
	GetSocial(ctx context.Context, in *GetSocialByIDRequest, opts ...grpc.CallOption) (*SocialBaseResponse, error)
	GetAllSocial(ctx context.Context, in *GetAllSocialRequest, opts ...grpc.CallOption) (*GetAllSocialResponse, error)
	UpdateSocial(ctx context.Context, in *UpdateSocialByIDRequest, opts ...grpc.CallOption) (*SocialBaseResponse, error)
	DeleteSocial(ctx context.Context, in *DeleteSocialByIDRequest, opts ...grpc.CallOption) (*SocialBaseResponse, error)
}

type socialServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSocialServiceClient(cc grpc.ClientConnInterface) SocialServiceClient {
	return &socialServiceClient{cc}
}

func (c *socialServiceClient) CreateSocial(ctx context.Context, in *CreateSocialRequest, opts ...grpc.CallOption) (*SocialBaseResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SocialBaseResponse)
	err := c.cc.Invoke(ctx, SocialService_CreateSocial_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *socialServiceClient) GetSocial(ctx context.Context, in *GetSocialByIDRequest, opts ...grpc.CallOption) (*SocialBaseResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SocialBaseResponse)
	err := c.cc.Invoke(ctx, SocialService_GetSocial_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *socialServiceClient) GetAllSocial(ctx context.Context, in *GetAllSocialRequest, opts ...grpc.CallOption) (*GetAllSocialResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllSocialResponse)
	err := c.cc.Invoke(ctx, SocialService_GetAllSocial_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *socialServiceClient) UpdateSocial(ctx context.Context, in *UpdateSocialByIDRequest, opts ...grpc.CallOption) (*SocialBaseResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SocialBaseResponse)
	err := c.cc.Invoke(ctx, SocialService_UpdateSocial_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *socialServiceClient) DeleteSocial(ctx context.Context, in *DeleteSocialByIDRequest, opts ...grpc.CallOption) (*SocialBaseResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SocialBaseResponse)
	err := c.cc.Invoke(ctx, SocialService_DeleteSocial_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SocialServiceServer is the server API for SocialService service.
// All implementations must embed UnimplementedSocialServiceServer
// for forward compatibility
type SocialServiceServer interface {
	CreateSocial(context.Context, *CreateSocialRequest) (*SocialBaseResponse, error)
	GetSocial(context.Context, *GetSocialByIDRequest) (*SocialBaseResponse, error)
	GetAllSocial(context.Context, *GetAllSocialRequest) (*GetAllSocialResponse, error)
	UpdateSocial(context.Context, *UpdateSocialByIDRequest) (*SocialBaseResponse, error)
	DeleteSocial(context.Context, *DeleteSocialByIDRequest) (*SocialBaseResponse, error)
	mustEmbedUnimplementedSocialServiceServer()
}

// UnimplementedSocialServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSocialServiceServer struct {
}

func (UnimplementedSocialServiceServer) CreateSocial(context.Context, *CreateSocialRequest) (*SocialBaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSocial not implemented")
}
func (UnimplementedSocialServiceServer) GetSocial(context.Context, *GetSocialByIDRequest) (*SocialBaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSocial not implemented")
}
func (UnimplementedSocialServiceServer) GetAllSocial(context.Context, *GetAllSocialRequest) (*GetAllSocialResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllSocial not implemented")
}
func (UnimplementedSocialServiceServer) UpdateSocial(context.Context, *UpdateSocialByIDRequest) (*SocialBaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSocial not implemented")
}
func (UnimplementedSocialServiceServer) DeleteSocial(context.Context, *DeleteSocialByIDRequest) (*SocialBaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSocial not implemented")
}
func (UnimplementedSocialServiceServer) mustEmbedUnimplementedSocialServiceServer() {}

// UnsafeSocialServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SocialServiceServer will
// result in compilation errors.
type UnsafeSocialServiceServer interface {
	mustEmbedUnimplementedSocialServiceServer()
}

func RegisterSocialServiceServer(s grpc.ServiceRegistrar, srv SocialServiceServer) {
	s.RegisterService(&SocialService_ServiceDesc, srv)
}

func _SocialService_CreateSocial_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSocialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SocialServiceServer).CreateSocial(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SocialService_CreateSocial_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SocialServiceServer).CreateSocial(ctx, req.(*CreateSocialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SocialService_GetSocial_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSocialByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SocialServiceServer).GetSocial(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SocialService_GetSocial_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SocialServiceServer).GetSocial(ctx, req.(*GetSocialByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SocialService_GetAllSocial_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllSocialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SocialServiceServer).GetAllSocial(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SocialService_GetAllSocial_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SocialServiceServer).GetAllSocial(ctx, req.(*GetAllSocialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SocialService_UpdateSocial_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSocialByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SocialServiceServer).UpdateSocial(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SocialService_UpdateSocial_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SocialServiceServer).UpdateSocial(ctx, req.(*UpdateSocialByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SocialService_DeleteSocial_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSocialByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SocialServiceServer).DeleteSocial(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SocialService_DeleteSocial_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SocialServiceServer).DeleteSocial(ctx, req.(*DeleteSocialByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SocialService_ServiceDesc is the grpc.ServiceDesc for SocialService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SocialService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.SocialService",
	HandlerType: (*SocialServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSocial",
			Handler:    _SocialService_CreateSocial_Handler,
		},
		{
			MethodName: "GetSocial",
			Handler:    _SocialService_GetSocial_Handler,
		},
		{
			MethodName: "GetAllSocial",
			Handler:    _SocialService_GetAllSocial_Handler,
		},
		{
			MethodName: "UpdateSocial",
			Handler:    _SocialService_UpdateSocial_Handler,
		},
		{
			MethodName: "DeleteSocial",
			Handler:    _SocialService_DeleteSocial_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "social/social_service.proto",
}