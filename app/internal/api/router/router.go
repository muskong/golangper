package router

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"

	"blackapp/internal/api/handler"
	"blackapp/internal/api/middleware"
	"blackapp/internal/infrastructure/persistence"
	"blackapp/internal/service/impl"
	"blackapp/pkg/config"
	"blackapp/pkg/database"
)

func InitRouter() *gin.Engine {
	gin.SetMode(config.GetString("app.mode"))
	app := gin.New()
	// 注册中间件
	app.Use(gin.Logger())
	app.Use(gin.Recovery())
	// app.Use(middleware.Logger())
	// app.Use(middleware.Cors())

	jwtSecret := config.GetString("jwt.secret")
	tokenExpire, err := time.ParseDuration(config.GetString("jwt.token_expire"))
	if err != nil {
		// 处理错误
		log.Fatalf("解析token过期时间失败: %v", err)
	}

	// 初始化系统服务依赖
	adminRepo := persistence.NewAdminRepository()
	loginLogRepo := persistence.NewLoginLogRepository()
	queryLogRepo := persistence.NewQueryLogRepository()

	// 初始化依赖
	merchantRepo := persistence.NewMerchantRepository()
	blacklistRepo := persistence.NewBlacklistRepository()
	systemService := impl.NewSystemService(
		adminRepo,
		loginLogRepo,
		queryLogRepo,
		jwtSecret,
		tokenExpire,
		database.RDB,
		database.DB,
	)

	merchantService := impl.NewMerchantService(merchantRepo, loginLogRepo, jwtSecret, tokenExpire)
	blacklistService := impl.NewBlacklistService(blacklistRepo, queryLogRepo)

	merchantHandler := handler.NewMerchantHandler(merchantService)
	blacklistHandler := handler.NewBlacklistHandler(blacklistService)
	systemHandler := handler.NewSystemHandler(systemService)

	// 公开接口
	public := app.Group("/api/v1")
	{
		public.POST("/admins/login", systemHandler.AdminLogin)
	}

	// 需要认证的接口
	authorized := app.Group("/api/v1")
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
		}

		// 系统监控
		authorized.GET("/system/metrics", systemHandler.GetSystemMetrics)

		// 管理员管理
		admins := authorized.Group("/admins")
		{
			admins.POST("", systemHandler.CreateAdmin)
			admins.PUT("/:id", systemHandler.UpdateAdmin)
			admins.GET("", systemHandler.ListAdmins)
			admins.PUT("/:id/status", systemHandler.UpdateAdminStatus)
		}

		// 日志管理
		logs := authorized.Group("/logs")
		{
			logs.GET("/login", systemHandler.ListLoginLogs)
			logs.GET("/query", systemHandler.ListQueryLogs)
		}
	}

	return app
}
