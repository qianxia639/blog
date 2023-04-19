package api

import (
	"Blog/core/logs"
	"bytes"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (server *Server) requestLogMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()

		b, _ := ctx.GetRawData()

		body := bytes.NewBuffer(b).String()

		ctx.Request.Body = io.NopCloser(bytes.NewBuffer(b))

		ctx.Next()

		defer func() {
			method := ctx.Request.Method
			ip := ctx.ClientIP()
			path := ctx.Request.URL.Path
			raw := ctx.Request.URL.RawQuery
			hostname := getHostname()

			if raw != "" {
				path += raw
			}

			ua := ctx.Request.Header.Get("User-Agent")

			statusCode := ctx.Writer.Status()

			contentType := ctx.ContentType()
			cost := time.Since(start).Milliseconds()

			logs.Logs.Info("Request Log",
				zap.String("Method", method),
				zap.String("Path", path),
				zap.Int32("StatusCode", int32(statusCode)),
				zap.String("Ip", ip),
				zap.String("Hostname", hostname),
				zap.String("RequestBody", body),
				zap.Int64("ResponseTime", cost),
				zap.String("UserAgent", ua),
				zap.String("ContentType", contentType),
			)

			// arg := db.InsertRequestLogParams{
			// 	Method:       method,
			// 	Path:         path,
			// 	StatusCode:   int32(statusCode),
			// 	Ip:           ip,
			// 	Hostname:     hostname,
			// 	RequestBody:  body,
			// 	ResponseTime: cost,
			// 	UserAgent:    ua,
			// 	ContentType:  contentType,
			// }

			// server.store.InsertRequestLog(ctx, arg)
		}()
	}
}

func getHostname() string {
	hostname, _ := os.Hostname()
	return hostname
}
