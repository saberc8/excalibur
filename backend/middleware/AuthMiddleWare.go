package middleware

import (
	"aquila/utils"
	"errors"

	"github.com/gin-gonic/gin"
)

// AuthMiddleWare 中间件校验token登录
func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			utils.FailExpired(c, "请求未携带token，无权限访问")
			c.Abort()
			return
		}
		j := utils.NewJWT()
		if len(token) > 7 && token[:7] == "Bearer " {
			token = token[7:]
		}
		claims, err := j.ParseToken(token)
		if err != nil {
			if errors.Is(utils.TokenExpired, err) {
				utils.FailExpired(c, "授权已过期")
				c.Abort()
				return
			}
			utils.Fail(c, err.Error())
			c.Abort()
			return
		}

		// 只在token即将过期时才刷新
		if j.NeedRefresh(claims) {
			newToken, err := j.RefreshToken(token)
			if err == nil {
				c.Header("New-Token", newToken)
			}
		}

		// 继续交由下一个路由处理,并将解析出的信息传递下去
		c.Set("uid", claims.UID)
		c.Next()
	}
}
