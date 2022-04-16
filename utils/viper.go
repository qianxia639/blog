package utils

import (
	"github.com/qianxia/blog/global"
	"github.com/spf13/viper"
)

func Viper() {
	v := viper.New()
	// 设置配置文件路径
	v.SetConfigFile("./config/config.toml")
	// 读取配置文件
	err := v.ReadInConfig()
	if err != nil {
		global.QX_LOG.Fatalf("Fatal error config file: %v", err)
		return
	}
	// 指定配置文件的扩展名
	v.SetConfigType("toml")

	// 反序列化到指定结构体上
	err = v.Unmarshal(&global.QX_CONFIG)
	if err != nil {
		global.QX_LOG.Fatalf("unable to read remote config: %v", err)
		return
	}
}
