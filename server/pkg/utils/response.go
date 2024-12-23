package utils

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Response 统一响应结构
type Response struct {
	Time    string      `json:"time"`    // 响应时间
	Code    int         `json:"code"`    // 响应码
	Message string      `json:"message"` // 响应信息
	Data    interface{} `json:"data"`    // 响应数据
}

// ResponseWithPagination 带分页的响应结构
type ResponseWithPagination struct {
	Time    string      `json:"time"`    // 响应时间
	Code    int         `json:"code"`    // 响应码
	Message string      `json:"message"` // 响应信息
	Data    interface{} `json:"data"`    // 响应数据
	Total   int64       `json:"total"`   // 总记录数
	Page    int         `json:"page"`    // 当前页码
	Size    int         `json:"size"`    // 每页大小
}

// Success 成功响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Time:    time.Now().Format("2006-01-02 15:04:05"),
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
	})
}

// SuccessWithPagination 带分页的成功响应
func SuccessWithPagination(c *gin.Context, data interface{}, total int64, page, size int) {
	c.JSON(http.StatusOK, ResponseWithPagination{
		Time:    time.Now().Format("2006-01-02 15:04:05"),
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
		Total:   total,
		Page:    page,
		Size:    size,
	})
}

// Error 错误响应
func Error(c *gin.Context, code int, message string) {
	c.JSON(code, Response{
		Time:    time.Now().Format("2006-01-02 15:04:05"),
		Code:    code,
		Message: message,
		Data:    nil,
	})
}
