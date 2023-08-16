// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.6.3
// - protoc             v4.23.4
// source: api/toktik/user/user.proto

package user

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

const OperationUserInfoUserInfoSrv = "/UserInfo/UserInfoService"

type UserInfoHTTPServer interface {
	UserInfoSrv(context.Context, *UserInfoRequest) (*UserInfoResponse, error)
}

func RegisterUserInfoHTTPServer(s *http.Server, srv UserInfoHTTPServer) {
	r := s.Route("/")
	r.GET("/douyin/user", _UserInfo_UserInfoSrv0_HTTP_Handler(srv))
}

func _UserInfo_UserInfoSrv0_HTTP_Handler(srv UserInfoHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserInfoRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserInfoUserInfoSrv)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UserInfoSrv(ctx, req.(*UserInfoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserInfoResponse)
		return ctx.Result(200, reply)
	}
}

type UserInfoHTTPClient interface {
	UserInfoSrv(ctx context.Context, req *UserInfoRequest, opts ...http.CallOption) (rsp *UserInfoResponse, err error)
}

type UserInfoHTTPClientImpl struct {
	cc *http.Client
}

func NewUserInfoHTTPClient(client *http.Client) UserInfoHTTPClient {
	return &UserInfoHTTPClientImpl{client}
}

func (c *UserInfoHTTPClientImpl) UserInfoSrv(ctx context.Context, in *UserInfoRequest, opts ...http.CallOption) (*UserInfoResponse, error) {
	var out UserInfoResponse
	pattern := "/douyin/user"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserInfoUserInfoSrv))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

const OperationUserLoginUserLoginSrv = "/UserLogin/UserLoginService"

type UserLoginHTTPServer interface {
	UserLoginSrv(context.Context, *UserLoginRequest) (*UserLoginResponse, error)
}

func RegisterUserLoginHTTPServer(s *http.Server, srv UserLoginHTTPServer) {
	r := s.Route("/")
	r.POST("/douyin/user/login", _UserLogin_UserLoginSrv0_HTTP_Handler(srv))
}

func _UserLogin_UserLoginSrv0_HTTP_Handler(srv UserLoginHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserLoginRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserLoginUserLoginSrv)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UserLoginSrv(ctx, req.(*UserLoginRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserLoginResponse)
		return ctx.Result(200, reply)
	}
}

type UserLoginHTTPClient interface {
	UserLoginSrv(ctx context.Context, req *UserLoginRequest, opts ...http.CallOption) (rsp *UserLoginResponse, err error)
}

type UserLoginHTTPClientImpl struct {
	cc *http.Client
}

func NewUserLoginHTTPClient(client *http.Client) UserLoginHTTPClient {
	return &UserLoginHTTPClientImpl{client}
}

func (c *UserLoginHTTPClientImpl) UserLoginSrv(ctx context.Context, in *UserLoginRequest, opts ...http.CallOption) (*UserLoginResponse, error) {
	var out UserLoginResponse
	pattern := "/douyin/user/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserLoginUserLoginSrv))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

const OperationUserRegisterUserRegisterSrv = "/UserRegister/UserRegisterService"

type UserRegisterHTTPServer interface {
	UserRegisterSrv(context.Context, *UserRegisterRequest) (*UserRegisterResponse, error)
}

func RegisterUserRegisterHTTPServer(s *http.Server, srv UserRegisterHTTPServer) {
	r := s.Route("/")
	r.POST("/douyin/user/register", _UserRegister_UserRegisterSrv0_HTTP_Handler(srv))
}

func _UserRegister_UserRegisterSrv0_HTTP_Handler(srv UserRegisterHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserRegisterRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserRegisterUserRegisterSrv)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UserRegisterSrv(ctx, req.(*UserRegisterRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserRegisterResponse)
		return ctx.Result(200, reply)
	}
}

type UserRegisterHTTPClient interface {
	UserRegisterSrv(ctx context.Context, req *UserRegisterRequest, opts ...http.CallOption) (rsp *UserRegisterResponse, err error)
}

type UserRegisterHTTPClientImpl struct {
	cc *http.Client
}

func NewUserRegisterHTTPClient(client *http.Client) UserRegisterHTTPClient {
	return &UserRegisterHTTPClientImpl{client}
}

func (c *UserRegisterHTTPClientImpl) UserRegisterSrv(ctx context.Context, in *UserRegisterRequest, opts ...http.CallOption) (*UserRegisterResponse, error) {
	var out UserRegisterResponse
	pattern := "/douyin/user/register"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserRegisterUserRegisterSrv))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
