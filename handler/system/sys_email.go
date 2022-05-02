package system

import (
	"github.com/gin-gonic/gin"
	"github.com/qianxia/blog/utils"
)

type EmailHandler struct{}

func (eh *EmailHandler) SendMail(ctx *gin.Context) {
	utils.SendMail()
}
