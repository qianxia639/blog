package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/qianxia/blog/command"
	"github.com/qianxia/blog/global"
)

func GetClaims(ctx *gin.Context) (*CustomClaims, error) {
	token := ctx.Request.Header.Get("X-Token")
	claims, err := ParseToken(token)
	if err != nil {
		global.QX_LOG.Error("解析jwt信息失败,请检查请求头是否存在X-Token")
		command.Failed(ctx, 401, err.Error())
		return nil, err
	}
	return claims, nil
}

// 从Gin的Context中获取jwt并解析其中的用户Id
func GetUserId(ctx *gin.Context) uint64 {
	if claims, exists := ctx.Get("claims"); !exists {
		if cl, err := GetClaims(ctx); err != nil {
			return 0
		} else {
			return cl.Id
		}
	} else {
		return claims.(*CustomClaims).Id
	}
}

// 从Gin的Context中获取jwt并解析其中的用户UUID
func GetUserUUID(ctx *gin.Context) string {
	if claims, exists := ctx.Get("claims"); !exists {
		if cl, err := GetClaims(ctx); err != nil {
			return ""
		} else {
			return cl.UUID
		}
	} else {
		return claims.(*CustomClaims).UUID
	}
}
