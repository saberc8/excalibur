package router

import (
	"aquila/api/system"
	"aquila/middleware"

	"github.com/gin-gonic/gin"
)

func InitMenuRouter(Router *gin.RouterGroup) {
	menuRouter := Router.Group("menu", middleware.AuthMiddleWare())
	menu := system.Menu{}
	{
		menuRouter.POST("/create", menu.CreateMenuApi) // 创建菜单
		menuRouter.GET("", menu.GetMenuApi)            // 获取单个菜单
		menuRouter.GET("/list", menu.GetMenuPageApi)   // 获取菜单列表
		menuRouter.POST("/update", menu.UpdateMenuApi) // 更新菜单
		menuRouter.POST("/delete", menu.DeleteMenuApi) // 删除菜单
		menuRouter.GET("/roles", menu.GetMenuRolesApi) // 获取拥有该菜单的角色列表
		menuRouter.GET("/tree", menu.GetMenuTreeApi)   // 获取菜单树形结构
	}
}
