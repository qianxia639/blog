package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	AccessControlAllowOrigin      = "Access-Control-Allow-Origin"
	AccessControlMaxAge           = "Access-Control-Max-Age"
	AccessControlAllowMethods     = "Access-Control-Allow-Methods"
	AccessControlAllowHeaders     = "Access-Control-Allow-Headers"
	AccessControlAllowCredentials = "Access-Control-Allow-Credentials"
	ContenType                    = "Content-Type"
)

// 跨域 -- CORS
func CORS() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set(AccessControlAllowOrigin, "*")
		ctx.Writer.Header().Set(AccessControlMaxAge, "86400")
		ctx.Writer.Header().Set(AccessControlAllowMethods, "POST,GET,PUT,DELETE,OPTIONS")
		ctx.Writer.Header().Set(AccessControlAllowHeaders, "Content-Type,X-CSRF-Token,X-Token,Authorization")
		ctx.Writer.Header().Set(AccessControlAllowCredentials, "true")
		ctx.Writer.Header().Set(ContenType, "application/json;charset=utf-8")
		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(http.StatusNoContent)
		}
		ctx.Next()
	}
}
