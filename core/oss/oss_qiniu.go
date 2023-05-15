package oss

import (
	"Blog/core/config"
	"Blog/core/logs"
	"Blog/utils"
	"context"
	"os"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"go.uber.org/zap"
)

type OssQiniu struct {
	Conf config.OssQiniu
}

func (o *OssQiniu) UploadImage(localfile string) (string, error) {
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

	buf, err := os.ReadFile(localfile)
	if err != nil {
		return "", err
	}

	md5Hash := utils.Md5(buf)

	err = formUploader.PutFile(context.Background(), &ret, upToken, md5Hash, localfile, nil)
	if err != nil {
		logs.Logs.Error("failed put file", zap.Error(err))
		return "", err
	}

	logs.Logs.Info("upload success", zap.String("file", localfile), zap.String("key", md5Hash))

	return o.Conf.ServerUrl + "/" + ret.Key, nil
}
