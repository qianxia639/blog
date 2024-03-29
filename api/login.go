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
	"go.uber.org/zap"
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
		logs.Logger.Error("Bind Param err: ", zap.Error(err))
		result.ParamError(ctx, errors.ParamErr.Error())
		return
	}

	loginAttemptsKey := fmt.Sprintf("loginAttempts:%s", req.Username)
	lockedAccountKey := fmt.Sprintf("lockedAccount:%s", req.Username)

	locked, err := server.cache.Get(ctx, lockedAccountKey).Bool()
	if err != nil && err != redis.Nil {
		logs.Logger.Error("get locked account err: ", zap.Error(err))
		result.ServerError(ctx, errors.ServerErr.Error())
		return
	}

	if locked {
		result.UnauthorizedError(ctx, errors.AccountLockedErr.Error())
		return
	}

	user, err := server.store.GetUser(ctx, req.Username)
	if err != nil {
		switch err {
		case errors.NoRowsErr:
			logs.Logger.Error("Not User err: ", zap.Error(err))
			result.UnauthorizedError(ctx, errors.NotExistsUserErr.Error())
		default:
			result.ServerError(ctx, errors.ServerErr.Error())
		}
		return
	}

	if err := utils.Decrypt(req.Password, user.Password); err != nil {
		if statusCode, err := server.accountLocked(ctx, loginAttemptsKey, lockedAccountKey); err != nil && statusCode != http.StatusOK {
			// 重置失败的登录次数
			if err := server.resetLoginAttempts(ctx, loginAttemptsKey); err != nil {
				logs.Logger.Error("reset login attempts", zap.Error(err))
				result.ServerError(ctx, errors.ServerErr.Error())
			}
			result.Error(ctx, statusCode, err.Error())
			return
		}

		logs.Logger.Error("Decrypt Password", zap.Error(err))
		result.UnauthorizedError(ctx, errors.PasswordErr.Error())
		return
	}

	// 重置登录失败次数
	if err := server.resetLoginAttempts(ctx, loginAttemptsKey); err != nil {
		logs.Logger.Error("reset login attempts", zap.Error(err))
		result.ServerError(ctx, errors.ServerErr.Error())
		return
	}

	token, err := server.maker.CreateToken(user.Username, server.conf.Token.AccessTokenDuration)
	if err != nil {
		result.ServerError(ctx, errors.ServerErr.Error())
		return
	}

	key := fmt.Sprintf("t_%s", user.Username)
	err = server.cache.Set(ctx, key, &user, 24*time.Hour).Err()
	if err != nil {
		logs.Logger.Error("redis err: ", zap.Error(err))
		result.ServerError(ctx, errors.ServerErr.Error())
		return
	}

	result.Obj(ctx, token)
}

func (server *Server) accountLocked(ctx context.Context, loginAttemptsKey, lockedAccountKey string) (int, error) {
	// 增加登录尝试次数
	attempts, err := server.cache.Incr(ctx, loginAttemptsKey).Result()
	if err != nil {
		logs.Logger.Error("redis err: ", zap.Error(err))
		return http.StatusInternalServerError, errors.ServerErr
	}

	if attempts > maxLoginAttempts {
		// 锁定用户1小时
		if err := server.cache.Set(ctx, lockedAccountKey, true, time.Hour).Err(); err != nil {
			logs.Logger.Error("set locked account", zap.Error(err))
			return http.StatusInternalServerError, errors.ServerErr
		}
		return http.StatusUnauthorized, errors.AccountLockedErr
	}

	return http.StatusOK, nil
}

func (server *Server) resetLoginAttempts(ctx context.Context, loginAttemptsKey string) error {
	return server.cache.Del(ctx, loginAttemptsKey).Err()
}
