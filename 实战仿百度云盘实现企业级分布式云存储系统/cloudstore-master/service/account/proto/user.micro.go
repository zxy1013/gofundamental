// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: user.proto

package go_micro_service_user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for UserService service

type UserService interface {
	// 用户注册
	SignUp(ctx context.Context, in *ReqSignUp, opts ...client.CallOption) (*RespSignUp, error)
	// 用户登录
	SignIn(ctx context.Context, in *ReqSignIn, opts ...client.CallOption) (*RespSignIn, error)
	// 获取用户信息
	UserInfo(ctx context.Context, in *ReqUserInfo, opts ...client.CallOption) (*RespUserInfo, error)
	// 获取用户文件
	UserFiles(ctx context.Context, in *ReqUserFile, opts ...client.CallOption) (*RespUserFile, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.service.user"
	}
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) SignUp(ctx context.Context, in *ReqSignUp, opts ...client.CallOption) (*RespSignUp, error) {
	req := c.c.NewRequest(c.name, "UserService.SignUp", in)
	out := new(RespSignUp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) SignIn(ctx context.Context, in *ReqSignIn, opts ...client.CallOption) (*RespSignIn, error) {
	req := c.c.NewRequest(c.name, "UserService.SignIn", in)
	out := new(RespSignIn)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) UserInfo(ctx context.Context, in *ReqUserInfo, opts ...client.CallOption) (*RespUserInfo, error) {
	req := c.c.NewRequest(c.name, "UserService.UserInfo", in)
	out := new(RespUserInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) UserFiles(ctx context.Context, in *ReqUserFile, opts ...client.CallOption) (*RespUserFile, error) {
	req := c.c.NewRequest(c.name, "UserService.UserFiles", in)
	out := new(RespUserFile)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceHandler interface {
	// 用户注册
	SignUp(context.Context, *ReqSignUp, *RespSignUp) error
	// 用户登录
	SignIn(context.Context, *ReqSignIn, *RespSignIn) error
	// 获取用户信息
	UserInfo(context.Context, *ReqUserInfo, *RespUserInfo) error
	// 获取用户文件
	UserFiles(context.Context, *ReqUserFile, *RespUserFile) error
}

func RegisterUserServiceHandler(s server.Server, hdlr UserServiceHandler, opts ...server.HandlerOption) error {
	type userService interface {
		SignUp(ctx context.Context, in *ReqSignUp, out *RespSignUp) error
		SignIn(ctx context.Context, in *ReqSignIn, out *RespSignIn) error
		UserInfo(ctx context.Context, in *ReqUserInfo, out *RespUserInfo) error
		UserFiles(ctx context.Context, in *ReqUserFile, out *RespUserFile) error
	}
	type UserService struct {
		userService
	}
	h := &userServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&UserService{h}, opts...))
}

type userServiceHandler struct {
	UserServiceHandler
}

func (h *userServiceHandler) SignUp(ctx context.Context, in *ReqSignUp, out *RespSignUp) error {
	return h.UserServiceHandler.SignUp(ctx, in, out)
}

func (h *userServiceHandler) SignIn(ctx context.Context, in *ReqSignIn, out *RespSignIn) error {
	return h.UserServiceHandler.SignIn(ctx, in, out)
}

func (h *userServiceHandler) UserInfo(ctx context.Context, in *ReqUserInfo, out *RespUserInfo) error {
	return h.UserServiceHandler.UserInfo(ctx, in, out)
}

func (h *userServiceHandler) UserFiles(ctx context.Context, in *ReqUserFile, out *RespUserFile) error {
	return h.UserServiceHandler.UserFiles(ctx, in, out)
}
