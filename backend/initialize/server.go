package initialize

import (
	"aquila/global"
	"aquila/middleware"
	"aquila/router"
	"fmt"

	"github.com/gin-gonic/gin"
)

func InitServer() {
	// 初始化路由
	routers := Routers()
	address := fmt.Sprintf(":%d", global.AquilaConfig.App.Port)
	// 启动服务
	if err := routers.Run(address); err != nil {
		panic(err)
	}
}

// Routers 配置全局的路由
func Routers() *gin.Engine {
	Router := gin.Default()
	// 注册中间件
	Router.Use(
		middleware.CorsMiddleWare(),
		middleware.LoggerMiddleWare(),
		middleware.RecoverMiddleWare(),
	)
	// 配置全局路径
	ApiGroup := Router.Group("/api")
	// 注册路由
	router.InitCommonRouter(ApiGroup) // 公共路由
	router.InitUserRouter(ApiGroup)   // 用户路由
	router.InitRoleRouter(ApiGroup)   // 角色路由
	router.InitMenuRouter(ApiGroup)   // 菜单路由
	router.InitFileRouter(ApiGroup)   // 文件路由
	router.InitInitDbRouter(ApiGroup) // 初始化路由
	router.InitManagerRouter(ApiGroup) // 业务路由

	return Router
}
