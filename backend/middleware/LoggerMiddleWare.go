package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func LoggerMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		//请求方式
		method := c.Request.Method
		//请求路由
		reqUrl := c.Request.RequestURI
		//状态码
		statusCode := c.Writer.Status()
		//请求ip
		clientIP := c.ClientIP()
		// 获取当前用户
		userId, isOk := c.Get("id")
		var userIdInt int
		if !isOk {
			userId = 0
		} else {
			userIdInt = int(userId.(float64))
		}
		// 打印日志
		fmt.Println("请求方式:", method, "请求路由:", reqUrl, "状态码:", statusCode, "请求ip:", clientIP, "用户id:", userIdInt)

		c.Next()
	}
}
