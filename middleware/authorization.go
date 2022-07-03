package middleware

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/qianxia/blog/command"
	"github.com/qianxia/blog/utils"
)

func Authorization(e *casbin.Enforcer) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取请求PATH
		obj := ctx.FullPath()

		// 获取请求方式
		act := ctx.Request.Method

		claims, _ := utils.GetClaims(ctx)
		// 获取用户角色
		sub := strconv.Itoa(int(claims.RoleId))
		fmt.Printf("obj = [%s], act = [%s], sub = [%s]\n", obj, act, sub)
		// 校验策略
		success, _ := e.Enforce(sub, obj, act)
		fmt.Printf("success: %v\n", success)
		if !success {
			command.Failed(ctx, http.StatusForbidden, "权限不足")
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
