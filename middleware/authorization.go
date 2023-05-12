package middleware

import (
	"Blog/core/result"
	"Blog/core/token"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

const (
	authorizationHeader     = "Authorization"
	authorizationPayloadKey = "Authorization_Payload"
)

func Authorization(maker token.Maker, rdb *redis.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorization := ctx.Request.Header.Get(authorizationHeader)

		if len(authorization) == 0 {
			ctx.Abort()
			result.UnauthorizedError(ctx, "Unauthorized")
			return
		}

		_, err := maker.VerifyToken(authorization)
		if err != nil {
			ctx.Abort()
			result.UnauthorizedError(ctx, err.Error())
			return
		}

		// ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}
}
