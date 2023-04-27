package api

import (
	"Blog/core/errors"
	"Blog/core/logs"
	"Blog/core/result"
	"Blog/utils"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

// max login attempts count
const maxLoginAttempts = 5

type loginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (server *Server) login(ctx *gin.Context) {

	var req loginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		logs.Logs.Error("Bind Param err: ", err)
		result.BadRequestError(ctx, errors.ParamErr.Error())
		return
	}

	user, err := server.store.GetUser(ctx, req.Username)
	if err != nil {
		switch err {
		case ErrNoRows:
			logs.Logs.Error("Not User err: ", err)
			result.Error(ctx, http.StatusNotFound, errors.NotExistsUserErr.Error())
		default:
			result.ServerError(ctx, errors.ServerErr.Error())
		}
	}

	loginAttemptsKey := fmt.Sprintf("loginAttempts:%s", user.Username)
	lockedAccountKey := fmt.Sprintf("lockedAccount:%s", user.Username)

	if err := utils.Decrypt(req.Password, user.Password); err != nil {
		if statusCode, err := server.accountLocked(ctx, loginAttemptsKey, lockedAccountKey); err != nil && statusCode != http.StatusOK {
			result.Error(ctx, statusCode, err.Error())
			return
		}

		logs.Logs.Error(err)
		result.UnauthorizedError(ctx, errors.PasswordErr.Error())
		return
	}

	locked, err := server.rdb.Get(ctx, lockedAccountKey).Bool()
	if err != nil && err != redis.Nil {
		logs.Logs.Error("get locked account err: ", err)
		result.ServerError(ctx, errors.ServerErr.Error())
		return
	}

	if locked {
		result.UnauthorizedError(ctx, errors.AccountLockedErr.Error())
		return
	}

	// 重置登录失败次数
	if err := server.resetLoginAttempts(ctx, loginAttemptsKey); err != nil {
		logs.Logs.Error(err)
		result.ServerError(ctx, errors.ServerErr.Error())
		return
	}

	token, err := server.maker.CreateToken(user.Username, server.conf.Token.AccessTokenDuration)
	if err != nil {
		result.ServerError(ctx, errors.ServerErr.Error())
		return
	}

	result.Obj(ctx, token)
}

func (server *Server) accountLocked(ctx context.Context, loginAttemptsKey, lockedAccountKey string) (int, error) {
	// 增加登录尝试次数
	attempts, err := server.rdb.Incr(ctx, loginAttemptsKey).Result()
	if err != nil {
		logs.Logs.Error("redis err: ", err)
		return http.StatusInternalServerError, errors.ServerErr
	}

	if attempts > maxLoginAttempts {
		// 锁定用户1小时
		if err := server.rdb.Set(ctx, lockedAccountKey, true, time.Hour).Err(); err == nil {
			// 重置失败的登录次数
			if err := server.resetLoginAttempts(ctx, loginAttemptsKey); err != nil {
				logs.Logs.Error(err)
				return http.StatusInternalServerError, errors.ServerErr
			}
			logs.Logs.Error(err)
			return http.StatusInternalServerError, errors.ServerErr
		}
		return http.StatusUnauthorized, errors.AccountLockedErr
	}

	return http.StatusOK, nil
}

func (server *Server) resetLoginAttempts(ctx context.Context, loginAttemptsKey string) error {
	return server.rdb.Del(ctx, loginAttemptsKey).Err()
}
