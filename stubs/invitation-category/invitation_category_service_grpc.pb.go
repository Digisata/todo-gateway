// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.21.12
// source: invitation-category/invitation_category_service.proto

package invitation_category

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
	InvitationCategoryService_CreateInvitationCategory_FullMethodName = "/proto.InvitationCategoryService/CreateInvitationCategory"
	InvitationCategoryService_GetInvitationCategory_FullMethodName    = "/proto.InvitationCategoryService/GetInvitationCategory"
	InvitationCategoryService_GetAllInvitationCategory_FullMethodName = "/proto.InvitationCategoryService/GetAllInvitationCategory"
	InvitationCategoryService_UpdateInvitationCategory_FullMethodName = "/proto.InvitationCategoryService/UpdateInvitationCategory"
	InvitationCategoryService_DeleteInvitationCategory_FullMethodName = "/proto.InvitationCategoryService/DeleteInvitationCategory"
)

// InvitationCategoryServiceClient is the client API for InvitationCategoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InvitationCategoryServiceClient interface {
	CreateInvitationCategory(ctx context.Context, in *CreateInvitationCategoryRequest, opts ...grpc.CallOption) (*InvitationCategoryBaseResponse, error)
	GetInvitationCategory(ctx context.Context, in *GetInvitationCategoryByIDRequest, opts ...grpc.CallOption) (*InvitationCategoryBaseResponse, error)
	GetAllInvitationCategory(ctx context.Context, in *GetAllInvitationCategoryRequest, opts ...grpc.CallOption) (*GetAllInvitationCategoryResponse, error)
	UpdateInvitationCategory(ctx context.Context, in *UpdateInvitationCategoryByIDRequest, opts ...grpc.CallOption) (*InvitationCategoryBaseResponse, error)
	DeleteInvitationCategory(ctx context.Context, in *DeleteInvitationCategoryByIDRequest, opts ...grpc.CallOption) (*InvitationCategoryBaseResponse, error)
}

type invitationCategoryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInvitationCategoryServiceClient(cc grpc.ClientConnInterface) InvitationCategoryServiceClient {
	return &invitationCategoryServiceClient{cc}
}

