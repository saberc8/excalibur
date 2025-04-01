package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response(ctx *gin.Context, code int, message string, data interface{}) {
	ctx.Header("Content-Type", "application/json; charset=utf-8")
	ctx.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": message,
		"result":  data,
	})
}

// Success 成功的请求
func Success(ctx *gin.Context, data interface{}) {
	ctx.Header("Content-Type", "application/json; charset=utf-8")
	Response(ctx, 0, "成功", data)
}

// Fail 失败的请求
func Fail(ctx *gin.Context, message string) {
	Response(ctx, 1, message, nil)
}

// 需要重新登录的请求
func FailExpired(ctx *gin.Context, message string) {
	Response(ctx, 4, message, nil)
}

type PageVo struct {
	Data     interface{} `json:"data"`     // 数据
	Total    int64       `json:"total"`    // 总条数
	PageSize int64       `json:"pageSize"` // 当前条数
	PageNum  int64       `json:"pageNum"`  // 当前页数
}
