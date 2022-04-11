package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/qianxia/blog/command"
	"github.com/qianxia/blog/global"
	uuid "github.com/satori/go.uuid"
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
func GetUserUUID(ctx *gin.Context) uuid.UUID {
	if claims, exists := ctx.Get("claims"); !exists {
		if cl, err := GetClaims(ctx); err != nil {
			return uuid.UUID{}
		} else {
			return cl.UUID
		}
	} else {
		return claims.(*CustomClaims).UUID
	}
}

// 从Gin的Context中获取jwt并解析其中的用户名
func GetUsername(ctx *gin.Context) string {
	if claims, exists := ctx.Get("claims"); !exists {
		if cl, err := GetClaims(ctx); err != nil {
			return ""
		} else {
			return cl.Username
		}
	} else {
		return claims.(*CustomClaims).Username
	}
}

// 从Gin的Context中获取jwt并解析其中的用户头像
func GetAvatar(ctx *gin.Context) string {
	if claims, exists := ctx.Get("claims"); !exists {
		if cl, err := GetClaims(ctx); err != nil {
			return ""
		} else {
			return cl.Avatar
		}
	} else {
		return claims.(*CustomClaims).Avatar
	}
}
