package routers

import (
	"database/sql"
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

func Load() *sql.DB {
	// 读取yaml配置文件
	global.RY_YAML_CONFIG = utils.DeCode()

	// 加载log日志
	global.RY_LOG = utils.InitLogger("./log/info."+time.Now().Format("2006-01-02")+".log", global.RY_WARN_PATH, zap.InfoLevel)

	// 加载MySQL配置信息
	global.RY_DB = utils.InitDb(global.RY_YAML_CONFIG)
	if global.RY_DB != nil {
		// 初始化表
		initialize.RegisterTables(global.RY_DB)
		db, _ := global.RY_DB.DB()
		// 关闭mysql连接
		return db
	}
	return nil
}
