package rest

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func NewRateLimiterMiddleware(redisClient *redis.Client, key string, limit int, slidingWindow time.Duration) gin.HandlerFunc {

	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprint("error init redis", err.Error()))
	}

	return func(c *gin.Context) {
		now := time.Now().UnixNano()
		userCntKey := fmt.Sprint(c.ClientIP(), ":", key)

		redisClient.ZRemRangeByScore(ctx, userCntKey,
			"0",
			fmt.Sprint(now-(slidingWindow.Nanoseconds()))).Result()

		reqs, _ := redisClient.ZRange(ctx, userCntKey, 0, -1).Result()

		if len(reqs) >= limit {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"status":  http.StatusTooManyRequests,
				"message": "too many request",
			})
			return
		}

		c.Next()
		redisClient.ZAddNX(ctx, userCntKey, &redis.Z{Score: float64(now), Member: float64(now)})
		redisClient.Expire(ctx, userCntKey, slidingWindow)
	}

}
