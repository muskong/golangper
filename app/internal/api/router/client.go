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
)

func InitClientRouter() *gin.Engine {
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

	// 初始化依赖
	merchantRepo := persistence.NewMerchantRepository()
	loginLogRepo := persistence.NewLoginLogRepository()
	blacklistRepo := persistence.NewBlacklistRepository()
	queryLogRepo := persistence.NewQueryLogRepository()

	merchantService := impl.NewMerchantService(merchantRepo, loginLogRepo, jwtSecret, tokenExpire)
	blacklistService := impl.NewBlacklistService(blacklistRepo, queryLogRepo)

	merchantHandler := handler.NewMerchantHandler(merchantService)
	blacklistHandler := handler.NewBlacklistHandler(blacklistService)

	// 公开接口
	public := app.Group("/api/v1")
	{
		public.POST("/merchants/login", merchantHandler.Login)
	}

	// 需要认证的接口
	authorized := app.Group("/api/v1")
	authorized.Use(middleware.JWTAuth())
	{
		// 黑名单管理
		blacklists := authorized.Group("/blacklists")
		blacklists.Use(middleware.RateLimit())
		{
			blacklists.POST("/check", blacklistHandler.Check)
		}
	}

	return app
}
