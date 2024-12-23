package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
)

// ResponseMiddleware 响应中间件
func ResponseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 处理请求
		c.Next()

		// 如果已经写入了响应，则不做处理
		if c.Writer.Written() {
			return
		}

		// 获取错误信息
		if len(c.Errors) > 0 {
			c.JSON(c.Writer.Status(), gin.H{
				"time":    time.Now().Format("2006-01-02 15:04:05"),
				"code":    c.Writer.Status(),
				"message": c.Errors.String(),
				"data":    nil,
			})
			return
		}
	}
}
