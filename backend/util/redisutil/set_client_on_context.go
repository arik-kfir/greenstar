package redisutil

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rueian/rueidis"
)

const key = "$$$redisKey$$$"

func GetRedis(ctx context.Context) rueidis.Client {
	if v := ctx.Value(key); v != nil {
		if r, ok := v.(rueidis.Client); ok {
			return r
		} else {
			panic(fmt.Sprintf("non-Redis client found in request context under key '%s'", key))
		}
	} else {
		panic(fmt.Sprintf("no Redis client found in request context under key '%s'", key))
	}
}

func CreateSetRedisMiddleware(r rueidis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(key, r)
		c.Next()
	}
}
