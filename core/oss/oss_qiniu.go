package oss

import (
	"Blog/core/config"
	"Blog/core/logs"
	"Blog/utils"
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"go.uber.org/zap"
)

type OssQiniu struct {
	Conf config.OssQiniu
}

func (o *OssQiniu) UploadImage(localfile string) (string, error) {
	buf, err := os.ReadFile(localfile)
	if err != nil {
		return "", err
	}

	contentType := http.DetectContentType(buf)
	if _, ok := utils.ImageTypes[contentType]; !ok {
		return "", fmt.Errorf("文件类型错误,当前文件类型为: %s", contentType)
	}

	putPolicy := storage.PutPolicy{
		Scope: o.Conf.Bucket,
	}

	mac := qbox.NewMac(o.Conf.AccessKey, o.Conf.SecretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{
		Region:        &storage.ZoneHuanan, // 空间对应的域名
		UseHTTPS:      true,                // 是否使用https域名
		UseCdnDomains: false,               // 上传是否使用CDN加速
	}

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	err = formUploader.PutFileWithoutKey(context.Background(), &ret, upToken, localfile, nil)
	if err != nil {
		logs.Logs.Error("failed put file", zap.Error(err))
		return "", err
	}

	logs.Logs.Info("upload success", zap.String("file", localfile), zap.String("key", ret.Key))

	return o.Conf.ServerUrl + "/" + ret.Key, nil
}
