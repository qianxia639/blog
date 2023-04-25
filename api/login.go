package api

import (
	"Blog/core/errors"
	"Blog/core/logs"
	"Blog/utils"
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
	Password string `json:"password"`
}

type loginResponse struct {
	Token string `json:"token"`
}

func (server *Server) login(ctx *gin.Context) {

	var req loginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		logs.Logs.Error("Bind Param err: ", err)
		ctx.SecureJSON(http.StatusBadRequest, errors.ParamErr.Error())
		return
	}

	user, err := server.store.GetUser(ctx, req.Username)
	if err != nil {
		if err == ErrNoRows {
			logs.Logs.Error("Not User err: ", err)
			ctx.SecureJSON(http.StatusNotFound, errors.NotExistsUserErr.Error())
			return
		}
		ctx.SecureJSON(http.StatusInternalServerError, errors.ServerErr.Error())
		return
	}

	loginAttemptskey := fmt.Sprintf("loginAttempts:%s", user.Username)
	lockedAccountKey := fmt.Sprintf("lockedAccount:%s", user.Username)

	if err := utils.Decrypt(req.Password, user.Password); err != nil {
		logs.Logs.Error(err)
		// 增加登录尝试次数
		attempts, err := server.rdb.Incr(ctx, loginAttemptskey).Result()
		if err != nil {
			logs.Logs.Error("redis err: ", err)
			ctx.JSON(http.StatusInternalServerError, errors.ServerErr.Error())
			return
		}

		if attempts >= maxLoginAttempts {
			// 锁定用户1小时
			err = server.rdb.Set(ctx, lockedAccountKey, true, time.Hour).Err()
			if err != nil {
				logs.Logs.Error("redis err: ", err)
				ctx.JSON(http.StatusInternalServerError, errors.ServerErr.Error())
				return
			}
			ctx.JSON(http.StatusUnauthorized, errors.AccountLockedErr.Error())
			return
		}

		ctx.SecureJSON(http.StatusUnauthorized, errors.PasswordErr.Error())
		return
	}

	locked, err := server.rdb.Get(ctx, lockedAccountKey).Bool()
	if err != nil && err != redis.Nil {
		ctx.SecureJSON(http.StatusInternalServerError, errors.ServerErr.Error())
		return
	}

	if locked {
		ctx.SecureJSON(http.StatusUnauthorized, errors.AccountLockedErr.Error())
		return
	}

	// 重置登录尝试次数
	if err := server.rdb.Del(ctx, loginAttemptskey).Err(); err != nil {
		logs.Logs.Error("Del Redis err: ", err)
		ctx.SecureJSON(http.StatusInternalServerError, errors.ServerErr.Error())
		return
	}

	token, err := server.maker.CreateToken(user.Username, server.conf.Token.AccessTokenDuration)
	if err != nil {
		ctx.SecureJSON(http.StatusInternalServerError, errors.ServerErr.Error())
		return
	}

	ctx.JSON(http.StatusOK, loginResponse{Token: token})
}
