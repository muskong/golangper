package router

import (
	"time"

	"github.com/gin-gonic/gin"

	"blackapp/internal/api/handler"
	"blackapp/internal/api/middleware"
	"blackapp/internal/infrastructure/persistence"
	"blackapp/internal/service/impl"
	"blackapp/pkg/config"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	jwtSecret := config.GetString("jwt.secret")
	tokenExpire, err := time.ParseDuration(config.GetString("jwt.token_expire"))
	if err != nil {
		// 处理错误
	}

	// 初始化依赖
	merchantRepo := persistence.NewMerchantRepository()
	blacklistRepo := persistence.NewBlacklistRepository()

	merchantService := impl.NewMerchantService(merchantRepo, jwtSecret, tokenExpire)
	blacklistService := impl.NewBlacklistService(blacklistRepo)

	merchantHandler := handler.NewMerchantHandler(merchantService)
	blacklistHandler := handler.NewBlacklistHandler(blacklistService)

	// 公开接口
	public := r.Group("/api/v1")
	{
		public.POST("/merchants/login", merchantHandler.Login)
	}

	// 需要认证的接口
	authorized := r.Group("/api/v1")
	authorized.Use(middleware.JWTAuth())
	{
		// 商户管理
		merchants := authorized.Group("/merchants")
		{
			merchants.POST("", merchantHandler.Create)
			merchants.PUT("/:id", merchantHandler.Update)
			merchants.DELETE("/:id", merchantHandler.Delete)
			merchants.GET("/:id", merchantHandler.GetByID)
			merchants.GET("", merchantHandler.List)
			merchants.PUT("/:id/status", merchantHandler.UpdateStatus)
		}

		// 黑名单管理
		blacklists := authorized.Group("/blacklists")
		blacklists.Use(middleware.RateLimit())
		{
			blacklists.POST("", blacklistHandler.Create)
			blacklists.PUT("/:id", blacklistHandler.Update)
			blacklists.DELETE("/:id", blacklistHandler.Delete)
			blacklists.GET("/:id", blacklistHandler.GetByID)
			blacklists.GET("", blacklistHandler.List)
			blacklists.PUT("/:id/status", blacklistHandler.UpdateStatus)
			blacklists.POST("/check", blacklistHandler.Check)
		}
	}

	return r
}
