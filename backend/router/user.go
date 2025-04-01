package router

import (
	"aquila/api/system"
	"aquila/middleware"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user", middleware.AuthMiddleWare())
	user := system.User{}
	{
		userRouter.POST("/create", user.CreateUserApi)
		userRouter.GET("", user.GetUserApi)
		userRouter.GET("/list", user.GetUserPageApi)
		userRouter.POST("/update", user.UpdateUserApi)
		userRouter.POST("/delete", user.DeleteUserApi)
		userRouter.POST("/bindRole", user.BindRoleApi)
		userRouter.POST("/unbindRole", user.UnbindRoleApi)
		userRouter.POST("/changePassword", user.ChangePasswordApi)
		userRouter.GET("/roles", user.GetUserRolesApi)
		userRouter.GET("/menus", user.GetUserMenuApi)

	}
}
