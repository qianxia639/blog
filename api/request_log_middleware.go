package api

import (
	db "Blog/db/sqlc"
	"bytes"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
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

			contentType := ctx.Request.Header.Get("Content-Type")
			cost := time.Since(start).Milliseconds()

			arg := db.InsertRequestLogParams{
				Method:       method,
				Path:         path,
				StatusCode:   int32(statusCode),
				Ip:           ip,
				Hostname:     hostname,
				RequestBody:  body,
				ResponseTime: cost,
				UserAgent:    ua,
				ContentType:  contentType,
			}

			server.store.InsertRequestLog(ctx, arg)
		}()
	}
}

func getHostname() string {
	hostname, _ := os.Hostname()
	return hostname
}
