package middleware

import (
	"github.com/gin-gonic/gin"
)

func LogFuncExecTime() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// start := time.Now()
		// path := ctx.Request.URL.Path
		// raw := ctx.Request.URL.RawQuery

		// ctx.Next()

		// method := ctx.Request.Method
		// ip := ctx.ClientIP()

		// if raw != "" {
		// 	path += raw
		// }

		// statusCode := ctx.Writer.Status()

		// latency := time.Since(start).Milliseconds()

		// time | statusCode | timeSub | ip | method | path
		// logs.Logs.Info("%s | %d | %5dms | %s | %s | %s",
		// 	start.Format("2006/01/02 15:04:05"), statusCode, latency, ip, method, path,
		// )
	}
}
