package utils

import (
	"log"
	"os"
	"sync"

	"github.com/qianxia/blog/global"
	"github.com/spf13/viper"
)

func Viper() {
	var once sync.Once

	// 设置配置文件名
	viper.SetConfigName("config")
	// 指定配置文件的扩展名
	viper.SetConfigType("toml")
	// 设置配置文件路径
	workDir, _ := os.Getwd()
	viper.AddConfigPath(workDir + "/config")
	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Fatal error config file: %v", err)
	}
	once.Do(func() {
		// 反序列化到指定结构体上
		err := viper.Unmarshal(&global.QX_CONFIG)
		if err != nil {
			log.Fatalf("unable to read remote config: %v", err)
		}
	})
}
