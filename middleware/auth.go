package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/qianxia/blog/command"
	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/model"
	"github.com/qianxia/blog/utils"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 从请求头中获取X-Token头信息
		token := ctx.GetHeader("X-Token")
		if token == "" {
			command.Failed(ctx, http.StatusUnauthorized, "未登录或非法访问")
			ctx.Abort()
			return
		}
		if !strings.HasPrefix(token, "Bearer ") {
			command.Failed(ctx, http.StatusUnauthorized, "token格式有误")
			ctx.Abort()
			return
		}
		token = token[7:]

		// 解析token
		claims, err := utils.ParseJwt(token)
		if err != nil {
			command.Failed(ctx, http.StatusUnauthorized, err.Error())
			ctx.Abort()
			return
		}

		// 验证通过或获取claims中的userId
		var user model.User
		if err := global.RY_DB.Debug().Select("id,username,email,avatar").Where("id = ?", claims.UserId).Find(&user).Error; err != nil {
			command.Failed(ctx, http.StatusInternalServerError, "用户名不存在")
			ctx.Abort()
			return
		}

		// 将用户信息写入上下文中
		ctx.Set("user", user)
		ctx.Next()
	}
}
