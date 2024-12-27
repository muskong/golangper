package middleware

import (
	"errors"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"github.com/muskong/gopermission/works/pkgs/config"
	"github.com/muskong/gopermission/works/pkgs/response"
)

func JWTAuthMerchant() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Unauthorized(c)
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			response.Unauthorized(c)
			c.Abort()
			return
		}

		// 验证token
		claims, err := MerchantParseToken(parts[1])
		if err != nil {
			response.Unauthorized(c)
			c.Abort()
			return
		}

		// 将用户信息存储到上下文
		c.Set("merchantID", claims.MerchantID)
		c.Next()
	}
}

type MerchantClaims struct {
	MerchantID int `json:"merchantID"`
	jwt.RegisteredClaims
}

func MerchantGenerateToken(merchantID int) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)

	claims := MerchantClaims{
		MerchantID: merchantID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			IssuedAt:  jwt.NewNumericDate(nowTime),
			NotBefore: jwt.NewNumericDate(nowTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.GetString("jwt.secret")))
}

func MerchantParseToken(tokenString string) (*MerchantClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MerchantClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetString("jwt.secret")), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*MerchantClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
