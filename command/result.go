package command

import (
	"github.com/gin-gonic/gin"
)

func result(ctx *gin.Context, httpStatus int, state bool, msg string, data interface{}) {
	ctx.SecureJSON(httpStatus, gin.H{"state": state, "msg": msg, "data": data})
}

func resultError(ctx *gin.Context, httpStatus int, state bool, msg string) {
	ctx.SecureJSON(httpStatus, gin.H{"state": state, "msg": msg})
}

func Success(ctx *gin.Context, msg string, data interface{}) {
	result(ctx, 200, true, msg, data)
}

func Failed(ctx *gin.Context, httpStatus int, msg string) {
	resultError(ctx, httpStatus, false, msg)
}

func RFailed(ctx *gin.Context, httpStatus int, msg string) {
	resultError(ctx, httpStatus, false, msg)
	return
}
