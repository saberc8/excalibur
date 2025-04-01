package initialize

import (
	"aquila/global"
	"flag"
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"aquila/constants"
)

// InitViper 优先级: 命令行 > 环境变量 > 默认值
func InitViper(path ...string) *viper.Viper {
	var config string

	if len(path) == 0 {
		// 定义命令行flag参数，格式：flag.TypeVar(Type指针, flag名, 默认值, 帮助信息)
		flag.StringVar(&config, "c", "", "请输入当前环境配置文件.")
		//
		// 定义好命令行flag参数后，需要通过调用flag.Parse()来对命令行参数进行解析。 如: go run main.go -c=config.release.yaml
		flag.Parse()

		// 判断命令行参数是否为空
		if config == "" {
			if configEnv := os.Getenv(constants.ConfigEnv); configEnv == "" {
				switch gin.Mode() {
				case gin.DebugMode:
					config = constants.ConfigDefaultFile
					fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.EnvGinMode, constants.ConfigDefaultFile)
				case gin.ReleaseMode:
					config = constants.ConfigReleaseFile
					fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.EnvGinMode, constants.ConfigReleaseFile)
				case gin.TestMode:
					config = constants.ConfigTestFile
					fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.EnvGinMode, constants.ConfigTestFile)
				default:
					config = constants.ConfigDebugFile
				}
			} else {
				// internal.ConfigEnv 常量存储的环境变量不为空 将值赋值于config
				config = configEnv
				fmt.Printf("您正在使用%s环境变量,config的路径为%s\n", constants.ConfigEnv, config)
			}
		} else {
			fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%s\n", config)
		}
	} else {
		config = path[0]
		fmt.Printf("您正在使用func Viper()传递的值,config的路径为%s\n", config)
	}

	vip := viper.New()
	vip.SetConfigFile(config)
	vip.SetConfigType("yaml")
	err := vip.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	vip.WatchConfig()
	vip.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = vip.Unmarshal(&global.AquilaConfig); err != nil {
			fmt.Println(err)
		}
	})
	if err = vip.Unmarshal(&global.AquilaConfig); err != nil {
		fmt.Println(err)
	}
	fmt.Println("====1-viper====: viper init config success")
	return vip
}
