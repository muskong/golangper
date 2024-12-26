package middleware

import (
	"admins/service/dto"
	"bytes"
	"io"
	"time"

	"github.com/gin-gonic/gin"
)

type responseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w responseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 记录请求体
		var requestBody []byte
		if c.Request.Body != nil {
			requestBody, _ = io.ReadAll(c.Request.Body)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
		}

		// 包装 ResponseWriter
		w := &responseWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
		c.Writer = w

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)

		// 获取当前用户
		var adminID int
		if claims, exists := c.Get("claims"); exists {
			if jwtClaims, ok := claims.(*AdminClaims); ok {
				adminID = jwtClaims.AdminID
			}
		}

		// 记录操作日志
		operationLog := &dto.OperationLogCreateDTO{
			AdminID:           adminID,
			OperationIP:       c.ClientIP(),
			OperationMethod:   c.Request.Method,
			OperationPath:     c.Request.URL.Path,
			OperationStatus:   c.Writer.Status(),
			OperationLatency:  int(latencyTime.Milliseconds()),
			OperationRequest:  string(requestBody),
			OperationResponse: w.body.String(),
		}

		// 异步保存日志
		go saveOperationLog(operationLog)
	}
}

func saveOperationLog(log *dto.OperationLogCreateDTO) {
	// TODO: 实现日志保存逻辑
}
