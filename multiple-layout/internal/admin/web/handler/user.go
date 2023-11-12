package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"basic-layout/internal/core/svc"
)

var _ Router = &userHandler{}

type UserHandler interface {
	Login(ctx *gin.Context)
	GetUser(ctx *gin.Context)
	RegisterRoutes(server *gin.Engine)
}

type userHandler struct {
	svc svc.UserService
}

func NewUserHandler(svc svc.UserService) UserHandler {
	return &userHandler{
		svc: svc,
	}
}

func (h *userHandler) RegisterRoutes(server *gin.Engine) {
	// 分组注册
	ug := server.Group("/users")
	ug.GET("/user/:id", h.GetUser)
	ug.POST("/login", h.Login)
}

func (h *userHandler) Login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Result{
		Code: 0,
		Msg:  "登录成功",
	})
}

func (h *userHandler) GetUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Result{
		Code: 0,
		Msg:  "查询用户成功",
		Data: gin.H{
			"id":   1,
			"name": "moocss",
		},
	})
}
