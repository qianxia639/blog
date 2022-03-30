package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qianxia/blog/command"
	"github.com/qianxia/blog/utils"
)

func CasbinMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claims, _ := utils.GetClaims(ctx)
		// 获取请求路径
		obj := ctx.Request.URL.Path
		// 获取请求方式
		act := ctx.Request.Method
		// 获取用户角色
		sub := claims.BaseClaims.Id
		e := utils.Casbin()
		// 判断策略
		success, _ := e.Enforce(sub, obj, act)
		if success {
			ctx.Next()
		} else {
			command.Failed(ctx, http.StatusForbidden, "权限不足")
			ctx.Abort()
			return
		}
	}
}
