package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/qianxia/blog/global"
)

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		method := ctx.Request.Method
		// body, _ := ioutil.ReadAll(ctx.Request.Body)

		global.QX_LOG.Infof("请求的信息: { %s | %s | {%s}}", method, path)

		ctx.Next()
	}
}
