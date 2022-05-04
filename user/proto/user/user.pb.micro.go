// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/user/user.proto

package user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for UserService service

func NewUserServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for UserService service

type UserService interface {
	Pagination(ctx context.Context, in *PaginationRequest, opts ...client.CallOption) (*PaginationResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*UserResponse, error)
	Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*UserResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UserResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*UserResponse, error)
	Auth(ctx context.Context, in *AuthRequest, opts ...client.CallOption) (*TokenResponse, error)
	ValidateToken(ctx context.Context, in *TokenRequest, opts ...client.CallOption) (*TokenResponse, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) Pagination(ctx context.Context, in *PaginationRequest, opts ...client.CallOption) (*PaginationResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.Pagination", in)
	out := new(PaginationResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Get(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*UserResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.Get", in)
	out := new(UserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*UserResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.Create", in)
	out := new(UserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UserResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.Update", in)
	out := new(UserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*UserResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.Delete", in)
	out := new(UserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Auth(ctx context.Context, in *AuthRequest, opts ...client.CallOption) (*TokenResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.Auth", in)
	out := new(TokenResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) ValidateToken(ctx context.Context, in *TokenRequest, opts ...client.CallOption) (*TokenResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.ValidateToken", in)
	out := new(TokenResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceHandler interface {
	Pagination(context.Context, *PaginationRequest, *PaginationResponse) error
	Get(context.Context, *GetRequest, *UserResponse) error
	Create(context.Context, *CreateRequest, *UserResponse) error
	Update(context.Context, *UpdateRequest, *UserResponse) error
	Delete(context.Context, *DeleteRequest, *UserResponse) error
	Auth(context.Context, *AuthRequest, *TokenResponse) error
	ValidateToken(context.Context, *TokenRequest, *TokenResponse) error
}

func RegisterUserServiceHandler(s server.Server, hdlr UserServiceHandler, opts ...server.HandlerOption) error {
	type userService interface {
		Pagination(ctx context.Context, in *PaginationRequest, out *PaginationResponse) error
		Get(ctx context.Context, in *GetRequest, out *UserResponse) error
		Create(ctx context.Context, in *CreateRequest, out *UserResponse) error
		Update(ctx context.Context, in *UpdateRequest, out *UserResponse) error
		Delete(ctx context.Context, in *DeleteRequest, out *UserResponse) error
		Auth(ctx context.Context, in *AuthRequest, out *TokenResponse) error
		ValidateToken(ctx context.Context, in *TokenRequest, out *TokenResponse) error
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

func (h *userServiceHandler) Pagination(ctx context.Context, in *PaginationRequest, out *PaginationResponse) error {
	return h.UserServiceHandler.Pagination(ctx, in, out)
}

func (h *userServiceHandler) Get(ctx context.Context, in *GetRequest, out *UserResponse) error {
	return h.UserServiceHandler.Get(ctx, in, out)
}

func (h *userServiceHandler) Create(ctx context.Context, in *CreateRequest, out *UserResponse) error {
	return h.UserServiceHandler.Create(ctx, in, out)
}

func (h *userServiceHandler) Update(ctx context.Context, in *UpdateRequest, out *UserResponse) error {
	return h.UserServiceHandler.Update(ctx, in, out)
}

func (h *userServiceHandler) Delete(ctx context.Context, in *DeleteRequest, out *UserResponse) error {
	return h.UserServiceHandler.Delete(ctx, in, out)
}

func (h *userServiceHandler) Auth(ctx context.Context, in *AuthRequest, out *TokenResponse) error {
	return h.UserServiceHandler.Auth(ctx, in, out)
}

func (h *userServiceHandler) ValidateToken(ctx context.Context, in *TokenRequest, out *TokenResponse) error {
	return h.UserServiceHandler.ValidateToken(ctx, in, out)
}
