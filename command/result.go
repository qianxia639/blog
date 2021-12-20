package command

import (
	"github.com/gin-gonic/gin"
)

func result(ctx *gin.Context, httpStatus, code int, msg string, data interface{}) {
	ctx.JSON(httpStatus, gin.H{"code": code, "msg": msg, "data": data})
}

func resultError(ctx *gin.Context, httpStatus, code int, msg string) {
	ctx.JSON(httpStatus, gin.H{"code": code, "msg": msg})
}

func Success(ctx *gin.Context, msg string, data interface{}) {
	result(ctx, 200, 200, msg, data)
}

func Failed(ctx *gin.Context, httpStatus, code int, msg string) {
	resultError(ctx, httpStatus, code, msg)
}
