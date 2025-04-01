package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"aquila/utils"
)

func RecoverMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				utils.Fail(c, fmt.Sprint(err))
				c.Abort()
				return
			}
		}()
	}
}
