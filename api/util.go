package api

import (
	"Blog/core/logs"
	"Blog/core/token"
	"context"
	"net/http"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

const (
	authorizationHeader = "Authorization"
)

// read token
func (s *Server) readToken(r *http.Request) (*token.Payload, error) {
	token := r.Header.Get(authorizationHeader)
	payload, err := s.maker.VerifyToken(token)
	if err != nil {
		return nil, err
	}

	return payload, nil
}

// upload file
func (s *Server) uploadFile(localFile string) (string, error) {

	// 判断是网络文件还是本地文件
	// if strings.HasPrefix(localFile, "http://") || strings.HasPrefix(localFile, "https://") {
	// 	return "", errors.New("暂不支持网络文件上传")
	// }

	putPolicy := storage.PutPolicy{
		Scope: s.conf.OssQiniu.Bucket,
	}

	mac := qbox.NewMac(s.conf.OssQiniu.AccessKey, s.conf.OssQiniu.SecretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{
		Region:        &storage.ZoneHuanan, // 空间对应的域名
		UseHTTPS:      true,                // 是否使用https域名
		UseCdnDomains: false,               // 上传是否使用CDN加速
	}

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	err := formUploader.PutFileWithoutKey(context.Background(), &ret, upToken, localFile, nil)
	if err != nil {
		return "", err
	}
	logs.Logs.Infoln("localFile: ", localFile)

	return s.conf.OssQiniu.ServerUrl + "/" + ret.Key, nil
}
