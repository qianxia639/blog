package routers

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/handler"
	"github.com/qianxia/blog/initialize"
	"github.com/qianxia/blog/middleware"
	"github.com/qianxia/blog/utils"
)

type Option func(*gin.Engine) *gin.Engine

var options = []Option{}

// 注册app的路由配置
func include(opts ...Option) {
	options = append(options, opts...)
}

// 初始化
func Init() *gin.Engine {
	// 加载app的配置路由
	include(handler.ExampleRouters, handler.SystemRouters)

	r := gin.Default()
	r.Use(middleware.CORS())
	for _, opt := range options {
		opt(r)
	}
	return r
}

func Load() {
	// 读取yaml配置文件
	global.QX_YAML_CONFIG = utils.DeCode()

	// 加载log日志
	global.QX_LOG = utils.InitLogger("./log/info."+time.Now().Format("2006-01-02")+".log", global.QX_WARN_PATH, zap.InfoLevel)

	// 加载MySQL配置信息
	global.QX_DB = utils.InitDb(global.QX_YAML_CONFIG)
	if global.QX_DB != nil {
		// 初始化表
		initialize.RegisterTables(global.QX_DB)
	}
}
