package router

import (
	"aquila/api/file"
	"github.com/gin-gonic/gin"
)

func InitFileRouter(Router *gin.RouterGroup) {
	fileRouter := Router.Group("file")
	fileUpload := file.FileUpload{}
	{
		fileRouter.POST("qiniu", fileUpload.UploadFileQiniuApi)
	}
}
