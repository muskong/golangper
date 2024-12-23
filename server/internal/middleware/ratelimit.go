package middleware

import (
	"context"
	"fmt"
	"time"

	"blacklist/internal/pkg/redis"
	"blacklist/pkg/utils"

	"github.com/gin-gonic/gin"
)

// RateLimit 请求频率限制中间件
func RateLimit(limit int, duration time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取商户ID
		merchantID, exists := c.Get("merchant_id")
		if !exists {
			utils.Error(c, 401, "未授权的访问")
			c.Abort()
			return
		}

		// 构造Redis key
		key := fmt.Sprintf("ratelimit:merchant:%d", merchantID)

		// 获取当前计数
		count, err := redis.RDB.Get(context.Background(), key).Int()
		if err == nil && count >= limit {
			utils.Error(c, 429, "请求过于频繁，请稍后再试")
			c.Abort()
			return
		}

		// 增加计数
		if err := redis.RDB.Incr(context.Background(), key).Err(); err != nil {
			utils.Error(c, 500, "服务器内部错误")
			c.Abort()
			return
		}

		// 设置过期时间
		if count == 0 {
			redis.RDB.Expire(context.Background(), key, duration)
		}

		c.Next()
	}
}
