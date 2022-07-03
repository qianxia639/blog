package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qianxia/blog/command"
	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/utils"
)

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 从请求头中获取X-Token头信息
		token := ctx.Request.Header.Get("X-Token")

		if token == "" {
			global.LOG.Error("token 为空")
			command.Failed(ctx, http.StatusUnauthorized, "未登录或非法访问")
			ctx.Abort()
			return
		}

		// 解析token
		claims, err := utils.ParseToken(token)
		if err != nil {
			global.LOG.Error(err)
			command.Failed(ctx, http.StatusUnauthorized, err.Error())
			ctx.Abort()
			return
		}

		if claims.UUID == "" {
			global.LOG.Error("用户UUID为空")
			command.Failed(ctx, http.StatusUnauthorized, "未登录或非法访问")
			ctx.Abort()
			return
		}

		// 将token信息写入Gin的Context中
		ctx.Set("claims", claims)
		ctx.Next()
	}
}
