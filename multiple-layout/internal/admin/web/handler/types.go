package handler

import (
	"github.com/gin-gonic/gin"
)

// Router 加载路由，使用侧提供接口，实现侧需要实现该接口
type Router interface {
	RegisterRoutes(engine *gin.Engine)
}

// Result defines HTTP JSON response
type Result struct {
	Code    int      `json:"code"`
	Msg     string   `json:"msg"`
	Data    any      `json:"data,omitempty"`
	Details []string `json:"details,omitempty"`
}