func (c *invitationCategoryServiceClient) CreateInvitationCategory(ctx context.Context, in *CreateInvitationCategoryRequest, opts ...grpc.CallOption) (*InvitationCategoryBaseResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(InvitationCategoryBaseResponse)
	err := c.cc.Invoke(ctx, InvitationCategoryService_CreateInvitationCategory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *invitationCategoryServiceClient) GetInvitationCategory(ctx context.Context, in *GetInvitationCategoryByIDRequest, opts ...grpc.CallOption) (*InvitationCategoryBaseResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(InvitationCategoryBaseResponse)
	err := c.cc.Invoke(ctx, InvitationCategoryService_GetInvitationCategory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *invitationCategoryServiceClient) GetAllInvitationCategory(ctx context.Context, in *GetAllInvitationCategoryRequest, opts ...grpc.CallOption) (*GetAllInvitationCategoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllInvitationCategoryResponse)
	err := c.cc.Invoke(ctx, InvitationCategoryService_GetAllInvitationCategory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *invitationCategoryServiceClient) UpdateInvitationCategory(ctx context.Context, in *UpdateInvitationCategoryByIDRequest, opts ...grpc.CallOption) (*InvitationCategoryBaseResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(InvitationCategoryBaseResponse)
	err := c.cc.Invoke(ctx, InvitationCategoryService_UpdateInvitationCategory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *invitationCategoryServiceClient) DeleteInvitationCategory(ctx context.Context, in *DeleteInvitationCategoryByIDRequest, opts ...grpc.CallOption) (*InvitationCategoryBaseResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(InvitationCategoryBaseResponse)
	err := c.cc.Invoke(ctx, InvitationCategoryService_DeleteInvitationCategory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InvitationCategoryServiceServer is the server API for InvitationCategoryService service.
// All implementations must embed UnimplementedInvitationCategoryServiceServer
// for forward compatibility
type InvitationCategoryServiceServer interface {
	CreateInvitationCategory(context.Context, *CreateInvitationCategoryRequest) (*InvitationCategoryBaseResponse, error)
	GetInvitationCategory(context.Context, *GetInvitationCategoryByIDRequest) (*InvitationCategoryBaseResponse, error)
	GetAllInvitationCategory(context.Context, *GetAllInvitationCategoryRequest) (*GetAllInvitationCategoryResponse, error)
	UpdateInvitationCategory(context.Context, *UpdateInvitationCategoryByIDRequest) (*InvitationCategoryBaseResponse, error)
	DeleteInvitationCategory(context.Context, *DeleteInvitationCategoryByIDRequest) (*InvitationCategoryBaseResponse, error)
	mustEmbedUnimplementedInvitationCategoryServiceServer()
}

// UnimplementedInvitationCategoryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedInvitationCategoryServiceServer struct {
}

func (UnimplementedInvitationCategoryServiceServer) CreateInvitationCategory(context.Context, *CreateInvitationCategoryRequest) (*InvitationCategoryBaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateInvitationCategory not implemented")
}
func (UnimplementedInvitationCategoryServiceServer) GetInvitationCategory(context.Context, *GetInvitationCategoryByIDRequest) (*InvitationCategoryBaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInvitationCategory not implemented")
}
func (UnimplementedInvitationCategoryServiceServer) GetAllInvitationCategory(context.Context, *GetAllInvitationCategoryRequest) (*GetAllInvitationCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllInvitationCategory not implemented")
}
func (UnimplementedInvitationCategoryServiceServer) UpdateInvitationCategory(context.Context, *UpdateInvitationCategoryByIDRequest) (*InvitationCategoryBaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateInvitationCategory not implemented")
}
func (UnimplementedInvitationCategoryServiceServer) DeleteInvitationCategory(context.Context, *DeleteInvitationCategoryByIDRequest) (*InvitationCategoryBaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteInvitationCategory not implemented")
}
func (UnimplementedInvitationCategoryServiceServer) mustEmbedUnimplementedInvitationCategoryServiceServer() {
}

// UnsafeInvitationCategoryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InvitationCategoryServiceServer will
// result in compilation errors.
type UnsafeInvitationCategoryServiceServer interface {
	mustEmbedUnimplementedInvitationCategoryServiceServer()
}

func RegisterInvitationCategoryServiceServer(s grpc.ServiceRegistrar, srv InvitationCategoryServiceServer) {
	s.RegisterService(&InvitationCategoryService_ServiceDesc, srv)
}

func _InvitationCategoryService_CreateInvitationCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateInvitationCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvitationCategoryServiceServer).CreateInvitationCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InvitationCategoryService_CreateInvitationCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvitationCategoryServiceServer).CreateInvitationCategory(ctx, req.(*CreateInvitationCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InvitationCategoryService_GetInvitationCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInvitationCategoryByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvitationCategoryServiceServer).GetInvitationCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InvitationCategoryService_GetInvitationCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvitationCategoryServiceServer).GetInvitationCategory(ctx, req.(*GetInvitationCategoryByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InvitationCategoryService_GetAllInvitationCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllInvitationCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvitationCategoryServiceServer).GetAllInvitationCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InvitationCategoryService_GetAllInvitationCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvitationCategoryServiceServer).GetAllInvitationCategory(ctx, req.(*GetAllInvitationCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InvitationCategoryService_UpdateInvitationCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateInvitationCategoryByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvitationCategoryServiceServer).UpdateInvitationCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InvitationCategoryService_UpdateInvitationCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvitationCategoryServiceServer).UpdateInvitationCategory(ctx, req.(*UpdateInvitationCategoryByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InvitationCategoryService_DeleteInvitationCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteInvitationCategoryByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvitationCategoryServiceServer).DeleteInvitationCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InvitationCategoryService_DeleteInvitationCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvitationCategoryServiceServer).DeleteInvitationCategory(ctx, req.(*DeleteInvitationCategoryByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InvitationCategoryService_ServiceDesc is the grpc.ServiceDesc for InvitationCategoryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InvitationCategoryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.InvitationCategoryService",
	HandlerType: (*InvitationCategoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateInvitationCategory",
			Handler:    _InvitationCategoryService_CreateInvitationCategory_Handler,
		},
		{
			MethodName: "GetInvitationCategory",
			Handler:    _InvitationCategoryService_GetInvitationCategory_Handler,
		},
		{
			MethodName: "GetAllInvitationCategory",
			Handler:    _InvitationCategoryService_GetAllInvitationCategory_Handler,
		},
		{
			MethodName: "UpdateInvitationCategory",
			Handler:    _InvitationCategoryService_UpdateInvitationCategory_Handler,
		},
		{
			MethodName: "DeleteInvitationCategory",
			Handler:    _InvitationCategoryService_DeleteInvitationCategory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "invitation-category/invitation_category_service.proto",
}
