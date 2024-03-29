// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: proto/test.proto

package test

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

// TestServiceClient is the client API for TestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TestServiceClient interface {
	AddUser(ctx context.Context, in *AddUserReq, opts ...grpc.CallOption) (*AddUserResp, error)
	ModUser(ctx context.Context, in *ModUserReq, opts ...grpc.CallOption) (*ModUserResp, error)
	DelUser(ctx context.Context, in *DelUserReq, opts ...grpc.CallOption) (*DelUserResp, error)
	DelUsersByOr(ctx context.Context, in *DelUsersByOrReq, opts ...grpc.CallOption) (*DelUsersByOrResp, error)
	ListUsers(ctx context.Context, in *ListUsersReq, opts ...grpc.CallOption) (*ListUsersResp, error)
	// search条件の他にpage情報も使う
	ListUsersByOr(ctx context.Context, in *ListUsersByOrReq, opts ...grpc.CallOption) (*ListUsersByOrResp, error)
}

type testServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTestServiceClient(cc grpc.ClientConnInterface) TestServiceClient {
	return &testServiceClient{cc}
}

func (c *testServiceClient) AddUser(ctx context.Context, in *AddUserReq, opts ...grpc.CallOption) (*AddUserResp, error) {
	out := new(AddUserResp)
	err := c.cc.Invoke(ctx, "/test.TestService/AddUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) ModUser(ctx context.Context, in *ModUserReq, opts ...grpc.CallOption) (*ModUserResp, error) {
	out := new(ModUserResp)
	err := c.cc.Invoke(ctx, "/test.TestService/ModUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) DelUser(ctx context.Context, in *DelUserReq, opts ...grpc.CallOption) (*DelUserResp, error) {
	out := new(DelUserResp)
	err := c.cc.Invoke(ctx, "/test.TestService/DelUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) DelUsersByOr(ctx context.Context, in *DelUsersByOrReq, opts ...grpc.CallOption) (*DelUsersByOrResp, error) {
	out := new(DelUsersByOrResp)
	err := c.cc.Invoke(ctx, "/test.TestService/DelUsersByOr", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) ListUsers(ctx context.Context, in *ListUsersReq, opts ...grpc.CallOption) (*ListUsersResp, error) {
	out := new(ListUsersResp)
	err := c.cc.Invoke(ctx, "/test.TestService/ListUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) ListUsersByOr(ctx context.Context, in *ListUsersByOrReq, opts ...grpc.CallOption) (*ListUsersByOrResp, error) {
	out := new(ListUsersByOrResp)
	err := c.cc.Invoke(ctx, "/test.TestService/ListUsersByOr", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestServiceServer is the server API for TestService service.
// All implementations must embed UnimplementedTestServiceServer
// for forward compatibility
type TestServiceServer interface {
	AddUser(context.Context, *AddUserReq) (*AddUserResp, error)
	ModUser(context.Context, *ModUserReq) (*ModUserResp, error)
	DelUser(context.Context, *DelUserReq) (*DelUserResp, error)
	DelUsersByOr(context.Context, *DelUsersByOrReq) (*DelUsersByOrResp, error)
	ListUsers(context.Context, *ListUsersReq) (*ListUsersResp, error)
	// search条件の他にpage情報も使う
	ListUsersByOr(context.Context, *ListUsersByOrReq) (*ListUsersByOrResp, error)
	mustEmbedUnimplementedTestServiceServer()
}

// UnimplementedTestServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTestServiceServer struct {
}

func (UnimplementedTestServiceServer) AddUser(context.Context, *AddUserReq) (*AddUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUser not implemented")
}
func (UnimplementedTestServiceServer) ModUser(context.Context, *ModUserReq) (*ModUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModUser not implemented")
}
func (UnimplementedTestServiceServer) DelUser(context.Context, *DelUserReq) (*DelUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelUser not implemented")
}
func (UnimplementedTestServiceServer) DelUsersByOr(context.Context, *DelUsersByOrReq) (*DelUsersByOrResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelUsersByOr not implemented")
}
func (UnimplementedTestServiceServer) ListUsers(context.Context, *ListUsersReq) (*ListUsersResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsers not implemented")
}
func (UnimplementedTestServiceServer) ListUsersByOr(context.Context, *ListUsersByOrReq) (*ListUsersByOrResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsersByOr not implemented")
}
func (UnimplementedTestServiceServer) mustEmbedUnimplementedTestServiceServer() {}

// UnsafeTestServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TestServiceServer will
// result in compilation errors.
type UnsafeTestServiceServer interface {
	mustEmbedUnimplementedTestServiceServer()
}

func RegisterTestServiceServer(s grpc.ServiceRegistrar, srv TestServiceServer) {
	s.RegisterService(&TestService_ServiceDesc, srv)
}

func _TestService_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.TestService/AddUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).AddUser(ctx, req.(*AddUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_ModUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).ModUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.TestService/ModUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).ModUser(ctx, req.(*ModUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_DelUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).DelUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.TestService/DelUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).DelUser(ctx, req.(*DelUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_DelUsersByOr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelUsersByOrReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).DelUsersByOr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.TestService/DelUsersByOr",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).DelUsersByOr(ctx, req.(*DelUsersByOrReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_ListUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUsersReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).ListUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.TestService/ListUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).ListUsers(ctx, req.(*ListUsersReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_ListUsersByOr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUsersByOrReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).ListUsersByOr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.TestService/ListUsersByOr",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).ListUsersByOr(ctx, req.(*ListUsersByOrReq))
	}
	return interceptor(ctx, in, info, handler)
}

// TestService_ServiceDesc is the grpc.ServiceDesc for TestService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TestService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "test.TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddUser",
			Handler:    _TestService_AddUser_Handler,
		},
		{
			MethodName: "ModUser",
			Handler:    _TestService_ModUser_Handler,
		},
		{
			MethodName: "DelUser",
			Handler:    _TestService_DelUser_Handler,
		},
		{
			MethodName: "DelUsersByOr",
			Handler:    _TestService_DelUsersByOr_Handler,
		},
		{
			MethodName: "ListUsers",
			Handler:    _TestService_ListUsers_Handler,
		},
		{
			MethodName: "ListUsersByOr",
			Handler:    _TestService_ListUsersByOr_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/test.proto",
}
