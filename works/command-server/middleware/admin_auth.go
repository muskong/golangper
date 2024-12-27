package middleware

import (
	"strings"
	"time"

	"github.com/muskong/gopermission/works/pkgs/config"
	"github.com/muskong/gopermission/works/pkgs/response"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuthAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Unauthorized(c)
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			response.Unauthorized(c)
			return
		}

		// 验证token
		claims, err := AdminParseToken(parts[1])
		if err != nil {
			response.Unauthorized(c)
			return
		}

		// 将用户信息存入上下文
		c.Set("adminId", claims.AdminID)
		c.Set("adminName", claims.AdminName)
		c.Set("claims", claims)

		c.Next()
	}
}

// CheckPermission 检查权限
func CheckPermission(permission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		adminID := c.GetInt("adminId")
		if adminID == 0 {
			response.Unauthorized(c)
			return
		}

		// TODO: 实现权限检查逻辑
		// 1. 获取用户角色
		// 2. 获取角色权限
		// 3. 检查是否包含所需权限

		c.Next()
	}
}

// CheckRole 检查角色
func CheckRole(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		adminID := c.GetInt("adminId")
		if adminID == 0 {
			response.Unauthorized(c)
			return
		}

		// TODO: 实现角色检查逻辑
		// 1. 获取用户角色
		// 2. 检查是否包含所需角色

		c.Next()
	}
}

type AdminClaims struct {
	AdminID   int    `json:"adminId"`
	AdminName string `json:"adminName"`
	jwt.RegisteredClaims
}

func AdminGenerateToken(adminID int, adminName string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)

	claims := AdminClaims{
		AdminID:   adminID,
		AdminName: adminName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			IssuedAt:  jwt.NewNumericDate(nowTime),
			NotBefore: jwt.NewNumericDate(nowTime),
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(config.GetString("jwt.secret")))

	return token, err
}

func AdminParseToken(token string) (*AdminClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &AdminClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetString("jwt.secret")), nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*AdminClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
