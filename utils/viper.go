package utils

import (
	"os"

	"github.com/qianxia/blog/global"
	"github.com/spf13/viper"
)

func Viper() {
	// 设置配置文件名
	viper.SetConfigName("config")
	// 指定配置文件的扩展名
	viper.SetConfigType("toml")
	// 设置配置文件路径
	workDir, _ := os.Getwd()
	viper.AddConfigPath(workDir + "/config/test")

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
		// global.QX_LOG.Fatalf("Fatal error config file: %v", err)
		// return
	}

	// 反序列化到指定结构体上
	err := viper.Unmarshal(&global.QX_CONFIG)
	if err != nil {
		panic(err)
		// global.QX_LOG.Fatalf("unable to read remote config: %v", err)
		// return
	}
}
