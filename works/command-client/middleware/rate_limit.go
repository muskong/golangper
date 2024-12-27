package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"

	"github.com/muskong/gopermission/works/pkgs/response"
)

var limiter = rate.NewLimiter(rate.Every(time.Second), 100)

func RateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !limiter.Allow() {
			response.Error(c, 429, "请求过于频繁，请稍后重试")
			c.Abort()
			return
		}
		c.Next()
	}
}
