package router

import (
	"aquila/api/system"
	"aquila/middleware"

	"github.com/gin-gonic/gin"
)

func InitRoleRouter(Router *gin.RouterGroup) {
	roleRouter := Router.Group("role", middleware.AuthMiddleWare())
	role := system.Role{}
	{
		roleRouter.POST("/create", role.CreateRoleApi)
		roleRouter.GET("", role.GetRoleApi)
		roleRouter.GET("/list", role.GetRolePageApi)
		roleRouter.POST("/update", role.UpdateRoleApi)
		roleRouter.POST("/delete", role.DeleteRoleApi)
		roleRouter.POST("/bindMenu", role.BindMenuApi)
		roleRouter.POST("/unbindMenu", role.UnbindMenuApi)
		roleRouter.GET("/menus", role.GetRoleMenusApi)
		roleRouter.GET("/users", role.GetRoleUsersApi)
	}
}
