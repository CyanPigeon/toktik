// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.6.3
// - protoc             v3.19.1
// source: api/toktik/message/message.proto

package message

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationMessageActionMessageActionSrv = "/MessageAction/MessageActionSrv"

type MessageActionHTTPServer interface {
	MessageActionSrv(context.Context, *MessageActionRequest) (*MessageActionResponse, error)
}

func RegisterMessageActionHTTPServer(s *http.Server, srv MessageActionHTTPServer) {
	r := s.Route("/")
	r.POST("/douyin/message/action", _MessageAction_MessageActionSrv0_HTTP_Handler(srv))
}

func _MessageAction_MessageActionSrv0_HTTP_Handler(srv MessageActionHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in MessageActionRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMessageActionMessageActionSrv)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.MessageActionSrv(ctx, req.(*MessageActionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*MessageActionResponse)
		return ctx.JSON(200, reply)
	}
}

type MessageActionHTTPClient interface {
	MessageActionSrv(ctx context.Context, req *MessageActionRequest, opts ...http.CallOption) (rsp *MessageActionResponse, err error)
}

type MessageActionHTTPClientImpl struct {
	cc *http.Client
}

func NewMessageActionHTTPClient(client *http.Client) MessageActionHTTPClient {
	return &MessageActionHTTPClientImpl{client}
}

func (c *MessageActionHTTPClientImpl) MessageActionSrv(ctx context.Context, in *MessageActionRequest, opts ...http.CallOption) (*MessageActionResponse, error) {
	var out MessageActionResponse
	pattern := "/douyin/message/action"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationMessageActionMessageActionSrv))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

const OperationMessageHistoryMessageHistorySrv = "/MessageHistory/MessageHistorySrv"

type MessageHistoryHTTPServer interface {
	MessageHistorySrv(context.Context, *MessageHistoryRequest) (*MessageHistoryResponse, error)
}

func RegisterMessageHistoryHTTPServer(s *http.Server, srv MessageHistoryHTTPServer) {
	r := s.Route("/")
	r.GET("/douyin/message/chat", _MessageHistory_MessageHistorySrv0_HTTP_Handler(srv))
}

func _MessageHistory_MessageHistorySrv0_HTTP_Handler(srv MessageHistoryHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in MessageHistoryRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMessageHistoryMessageHistorySrv)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.MessageHistorySrv(ctx, req.(*MessageHistoryRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*MessageHistoryResponse)
		return ctx.JSON(200, reply)
	}
}

type MessageHistoryHTTPClient interface {
	MessageHistorySrv(ctx context.Context, req *MessageHistoryRequest, opts ...http.CallOption) (rsp *MessageHistoryResponse, err error)
}

type MessageHistoryHTTPClientImpl struct {
	cc *http.Client
}

func NewMessageHistoryHTTPClient(client *http.Client) MessageHistoryHTTPClient {
	return &MessageHistoryHTTPClientImpl{client}
}

func (c *MessageHistoryHTTPClientImpl) MessageHistorySrv(ctx context.Context, in *MessageHistoryRequest, opts ...http.CallOption) (*MessageHistoryResponse, error) {
	var out MessageHistoryResponse
	pattern := "/douyin/message/chat"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationMessageHistoryMessageHistorySrv))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
