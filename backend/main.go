package main

import (
	"aquila/global"
	"aquila/initialize"
	"aquila/scheduler"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const AppMode = "debug"

// const AppMode = "release"

func main() {
	gin.SetMode(AppMode)
	// 初始化配置
	global.AquilaViper = initialize.InitViper()
	// 初始化zap日志库
	global.AquilaLog = initialize.InitZap()
	zap.ReplaceGlobals(global.AquilaLog)
	global.AquilaLog.Info("server run success on ", zap.String("zap_log", "zap_log"))
	// 初始化数据库
	global.AquilaDb = initialize.InitGorm()
	// 初始化minio
	initialize.InitMinio()
	// 初始化redis
	initialize.InitRedis()
	// 启动定时任务
	scheduler.StartScheduler()
	// 初始化路由
	initialize.InitServer()
}
