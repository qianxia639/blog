package system

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qianxia/blog/command"
	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/utils"
)

type UploadHandler struct{}

func (uh *UploadHandler) UploadMdFile(ctx *gin.Context) {
	file, fileHeader, err := ctx.Request.FormFile("file")
	if err != nil {
		global.QX_LOG.Error(err)
		command.Failed(ctx, http.StatusInternalServerError, "上传失败")
		return
	}

	if url, err := utils.UploadFile(file, fileHeader.Size); err != nil {
		global.QX_LOG.Error(err)
		command.Failed(ctx, http.StatusInternalServerError, "上传失败")
		return
	} else {
		ctx.SecureJSON(http.StatusOK, gin.H{"url": url})
	}
}
