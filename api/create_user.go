package api

import (
	"Blog/core/errors"
	"Blog/core/logs"
	"Blog/core/result"
	"Blog/core/task"
	db "Blog/db/sqlc"
	"Blog/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
	"github.com/lib/pq"
	"go.uber.org/zap"
)

type createUserRequest struct {
	Username string `json:"username" binding:"alphanum"`
	Email    string `json:"email" binding:"email"`
	Password string `json:"password" binding:"min=6,max=20"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		logs.Logs.Error("bind params", zap.Error(err))
		result.ParamError(ctx, errors.ParamErr.Error())
		return
	}

	hashPassword, err := utils.Encrypt(req.Password)
	if err != nil {
		logs.Logs.Error("encrypt password", zap.Error(err))
		result.ServerError(ctx, errors.ServerErr.Error())
		return
	}

	arg := &db.CreateUserTxParams{
		CreateUserParams: db.CreateUserParams{
			Username:     req.Username,
			Email:        req.Email,
			Nickname:     req.Email,
			Password:     hashPassword,
			RegisterTime: time.Now(),
		},
		AfterCreate: func(user db.User) error {
			taskPayload := &task.SendVerifyEmailPayload{
				Email: user.Email,
			}

			opts := []asynq.Option{
				asynq.MaxRetry(10),
				asynq.ProcessIn(time.Duration(utils.RandomInt(5, 10)) * time.Second),
				asynq.Queue(task.QueueCritical),
			}

			return server.taskDistributor.DistributeTaskSendVerifyEmail(ctx, taskPayload, opts...)
		},
	}

	_, err = server.store.CreateUserTx(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case errors.UniqueViolationErr:
				result.Error(ctx, http.StatusForbidden, errors.UsernameOrEmailEexistsErr.Error())
				return
			}
		}
		result.ServerError(ctx, errors.ServerErr.Error())
		return
	}

	result.OK(ctx, "Create Successfully")
}
