package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetCurrentMerchantID 获取当前商户ID
func GetCurrentMerchantID(c *gin.Context) uint {
	id, exists := c.Get("merchant_id")
	if !exists {
		return 0
	}
	if merchantID, ok := id.(uint); ok {
		return merchantID
	}
	return 0
}

// GetPage 获取页码
func GetPage(c *gin.Context) int {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	if page < 1 {
		page = 1
	}
	return page
}

// GetPageSize 获取每页数量
func GetPageSize(c *gin.Context) int {
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	if pageSize < 1 {
		pageSize = 10
	}
	return pageSize
}
