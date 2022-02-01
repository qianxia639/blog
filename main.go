package main

import (
	"fmt"

	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/initialize"
	"github.com/qianxia/blog/routers"
	"github.com/qianxia/blog/utils"
	"gopkg.in/yaml.v2"
)

func main() {
	// 初始化路由
	r := routers.Init()

	// 读取yaml配置文件
	dc, y := utils.DeCode()
	yaml.Unmarshal(dc, &y)

	// 加载MySQL配置信息
	global.RY_DB = utils.InitDb(y)
	if global.RY_DB != nil {
		// 初始化表
		initialize.RegisterTables(global.RY_DB)
		db, _ := global.RY_DB.DB()
		// 关闭mysql连接
		defer db.Close()
	}
	// 加载Redis链接池配置信息
	// Pool := utils.InitRedis(y)
	// // 关闭redis链接池信息
	// defer Pool.Close()

	if y.Server.Port != -1 {
		panic(r.Run(fmt.Sprintf("%s:%d", y.Server.Host, y.Server.Port)))
	}

	go func() {
		panic(r.Run())
	}()
}
