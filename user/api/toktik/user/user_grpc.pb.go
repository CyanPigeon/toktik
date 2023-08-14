// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: api/toktik/user/user.proto

package user

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
	UserInfo_UserInfoSrv_FullMethodName = "/UserInfo/UserInfoSrv"
)

// UserInfoClient is the client API for UserInfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserInfoClient interface {
	UserInfoSrv(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error)
}

type userInfoClient struct {
	cc grpc.ClientConnInterface
}

func NewUserInfoClient(cc grpc.ClientConnInterface) UserInfoClient {
	return &userInfoClient{cc}
}

func (c *userInfoClient) UserInfoSrv(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error) {
	out := new(UserInfoResponse)
	err := c.cc.Invoke(ctx, UserInfo_UserInfoSrv_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserInfoServer is the server API for UserInfo service.
// All implementations must embed UnimplementedUserInfoServer
// for forward compatibility
type UserInfoServer interface {
	UserInfoSrv(context.Context, *UserInfoRequest) (*UserInfoResponse, error)
	mustEmbedUnimplementedUserInfoServer()
}

// UnimplementedUserInfoServer must be embedded to have forward compatible implementations.
type UnimplementedUserInfoServer struct {
}

func (UnimplementedUserInfoServer) UserInfoSrv(context.Context, *UserInfoRequest) (*UserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserInfoSrv not implemented")
}
func (UnimplementedUserInfoServer) mustEmbedUnimplementedUserInfoServer() {}

// UnsafeUserInfoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserInfoServer will
// result in compilation errors.
type UnsafeUserInfoServer interface {
	mustEmbedUnimplementedUserInfoServer()
}

func RegisterUserInfoServer(s grpc.ServiceRegistrar, srv UserInfoServer) {
	s.RegisterService(&UserInfo_ServiceDesc, srv)
}

func _UserInfo_UserInfoSrv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserInfoServer).UserInfoSrv(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserInfo_UserInfoSrv_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserInfoServer).UserInfoSrv(ctx, req.(*UserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserInfo_ServiceDesc is the grpc.ServiceDesc for UserInfo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserInfo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "UserInfo",
	HandlerType: (*UserInfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserInfoSrv",
			Handler:    _UserInfo_UserInfoSrv_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/toktik/user/user.proto",
}

const (
	UserLogin_UserLoginSrv_FullMethodName = "/UserLogin/UserLoginSrv"
)

// UserLoginClient is the client API for UserLogin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserLoginClient interface {
	UserLoginSrv(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserLoginResponse, error)
}

type userLoginClient struct {
	cc grpc.ClientConnInterface
}

func NewUserLoginClient(cc grpc.ClientConnInterface) UserLoginClient {
	return &userLoginClient{cc}
}

func (c *userLoginClient) UserLoginSrv(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserLoginResponse, error) {
	out := new(UserLoginResponse)
	err := c.cc.Invoke(ctx, UserLogin_UserLoginSrv_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserLoginServer is the server API for UserLogin service.
// All implementations must embed UnimplementedUserLoginServer
// for forward compatibility
type UserLoginServer interface {
	UserLoginSrv(context.Context, *UserLoginRequest) (*UserLoginResponse, error)
	mustEmbedUnimplementedUserLoginServer()
}

// UnimplementedUserLoginServer must be embedded to have forward compatible implementations.
type UnimplementedUserLoginServer struct {
}

func (UnimplementedUserLoginServer) UserLoginSrv(context.Context, *UserLoginRequest) (*UserLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLoginSrv not implemented")
}
func (UnimplementedUserLoginServer) mustEmbedUnimplementedUserLoginServer() {}

// UnsafeUserLoginServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserLoginServer will
// result in compilation errors.
type UnsafeUserLoginServer interface {
	mustEmbedUnimplementedUserLoginServer()
}

func RegisterUserLoginServer(s grpc.ServiceRegistrar, srv UserLoginServer) {
	s.RegisterService(&UserLogin_ServiceDesc, srv)
}

func _UserLogin_UserLoginSrv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserLoginServer).UserLoginSrv(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserLogin_UserLoginSrv_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserLoginServer).UserLoginSrv(ctx, req.(*UserLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserLogin_ServiceDesc is the grpc.ServiceDesc for UserLogin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserLogin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "UserLogin",
	HandlerType: (*UserLoginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserLoginSrv",
			Handler:    _UserLogin_UserLoginSrv_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/toktik/user/user.proto",
}

const (
	UserRegister_UserRegisterSrv_FullMethodName = "/UserRegister/UserRegisterSrv"
)

// UserRegisterClient is the client API for UserRegister service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserRegisterClient interface {
	UserRegisterSrv(ctx context.Context, in *UserRegisterRequest, opts ...grpc.CallOption) (*UserRegisterResponse, error)
}

type userRegisterClient struct {
	cc grpc.ClientConnInterface
}

func NewUserRegisterClient(cc grpc.ClientConnInterface) UserRegisterClient {
	return &userRegisterClient{cc}
}

func (c *userRegisterClient) UserRegisterSrv(ctx context.Context, in *UserRegisterRequest, opts ...grpc.CallOption) (*UserRegisterResponse, error) {
	out := new(UserRegisterResponse)
	err := c.cc.Invoke(ctx, UserRegister_UserRegisterSrv_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserRegisterServer is the server API for UserRegister service.
// All implementations must embed UnimplementedUserRegisterServer
// for forward compatibility
type UserRegisterServer interface {
	UserRegisterSrv(context.Context, *UserRegisterRequest) (*UserRegisterResponse, error)
	mustEmbedUnimplementedUserRegisterServer()
}

// UnimplementedUserRegisterServer must be embedded to have forward compatible implementations.
type UnimplementedUserRegisterServer struct {
}

func (UnimplementedUserRegisterServer) UserRegisterSrv(context.Context, *UserRegisterRequest) (*UserRegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserRegisterSrv not implemented")
}
func (UnimplementedUserRegisterServer) mustEmbedUnimplementedUserRegisterServer() {}

// UnsafeUserRegisterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserRegisterServer will
// result in compilation errors.
type UnsafeUserRegisterServer interface {
	mustEmbedUnimplementedUserRegisterServer()
}

func RegisterUserRegisterServer(s grpc.ServiceRegistrar, srv UserRegisterServer) {
	s.RegisterService(&UserRegister_ServiceDesc, srv)
}

func _UserRegister_UserRegisterSrv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRegisterServer).UserRegisterSrv(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserRegister_UserRegisterSrv_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRegisterServer).UserRegisterSrv(ctx, req.(*UserRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserRegister_ServiceDesc is the grpc.ServiceDesc for UserRegister service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserRegister_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "UserRegister",
	HandlerType: (*UserRegisterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserRegisterSrv",
			Handler:    _UserRegister_UserRegisterSrv_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/toktik/user/user.proto",
}
