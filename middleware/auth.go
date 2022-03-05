package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/qianxia/blog/command"
	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/utils"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 从请求头中获取X-Token头信息
		token := ctx.Request.Header.Get("X-Token")
		if token == "" {
			global.QX_LOG.Error("未登录或非法访问")
			command.Failed(ctx, http.StatusUnauthorized, "未登录或非法访问")
			ctx.Abort()
			return
		}
		if !strings.HasPrefix(token, "Bearer ") {
			global.QX_LOG.Error("token格式有误")
			command.Failed(ctx, http.StatusUnauthorized, "token格式有误")
			ctx.Abort()
			return
		}
		token = token[7:]

		// 解析token
		claims, err := utils.ParseToken(token)
		if err != nil {
			global.QX_LOG.Error(err)
			command.Failed(ctx, http.StatusUnauthorized, err.Error())
			ctx.Abort()
			return
		}

		// // 将token信息写入Gin的Context中
		ctx.Set("claims", claims)
		ctx.Next()
	}
}
