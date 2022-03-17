package utils

import (
	"context"
	"mime/multipart"

	"github.com/qianxia/blog/global"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

// 文件上传(七牛云存储)
func UploadFile(file multipart.File, fileSize int64) (string, error) {
	mac := qbox.NewMac(global.QX_CONFIG.Qiniu.AccessKey, global.QX_CONFIG.Qiniu.SecretKey)
	putPolicy := storage.PutPolicy{
		Scope: global.QX_CONFIG.Qiniu.Bucket,
	}

	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{
		// 空间对应的机房
		Zone: &storage.ZoneHuanan,
		// 是否使用https域名
		UseHTTPS: false,
		// 上传是否使用CDS加速
		UseCdnDomains: false,
	}

	formUpload := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	if err := formUpload.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, nil); err != nil {
		return "", err
	} else {
		return global.QX_CONFIG.Qiniu.ServerUrl + ret.Key, nil
	}
}
