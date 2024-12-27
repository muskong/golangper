package router

import (
	"command-client/middleware"
	"log"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/muskong/gopermission/works/pkgs/config"
	"github.com/muskong/gopermission/works/pkgs/database"

	blacklistMapper "github.com/muskong/gopermission/works/blacklists/mapper"

	merchantMapper "github.com/muskong/gopermission/works/merchants/mapper"

	blacklistHandler "github.com/muskong/gopermission/works/blacklists/handler"

	merchantHandler "github.com/muskong/gopermission/works/merchants/handler"

	blacklistService "github.com/muskong/gopermission/works/blacklists/service/impl"

	merchantService "github.com/muskong/gopermission/works/merchants/service/impl"
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
	merchantRepo := merchantMapper.NewMerchantRepository(database.DB)
	loginLogRepo := merchantMapper.NewLoginLogRepository(database.DB)
	blacklistRepo := blacklistMapper.NewBlacklistRepository(database.DB)
	queryLogRepo := blacklistMapper.NewQueryLogRepository(database.DB)

	merchantService := merchantService.NewMerchantService(merchantRepo, loginLogRepo, jwtSecret, tokenExpire)
	blacklistService := blacklistService.NewBlacklistService(blacklistRepo, queryLogRepo)

	merchantHandler := merchantHandler.NewMerchantHandler(merchantService)
	blacklistHandler := blacklistHandler.NewBlacklistHandler(blacklistService)

	// 公开接口
	public := app.Group("/api/v1")
	{
		public.POST("/merchants/login", merchantHandler.Login)
	}

	// 需要认证的接口
	authorized := app.Group("/api/v1")
	authorized.Use(middleware.JWTAuthMerchant())
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
