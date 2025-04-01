package router

import (
	"aquila/api/system"
	"github.com/gin-gonic/gin"
)

func InitInitDbRouter(Router *gin.RouterGroup) {
	initRouter := Router.Group("init")
	db := system.InitDb{}
	{
		initRouter.POST("db", db.InitializeDBApi) // 初始化数据库
	}
}
