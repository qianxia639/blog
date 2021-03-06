package system

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qianxia/blog/command"
	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/utils"
)

type UploadHandler struct{}

// @Summary      图片上传
// @Tags         System/Upload
// @Accept       json
// @Produce      json
// @Param		 file  formData	  file  true	"图片"
// @Success 	 200  {object}  string {data=url}
// @Security 	 X-Token
// @Router       /upload/mdFile [post]
func (uh *UploadHandler) UploadMdFile(ctx *gin.Context) {

	var url []string
	form, err := ctx.MultipartForm()
	if err != nil {
		global.LOG.Errorf("get err %s", err.Error())
		command.Failed(ctx, http.StatusBadRequest, "上传失败")
		return
	}

	files := form.File["files"]
	for index, file := range files {
		f, _ := files[index].Open()

		if u, err := utils.UploadFile(f, file.Size); err != nil {
			global.LOG.Error(err)
			command.Failed(ctx, http.StatusInternalServerError, "上传失败")
			return
		} else {
			url = append(url, u)
		}
	}
	ctx.SecureJSON(http.StatusOK, gin.H{"url": url})

	// file, fileHeader, err := ctx.Request.FormFile("file")
	// if err != nil {
	// 	global.LOG.Error(err)
	// 	command.Failed(ctx, http.StatusInternalServerError, "服务异常")
	// 	return
	// }

	// if url, err := utils.UploadFile(file, fileHeader.Size); err != nil {
	// 	global.LOG.Error(err)
	// 	command.Failed(ctx, http.StatusInternalServerError, "上传失败")
	// 	return
	// } else {
	// 	ctx.SecureJSON(http.StatusOK, gin.H{"url": url})
	// }
}
