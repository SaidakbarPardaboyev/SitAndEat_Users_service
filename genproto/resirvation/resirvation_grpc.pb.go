// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: resirvation.proto

package resirvation

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

// ResirvationClient is the client API for Resirvation service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ResirvationClient interface {
	CreateRestaurant(ctx context.Context, in *Restuarant, opts ...grpc.CallOption) (*Status, error)
	GetAllRestaurants(ctx context.Context, in *AllRestuarant, opts ...grpc.CallOption) (*Restuanants, error)
	GetRestuarant(ctx context.Context, in *RestuanantId, opts ...grpc.CallOption) (*GetRes, error)
	UpdateRestuarant(ctx context.Context, in *GetRes, opts ...grpc.CallOption) (*Status, error)
	DeleteRestuarant(ctx context.Context, in *RestuanantId, opts ...grpc.CallOption) (*Status, error)
}

type resirvationClient struct {
	cc grpc.ClientConnInterface
}

func NewResirvationClient(cc grpc.ClientConnInterface) ResirvationClient {
	return &resirvationClient{cc}
}

func (c *resirvationClient) CreateRestaurant(ctx context.Context, in *Restuarant, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/resirvation.Resirvation/CreateRestaurant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resirvationClient) GetAllRestaurants(ctx context.Context, in *AllRestuarant, opts ...grpc.CallOption) (*Restuanants, error) {
	out := new(Restuanants)
	err := c.cc.Invoke(ctx, "/resirvation.Resirvation/GetAllRestaurants", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resirvationClient) GetRestuarant(ctx context.Context, in *RestuanantId, opts ...grpc.CallOption) (*GetRes, error) {
	out := new(GetRes)
	err := c.cc.Invoke(ctx, "/resirvation.Resirvation/GetRestuarant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resirvationClient) UpdateRestuarant(ctx context.Context, in *GetRes, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/resirvation.Resirvation/UpdateRestuarant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resirvationClient) DeleteRestuarant(ctx context.Context, in *RestuanantId, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/resirvation.Resirvation/DeleteRestuarant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResirvationServer is the server API for Resirvation service.
// All implementations must embed UnimplementedResirvationServer
// for forward compatibility
type ResirvationServer interface {
	CreateRestaurant(context.Context, *Restuarant) (*Status, error)
	GetAllRestaurants(context.Context, *AllRestuarant) (*Restuanants, error)
	GetRestuarant(context.Context, *RestuanantId) (*GetRes, error)
	UpdateRestuarant(context.Context, *GetRes) (*Status, error)
	DeleteRestuarant(context.Context, *RestuanantId) (*Status, error)
	mustEmbedUnimplementedResirvationServer()
}

// UnimplementedResirvationServer must be embedded to have forward compatible implementations.
type UnimplementedResirvationServer struct {
}

func (UnimplementedResirvationServer) CreateRestaurant(context.Context, *Restuarant) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRestaurant not implemented")
}
func (UnimplementedResirvationServer) GetAllRestaurants(context.Context, *AllRestuarant) (*Restuanants, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllRestaurants not implemented")
}
func (UnimplementedResirvationServer) GetRestuarant(context.Context, *RestuanantId) (*GetRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRestuarant not implemented")
}
func (UnimplementedResirvationServer) UpdateRestuarant(context.Context, *GetRes) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRestuarant not implemented")
}
func (UnimplementedResirvationServer) DeleteRestuarant(context.Context, *RestuanantId) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRestuarant not implemented")
}
func (UnimplementedResirvationServer) mustEmbedUnimplementedResirvationServer() {}

// UnsafeResirvationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResirvationServer will
// result in compilation errors.
type UnsafeResirvationServer interface {
	mustEmbedUnimplementedResirvationServer()
}

func RegisterResirvationServer(s grpc.ServiceRegistrar, srv ResirvationServer) {
	s.RegisterService(&Resirvation_ServiceDesc, srv)
}

func _Resirvation_CreateRestaurant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Restuarant)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResirvationServer).CreateRestaurant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resirvation.Resirvation/CreateRestaurant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResirvationServer).CreateRestaurant(ctx, req.(*Restuarant))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resirvation_GetAllRestaurants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllRestuarant)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResirvationServer).GetAllRestaurants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resirvation.Resirvation/GetAllRestaurants",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResirvationServer).GetAllRestaurants(ctx, req.(*AllRestuarant))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resirvation_GetRestuarant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestuanantId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResirvationServer).GetRestuarant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resirvation.Resirvation/GetRestuarant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResirvationServer).GetRestuarant(ctx, req.(*RestuanantId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resirvation_UpdateRestuarant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResirvationServer).UpdateRestuarant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resirvation.Resirvation/UpdateRestuarant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResirvationServer).UpdateRestuarant(ctx, req.(*GetRes))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resirvation_DeleteRestuarant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestuanantId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResirvationServer).DeleteRestuarant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resirvation.Resirvation/DeleteRestuarant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResirvationServer).DeleteRestuarant(ctx, req.(*RestuanantId))
	}
	return interceptor(ctx, in, info, handler)
}

// Resirvation_ServiceDesc is the grpc.ServiceDesc for Resirvation service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Resirvation_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "resirvation.Resirvation",
	HandlerType: (*ResirvationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRestaurant",
			Handler:    _Resirvation_CreateRestaurant_Handler,
		},
		{
			MethodName: "GetAllRestaurants",
			Handler:    _Resirvation_GetAllRestaurants_Handler,
		},
		{
			MethodName: "GetRestuarant",
			Handler:    _Resirvation_GetRestuarant_Handler,
		},
		{
			MethodName: "UpdateRestuarant",
			Handler:    _Resirvation_UpdateRestuarant_Handler,
		},
		{
			MethodName: "DeleteRestuarant",
			Handler:    _Resirvation_DeleteRestuarant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "resirvation.proto",
}
