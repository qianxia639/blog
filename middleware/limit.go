package middleware

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis_rate/v10"
)

func Limit(limiter *redis_rate.Limiter) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		res, err := limiter.Allow(ctx, "192.168.172.136:9010", redis_rate.PerSecond(5))
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		ctx.Writer.Header().Set("RateLimit-Remaining", strconv.Itoa(res.Remaining))

		if res.Allowed == 0 {
			seconds := int(res.RetryAfter / time.Second)
			ctx.Writer.Header().Set("RateLimit-RetryAfter", strconv.Itoa(seconds))
			ctx.AbortWithStatus(http.StatusTooManyRequests)
			return
		}

		ctx.Next()
	}
}
