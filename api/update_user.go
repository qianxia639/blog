package api

import (
	"Blog/core/errors"
	"Blog/core/logs"
	"Blog/core/oss"
	"Blog/core/result"
	db "Blog/db/sqlc"
	"Blog/utils"
	"database/sql"
	"fmt"
	"mime"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"go.uber.org/zap"
)

type updateUserRequest struct {
	Username string  `json:"username" binding:"required"`
	Email    *string `json:"email"`
	Nickname *string `json:"nickname"`
	Password *string `json:"password"`
	Avatar   *string `json:"avatar"`
}

func (server *Server) updateUser(ctx *gin.Context) {
	var req updateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		logs.Logs.Error("bind pramar err", zap.Error(err))
		result.ParamError(ctx, errors.ParamErr.Error())
		return
	}

	payload, err := server.readToken(ctx.Request)
	if err != nil {
		result.UnauthorizedError(ctx, err.Error())
		return
	}

	if req.Username != payload.Username {
		result.UnauthorizedError(ctx, errors.UnauthorizedErr.Error())
		return
	}

	arg := &db.UpdateUserParams{
		Username: req.Username,
		Email:    newNullString(req.Email),
		Nickname: newNullString(req.Nickname),
		Avatar:   newNullString(req.Avatar),
	}

	if req.Password != nil {
		hashPassword, err := utils.Encrypt(*req.Password)
		if err != nil {
			result.ServerError(ctx, errors.ServerErr.Error())
			return
		}
		arg.Password = sql.NullString{
			String: hashPassword,
			Valid:  true,
		}
	}

	if req.Avatar != nil {
		fileUrl, err := server.upload(*req.Avatar)
		if err != nil {
			logs.Logs.Error("upload file err", zap.Error(err))
			if wr, ok := err.(*errors.WrapError); ok {
				switch err {
				case wr:
					result.ServerError(ctx, wr.Error())
					return
				}
			}
			result.ServerError(ctx, errors.ServerErr.Error())
			return
		}
		arg.Avatar = newNullString(&fileUrl)
	}

	user, err := server.store.UpdateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case errors.UniqueViolationErr:
				result.Error(ctx, http.StatusForbidden, errors.NicknameExistsErr.Error())
				return
			}
		}
		result.ServerError(ctx, errors.ServerErr.Error())
		return
	}

	key := fmt.Sprintf("t_%s", user.Username)
	err = server.rdb.Set(ctx, key, &user, 24*time.Hour).Err()
	if err != nil {
		logs.Logs.Error("redis err: ", zap.Error(err))
		result.ServerError(ctx, errors.ServerErr.Error())
		return
	}

	result.OK(ctx, "Update Successfully")
}

func (server *Server) upload(localFile string) (string, error) {
	logs.Logs.Info("upload filename", zap.String("filename", localFile))

	buf, err := os.ReadFile(localFile)
	if err != nil {
		return "", err
	}

	if len(buf) > utils.ImageMaxSize {
		return "", errors.ExceedingLenggthErr
	}

	contentType := http.DetectContentType(buf)
	mediaType, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		return "", err
	}
	if _, ok := utils.GetInstance()[contentType]; !ok {
		return "", errors.NewWrapError("只支持上传图片(.gif也不行)，当前文件类型为: " + mediaType)
	}

	up := oss.Upload{
		ImageStrategy: &oss.OssQiniu{
			Conf: server.conf.OssQiniu,
		},
		LocalFile: localFile,
	}
	return up.UploadImage()
}
