package main

import (
	"fmt"

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
	db := utils.InitDb(y)
	// 关闭MySQL连接信息
	defer db.Close()

	// 加载Redis链接池配置信息
	// Pool := utils.InitRedis(y)
	// // 关闭redis链接池信息
	// defer Pool.Close()

	if y.Server.Port != -1 {
		panic(r.Run(fmt.Sprintf("%s:%d", y.Server.Host, y.Server.Port)))
	}

	panic(r.Run())
}
