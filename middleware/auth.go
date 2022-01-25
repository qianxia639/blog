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
		// 从请求头中获取Authorization头信息
		tokenStr := ctx.GetHeader("Authorization")
		if tokenStr == "" || !strings.HasPrefix(tokenStr, "Bearer ") {
			command.Failed(ctx, http.StatusUnauthorized, "token不存在或格式不正确")
			ctx.Abort()
			return
		}

		tokenStr = tokenStr[7:]
		// 常量
		// 解析token
		token, claims, err := utils.ParseJwt(tokenStr)
		if err != nil || !token.Valid {
			command.Failed(ctx, http.StatusUnauthorized, "token解析失败")
			ctx.Abort()
			return
		}

		// 验证通过或获取claims中的userId
		// userId := claims.UserId
		// DB := utils.GetDB()
		var user model.User
		// if err := DB.Raw("SELECT id,username,email,avatar FROM "+command.DBUser+" WHERE id = ?", userId).Scan(&user).Error; err != nil {
		// 	command.Failed(ctx, http.StatusInternalServerError, "用户名不存在")
		// 	ctx.Abort()
		// 	return
		// }
		if err := global.RY_DB.Debug().Select("id,username,email,avatar").Where("id = ?", claims.UserId).Find(&user).Error; err != nil {
			command.Failed(ctx, http.StatusInternalServerError, "用户名不存在")
			ctx.Abort()
			return
		}
		// if err := DB.Table(command.DBUser).First(&user, userId).Error; err != nil {
		// 	command.Failed(ctx, http.StatusInternalServerError, "用户名不存在")
		// 	ctx.Abort()
		// 	return
		// }

		// 将用户信息写入上下文
		ctx.Set("user", user)
		ctx.Next()
	}
}
