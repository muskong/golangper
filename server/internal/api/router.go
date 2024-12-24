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
	merchantLoginLogRepo := repository.NewMerchantLoginLogRepository(db)
	merchantService := service.NewMerchantService(merchantRepo, merchantLoginLogRepo)
	merchantHandler := v1.NewMerchantHandler(merchantService)

	blacklistRepo := repository.NewBlacklistRepository(db)
	blacklistQueryLogRepo := repository.NewBlacklistQueryLogRepository(db)
	blacklistService := service.NewBlacklistService(blacklistRepo, blacklistQueryLogRepo)
	blacklistHandler := v1.NewBlacklistHandler(blacklistService)

	adminRepo := repository.NewAdminRepository(db)
	adminService := service.NewAdminService(adminRepo)
	adminHandler := v1.NewAdminHandler(adminService)

	adminGroup := r.Group("/admin")
	{
		adminGroup.Use(middleware.MerchantAuth(merchantService))
		// 管理员相关接口
		admin := adminGroup.Group("/admin")
		{
			admin.GET("", adminHandler.ListAdmins)
			admin.GET("/:id", adminHandler.GetAdmin)
			admin.POST("", adminHandler.CreateAdmin)
			admin.PUT("/:id", adminHandler.UpdateAdmin)
			admin.DELETE("/:id", adminHandler.DeleteAdmin)
		}
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
			merchant.GET("/login-logs", merchantHandler.GetLoginLogs)
		}

		// 黑名单管理
		blacklist := adminGroup.Group("/blacklist")
		{
			blacklist.GET("", blacklistHandler.ListBlacklistUsers)
			blacklist.GET("/:id", blacklistHandler.GetBlacklistUser)
			blacklist.PUT("/:id", blacklistHandler.UpdateBlacklistUser)
			// 查询日志相关接口
			blacklist.GET("/query-logs", blacklistHandler.GetAllQueryLogs)                   // 获取所有查询日志
			blacklist.GET("/query-logs/merchant/:id", blacklistHandler.GetMerchantQueryLogs) // 获取指定商户的查询日志
			blacklist.GET("/query-logs/phone", blacklistHandler.GetPhoneQueryLogs)           // 获取指定手机号的查询日志
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
			blacklist.POST("", blacklistHandler.CreateBlacklistUser)
			blacklist.GET("/check", blacklistHandler.CheckPhoneExists)
			blacklist.GET("/exists", blacklistHandler.CheckExists)
		}
	}
}