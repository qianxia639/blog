package oss

import (
	"Blog/core/config"
	"Blog/core/logs"
	"context"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"go.uber.org/zap"
)

type OssQiniu struct {
	Conf config.OssQiniu
}

func (o *OssQiniu) UploadImage(localfile string) (string, error) {
	upToken, formUploader := o.generate()
	ret := storage.PutRet{}
	err := formUploader.PutFileWithoutKey(context.Background(), &ret, upToken, localfile, nil)
	if err != nil {
		logs.Logs.Error("failed put file", zap.Error(err))
		return "", err
	}

	logs.Logs.Info("upload success", zap.String("file", localfile), zap.String("key", ret.Key))

	return o.Conf.ServerUrl + "/" + ret.Key, nil
}

func (o *OssQiniu) generate() (string, *storage.FormUploader) {
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
	return upToken, storage.NewFormUploader(&cfg)
}
