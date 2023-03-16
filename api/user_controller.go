package api

import (
	db "Blog/db/sqlc"
	"Blog/utils"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type CreateUserRequest struct {
	Username string `json:"username" binding:"alphanum"`
	Email    string `json:"email" binding:"email"`
	Password string `json:"password" binding:"min=6,max=20"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.SecureJSON(http.StatusBadRequest, err.Error())
		return
	}

	hashPassword, err := utils.Encrypt(req.Password)
	if err != nil {
		ctx.SecureJSON(http.StatusInternalServerError, err.Error())
		return
	}

	arg := db.CreateUserParams{
		Username: req.Username,
		Email:    req.Email,
		Nickname: req.Email,
		Password: hashPassword,
	}

	_, err = server.store.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.SecureJSON(http.StatusForbidden, err.Error())
				return
			}
		}
		ctx.SecureJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.SecureJSON(http.StatusOK, "Create User Successfully")
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string  `json:"token"`
	User  db.User `json:"user"`
}

func (server *Server) login(ctx *gin.Context) {

	var req LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.SecureJSON(http.StatusBadRequest, err.Error())
		return
	}
	user, err := server.store.GetUser(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.SecureJSON(http.StatusNotFound, err.Error())
			return
		}
		ctx.SecureJSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = utils.Decrypt(req.Password, user.Password)
	if err != nil {
		ctx.SecureJSON(http.StatusUnauthorized, err.Error())
		return
	}

	token, err := server.maker.CreateToken(user.Username, server.conf.Token.AccessTokenDuration)
	if err != nil {
		ctx.SecureJSON(http.StatusInternalServerError, err.Error())
		return
	}

	resp := LoginResponse{
		Token: token,
		User: db.User{
			ID:           user.ID,
			Username:     user.Username,
			Nickname:     user.Nickname,
			Email:        user.Email,
			Avatar:       user.Avatar,
			RegisterTime: user.RegisterTime,
		},
	}

	ctx.SecureJSON(http.StatusOK, resp)
}
