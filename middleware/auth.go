package middleware

import (
	"fmt"
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
		fmt.Println("开始执行 Auth 中间件")
		// 从请求头中获取X-Token头信息
		token := ctx.GetHeader("X-Token")
		if token == "" {
			global.RY_LOG.Error("未登录或非法访问")
			command.Failed(ctx, http.StatusUnauthorized, "未登录或非法访问")
			ctx.Abort()
			return
		}
		if !strings.HasPrefix(token, "Bearer ") {
			global.RY_LOG.Error("token格式有误")
			command.Failed(ctx, http.StatusUnauthorized, "token格式有误")
			ctx.Abort()
			return
		}
		token = token[7:]

		// 解析token
		claims, err := utils.ParseToken(token)
		if err != nil {
			global.RY_LOG.Error(err)
			command.Failed(ctx, http.StatusUnauthorized, err.Error())
			ctx.Abort()
			return
		}
		fmt.Printf("程序执行到了这里且 claims = %v\n", claims)
		// 验证通过或获取claims中的userId
		var user model.User
		if err := global.RY_DB.Debug().Select("id,username,avatar").Where("username = ?", claims.Username).Find(&user).Error; err != nil {
			command.Failed(ctx, http.StatusInternalServerError, "用户名不存在")
			ctx.Abort()
			return
		}

		fmt.Println("set 前 user ===> ", user)

		// 将用户信息写入上下文中
		ctx.Set("user", user)

		fmt.Println("set 后 user ===> ", user)

		ctx.Next()
	}
}
