package middleware

import (
	"blacklist/internal/service"
	"blacklist/pkg/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

// MerchantAuth 商户认证中间件
func MerchantAuth(merchantService *service.MerchantService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取Token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.Error(c, 401, "未提供认证令牌")
			c.Abort()
			return
		}

		// 解析Token
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			utils.Error(c, 401, "无效的认证令牌格式")
			c.Abort()
			return
		}

		// 验证Token
		claims, err := utils.ParseToken(parts[1])
		if err != nil {
			utils.Error(c, 401, "无效的认证令牌")
			c.Abort()
			return
		}

		// 获取商户信息
		merchant, err := merchantService.GetByID(claims.MerchantID)
		if err != nil {
			utils.Error(c, 401, "商户不存在")
			c.Abort()
			return
		}

		if merchant.Status != 1 {
			utils.Error(c, 403, "商户已被禁用")
			c.Abort()
			return
		}

		// 将商户ID存入上下文
		c.Set("merchant_id", merchant.ID)
		c.Set("merchant_name", merchant.Name)
		c.Next()
	}
}
