package router

import (
	"aquila/api/manager"
	"aquila/middleware"

	"github.com/gin-gonic/gin"
)

func InitManagerRouter(Router *gin.RouterGroup) {
	managerRouter := Router.Group("")
	InitSpaceRouter(managerRouter)
	InitProjectRouter(managerRouter)
	InitTaskRouter(managerRouter)
}

func InitSpaceRouter(Router *gin.RouterGroup) {
	spaceRouter := Router.Group("space", middleware.AuthMiddleWare())
	space := manager.Space{}
	{
		spaceRouter.GET("/list", space.GetSpaceListApi)
		spaceRouter.GET("", space.GetSpaceApi)
		spaceRouter.POST("/create", space.CreateSpaceApi)
		spaceRouter.POST("/update", space.UpdateSpaceApi)
		spaceRouter.POST("/delete", space.DeleteSpaceApi)
	}
}

func InitProjectRouter(Router *gin.RouterGroup) {
	projectRouter := Router.Group("project", middleware.AuthMiddleWare())
	project := manager.Project{}
	{
		projectRouter.GET("/list", project.GetProjectListApi)
		projectRouter.GET("", project.GetProjectApi)
		projectRouter.POST("/create", project.CreateProjectApi)
		projectRouter.POST("/update", project.UpdateProjectApi)
		projectRouter.POST("/delete", project.DeleteProjectApi)
	}
}

func InitTaskRouter(Router *gin.RouterGroup) {
	taskRouter := Router.Group("task", middleware.AuthMiddleWare())
	task := manager.Task{}
	{
		taskRouter.GET("/list", task.GetTaskListApi)
		taskRouter.GET("", task.GetTaskApi)
		taskRouter.POST("/create", task.CreateTaskApi)
		taskRouter.POST("/update", task.UpdateTaskApi)
		taskRouter.POST("/delete", task.DeleteTaskApi)
	}
}
