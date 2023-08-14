// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.1
// source: api/toktik/message/message.proto

package message

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
	MessageAction_MessageActionSrv_FullMethodName = "/MessageAction/MessageActionSrv"
)

// MessageActionClient is the client API for MessageAction service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessageActionClient interface {
	MessageActionSrv(ctx context.Context, in *MessageActionRequest, opts ...grpc.CallOption) (*MessageActionResponse, error)
}

type messageActionClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageActionClient(cc grpc.ClientConnInterface) MessageActionClient {
	return &messageActionClient{cc}
}

func (c *messageActionClient) MessageActionSrv(ctx context.Context, in *MessageActionRequest, opts ...grpc.CallOption) (*MessageActionResponse, error) {
	out := new(MessageActionResponse)
	err := c.cc.Invoke(ctx, MessageAction_MessageActionSrv_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageActionServer is the server API for MessageAction service.
// All implementations must embed UnimplementedMessageActionServer
// for forward compatibility
type MessageActionServer interface {
	MessageActionSrv(context.Context, *MessageActionRequest) (*MessageActionResponse, error)
	mustEmbedUnimplementedMessageActionServer()
}

// UnimplementedMessageActionServer must be embedded to have forward compatible implementations.
type UnimplementedMessageActionServer struct {
}

func (UnimplementedMessageActionServer) MessageActionSrv(context.Context, *MessageActionRequest) (*MessageActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MessageActionSrv not implemented")
}
func (UnimplementedMessageActionServer) mustEmbedUnimplementedMessageActionServer() {}

// UnsafeMessageActionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessageActionServer will
// result in compilation errors.
type UnsafeMessageActionServer interface {
	mustEmbedUnimplementedMessageActionServer()
}

func RegisterMessageActionServer(s grpc.ServiceRegistrar, srv MessageActionServer) {
	s.RegisterService(&MessageAction_ServiceDesc, srv)
}

func _MessageAction_MessageActionSrv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageActionServer).MessageActionSrv(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageAction_MessageActionSrv_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageActionServer).MessageActionSrv(ctx, req.(*MessageActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MessageAction_ServiceDesc is the grpc.ServiceDesc for MessageAction service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MessageAction_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "MessageAction",
	HandlerType: (*MessageActionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MessageActionSrv",
			Handler:    _MessageAction_MessageActionSrv_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/toktik/message/message.proto",
}

const (
	MessageHistory_MessageHistorySrv_FullMethodName = "/MessageHistory/MessageHistorySrv"
)

// MessageHistoryClient is the client API for MessageHistory service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessageHistoryClient interface {
	MessageHistorySrv(ctx context.Context, in *MessageHistoryRequest, opts ...grpc.CallOption) (*MessageHistoryResponse, error)
}

type messageHistoryClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageHistoryClient(cc grpc.ClientConnInterface) MessageHistoryClient {
	return &messageHistoryClient{cc}
}

func (c *messageHistoryClient) MessageHistorySrv(ctx context.Context, in *MessageHistoryRequest, opts ...grpc.CallOption) (*MessageHistoryResponse, error) {
	out := new(MessageHistoryResponse)
	err := c.cc.Invoke(ctx, MessageHistory_MessageHistorySrv_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageHistoryServer is the server API for MessageHistory service.
// All implementations must embed UnimplementedMessageHistoryServer
// for forward compatibility
type MessageHistoryServer interface {
	MessageHistorySrv(context.Context, *MessageHistoryRequest) (*MessageHistoryResponse, error)
	mustEmbedUnimplementedMessageHistoryServer()
}

// UnimplementedMessageHistoryServer must be embedded to have forward compatible implementations.
type UnimplementedMessageHistoryServer struct {
}

func (UnimplementedMessageHistoryServer) MessageHistorySrv(context.Context, *MessageHistoryRequest) (*MessageHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MessageHistorySrv not implemented")
}
func (UnimplementedMessageHistoryServer) mustEmbedUnimplementedMessageHistoryServer() {}

// UnsafeMessageHistoryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessageHistoryServer will
// result in compilation errors.
type UnsafeMessageHistoryServer interface {
	mustEmbedUnimplementedMessageHistoryServer()
}

func RegisterMessageHistoryServer(s grpc.ServiceRegistrar, srv MessageHistoryServer) {
	s.RegisterService(&MessageHistory_ServiceDesc, srv)
}

func _MessageHistory_MessageHistorySrv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageHistoryServer).MessageHistorySrv(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageHistory_MessageHistorySrv_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageHistoryServer).MessageHistorySrv(ctx, req.(*MessageHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MessageHistory_ServiceDesc is the grpc.ServiceDesc for MessageHistory service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MessageHistory_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "MessageHistory",
	HandlerType: (*MessageHistoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MessageHistorySrv",
			Handler:    _MessageHistory_MessageHistorySrv_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/toktik/message/message.proto",
}
