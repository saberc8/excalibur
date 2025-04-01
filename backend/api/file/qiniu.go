package file

import (
	"aquila/global"
	"aquila/utils"
	"context"
	"fmt"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

type FileUpload struct{}

func (f FileUpload) UploadFileQiniuApi(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	fmt.Printf("Received file: %+v\n", file)
	if err != nil {
		utils.Fail(ctx, "获取文件失败"+err.Error())
		return
	}

	// 检查配置是否存在
	if global.AquilaConfig.Qiniu.AccessKey == "" || global.AquilaConfig.Qiniu.SecretKey == "" {
		utils.Fail(ctx, "七牛云配置错误")
		return
	}

	// 生成上传Token，增加过期时间
	putPolicy := storage.PutPolicy{
		Scope:      global.AquilaConfig.Qiniu.Bucket,
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"name":"$(fname)","mimeType":"$(mimeType)"}`,
		Expires:    7200,
	}

	// 打印配置信息用于调试
	fmt.Printf("Qiniu Config - AccessKey: %s, Bucket: %s\n",
		global.AquilaConfig.Qiniu.AccessKey,
		global.AquilaConfig.Qiniu.Bucket)

	mac := qbox.NewMac(global.AquilaConfig.Qiniu.AccessKey, global.AquilaConfig.Qiniu.SecretKey)
	upToken := putPolicy.UploadToken(mac)

	// 打印生成的token用于调试
	fmt.Printf("Generated Upload Token: %s\n", upToken)

	// 创建自定义区域
	customZone := &storage.Zone{
		SrcUpHosts: []string{
			"up-cn-east-2.qiniup.com",
		},
		CdnUpHosts: []string{
			"up-cn-east-2.qiniup.com",
		},
	}

	cfg := storage.Config{
		Zone:          customZone,
		UseCdnDomains: false,
		UseHTTPS:      true,
	}
	formUploader := storage.NewFormUploader(&cfg)

	fileHandle, err := file.Open()
	if err != nil {
		utils.Fail(ctx, "打开文件失败"+err.Error())
		return
	}
	defer fileHandle.Close()

	fileExt := path.Ext(file.Filename)
	key := fmt.Sprintf("qdd/%s%s", uuid.New().String(), fileExt)

	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": file.Filename,
		},
	}

	err = formUploader.Put(context.Background(), &ret, upToken, key, fileHandle, file.Size, &putExtra)
	if err != nil {
		fmt.Printf("上传错误详情: %v\n", err)
		errMsg := "上传文件失败: " + err.Error()
		utils.Fail(ctx, errMsg)
		return
	}

	if ret.Key == "" {
		fmt.Println("未获取到文件Key")
		utils.Fail(ctx, "上传失败，未获取到文件地址")
		return
	}

	fileUrl := fmt.Sprintf("https://qiniu.linkxspace.cn/%s", ret.Key)
	utils.Success(ctx, gin.H{
		"url": fileUrl,
	})
}
