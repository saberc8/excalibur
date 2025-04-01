package router

import (
	"aquila/api/common"
	"aquila/api/minio" // 新增导入

	"github.com/gin-gonic/gin"
)

func InitCommonRouter(Router *gin.RouterGroup) {
	resisterRouter := Router.Group("common")
	auth := common.Auth{}
	m := minio.Minio{} // 新增实例化
	resisterRouter.GET("/captcha", auth.Captcha)
	resisterRouter.POST("/login", auth.Login)
	resisterRouter.POST("/logout", auth.Logout)
	resisterRouter.POST("/register", auth.Register)
	resisterRouter.GET("/buckets", m.GetBuckets) // 新增路由
}
