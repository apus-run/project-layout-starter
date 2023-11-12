// Code generated protoc-gen-go-gin. DO NOT EDIT.
// protoc-gen-go-gin v1.10.1

package v1

import (
	context "context"
	errcode "github.com/apus-run/gaia/pkg/errcode"
	ginx "github.com/apus-run/gaia/pkg/ginx"
	gin "github.com/gin-gonic/gin"
	metadata "google.golang.org/grpc/metadata"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the eagle package it is being compiled against.

// context.
// metadata.
// gin.ginx.errcode.

type UserHTTPServer interface {
	GetUser(context.Context, *GetUserRequest) (*UserReply, error)
	Login(context.Context, *LoginRequest) (*UserReply, error)
	Register(context.Context, *RegisterRequest) (*UserReply, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UserReply, error)
}

func RegisterUserHTTPServer(r gin.IRouter, srv UserHTTPServer) {
	s := &User{
		server: srv,
		router: r,
	}
	s.RegisterService()
}

type User struct {
	server UserHTTPServer
	router gin.IRouter
}

func (s *User) Login_0_HTTP_Handler(ctx *ginx.Context) {
	var in LoginRequest

	if err := ctx.ShouldBindJSON(&in); err != nil {
		e := errcode.ErrInvalidParam.WithDetails(err.Error())
		ctx.Error(e)
		return
	}

	md := metadata.New(nil)
	for k, v := range ctx.Request.Header {
		md.Set(k, v...)
	}
	newCtx := metadata.NewIncomingContext(ctx, md)
	out, err := s.server.(UserHTTPServer).Login(newCtx, &in)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.Success(out)
}

func (s *User) Register_0_HTTP_Handler(ctx *ginx.Context) {
	var in RegisterRequest

	if err := ctx.ShouldBindJSON(&in); err != nil {
		e := errcode.ErrInvalidParam.WithDetails(err.Error())
		ctx.Error(e)
		return
	}

	md := metadata.New(nil)
	for k, v := range ctx.Request.Header {
		md.Set(k, v...)
	}
	newCtx := metadata.NewIncomingContext(ctx, md)
	out, err := s.server.(UserHTTPServer).Register(newCtx, &in)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.Success(out)
}

func (s *User) UpdateUser_0_HTTP_Handler(ctx *ginx.Context) {
	var in UpdateUserRequest

	if err := ctx.ShouldBindJSON(&in); err != nil {
		e := errcode.ErrInvalidParam.WithDetails(err.Error())
		ctx.Error(e)
		return
	}

	md := metadata.New(nil)
	for k, v := range ctx.Request.Header {
		md.Set(k, v...)
	}
	newCtx := metadata.NewIncomingContext(ctx, md)
	out, err := s.server.(UserHTTPServer).UpdateUser(newCtx, &in)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.Success(out)
}

func (s *User) GetUser_0_HTTP_Handler(ctx *ginx.Context) {
	var in GetUserRequest

	if err := ctx.ShouldBindUri(&in); err != nil {
		e := errcode.ErrInvalidParam.WithDetails(err.Error())
		ctx.Error(e)
		return
	}

	md := metadata.New(nil)
	for k, v := range ctx.Request.Header {
		md.Set(k, v...)
	}
	newCtx := metadata.NewIncomingContext(ctx, md)
	out, err := s.server.(UserHTTPServer).GetUser(newCtx, &in)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.Success(out)
}

func (s *User) RegisterService() {
	s.router.Handle("POST", "/api/users/login", ginx.Handle(s.Login_0_HTTP_Handler))
	s.router.Handle("POST", "/api/user", ginx.Handle(s.Register_0_HTTP_Handler))
	s.router.Handle("PUT", "/api/user", ginx.Handle(s.UpdateUser_0_HTTP_Handler))
	s.router.Handle("GET", "/api/user/:id", ginx.Handle(s.GetUser_0_HTTP_Handler))
}
