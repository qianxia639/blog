package oss

import (
	"Blog/utils"
	"context"
	"os"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

type OssQiniu struct {
	AccessKey string
	SecretKey string
	Bucket    string
	ServerUrl string
}

func (o *OssQiniu) UploadImage(localfile string) (string, error) {
	putPolicy := storage.PutPolicy{
		Scope: o.Bucket,
	}

	mac := qbox.NewMac(o.AccessKey, o.SecretKey)
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
		return "", err
	}

	return o.ServerUrl + "/" + ret.Key, nil
}
