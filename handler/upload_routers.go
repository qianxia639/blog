package handler

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

const (
	accessKey = "vz0iB4SnKnxXAEoxbW6TQvVxrfqOspnhYQzrPL3j"
	secretKey = "GufJ3mD3z13EFXAaW_Y_PzI4n-uj4_I7ax-Vz6Dr"
	bucket    = "lyy-blog"
)

func UploadRouters(e *gin.Engine) *gin.Engine {

	// 获取凭证
	e.GET("/upload/token", func(ctx *gin.Context) {
		mac := qbox.NewMac(accessKey, secretKey)
		putPolicy := storage.PutPolicy{
			Scope: bucket,
		}
		ctx.JSONP(200, map[string]string{
			"upload_token": putPolicy.UploadToken(mac),
		})
	})

	// 正式上传文件
	e.POST("/upload/putFile", func(ctx *gin.Context) {

		mac := qbox.NewMac(accessKey, secretKey)
		putPolicy := storage.PutPolicy{
			Scope: bucket,
		}
		token := putPolicy.UploadToken(mac)

		// var uploadMap map[string]string

		header, err := ctx.FormFile("file")
		// fmt.Printf("file = %v, err = %s\n", file, err.Error())

		fmt.Printf(", err = %s\n", err.Error())

		fname := header.Filename
		f := header.Header.Get("file")

		fmt.Printf("fname = %s, f = %s\n", fname, f)

		cfg := storage.Config{}

		// 空间对应的机房
		// Zone: &storage.ZoneHuanan,
		// // 是否使用https域名
		// UseHTTPS: true,
		// // 上传是否使用CDS加速
		// UseCdnDomains: false,

		// 空间对应的机房
		cfg.Zone = &storage.ZoneHuabei
		// 是否使用https域名
		cfg.UseHTTPS = true
		// 上传是否使用CDN加速
		cfg.UseCdnDomains = false
		// 构建表单上传对象
		// file := "/c/Users/Administrator/Desktop/bg.jpe"
		formUpload := storage.NewFormUploader(&cfg)
		ret := storage.PutRet{}
		if err := formUpload.PutFile(context.Background(), &ret, token, "", "", nil); err != nil {
			// 出错了
			ctx.JSON(500, err)
		} else {
			ctx.JSON(200, map[string]string{
				"key":  ret.Key,
				"hash": ret.Hash,
			})
		}
	})

	return e
}
