package api

import (
	v1 "blacklist/internal/api/v1"
	"blacklist/internal/middleware"
	"blacklist/internal/pkg/database"
	"blacklist/internal/repository"
	"blacklist/internal/service"
	"time"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, db *database.PostgresDB) {
	// 注册全局中间件
	r.Use(middleware.ResponseMiddleware())

	// 初始化依赖
	merchantRepo := repository.NewMerchantRepository(db)
	merchantService := service.NewMerchantService(merchantRepo)
	merchantHandler := v1.NewMerchantHandler(merchantService)

	blacklistRepo := repository.NewBlacklistRepository(db)
	blacklistService := service.NewBlacklistService(blacklistRepo)
	blacklistHandler := v1.NewBlacklistHandler(blacklistService)

	adminGroup := r.Group("/admin")
	{
		// 商户管理
		merchant := adminGroup.Group("/merchants")
		{
			// 后台管理接口
			merchant.POST("", merchantHandler.CreateMerchant)
			merchant.GET("", merchantHandler.ListMerchants)
			merchant.GET("/:id", merchantHandler.GetMerchant)
			merchant.PUT("/:id", merchantHandler.UpdateMerchant)
			merchant.DELETE("/:id", merchantHandler.DeleteMerchant)
			merchant.PUT("/:id/status", merchantHandler.UpdateMerchantStatus)
			merchant.POST("/:id/regenerate", merchantHandler.RegenerateAPICredentials)
		}
	}

	// API v1
	v1Group := r.Group("/api/v1")
	{
		// 商户登录
		v1Group.POST("/login", merchantHandler.Login)

		// 商户API接口（需要认证）
		blacklist := v1Group.Group("/blacklist")
		blacklist.Use(middleware.MerchantAuth(merchantService))
		blacklist.Use(middleware.RateLimit(100, time.Minute)) // 限制每分钟100次请求
		{
			// 黑名单相关接口
			blacklist.GET("", blacklistHandler.ListBlacklistUsers)
			blacklist.POST("", blacklistHandler.CreateBlacklistUser)
			blacklist.GET("/:id", blacklistHandler.GetBlacklistUser)
			blacklist.PUT("/:id", blacklistHandler.UpdateBlacklistUser)
			blacklist.GET("/check", blacklistHandler.CheckPhoneExists)
			blacklist.GET("/exists", blacklistHandler.CheckExists)
		}
	}
}
