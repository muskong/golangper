package router

import (
	"github.com/gin-gonic/gin"

	"blackapp/internal/api/handler"
	"blackapp/internal/api/middleware"
	"blackapp/internal/infrastructure/persistence"
	"blackapp/internal/service/impl"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	// 初始化依赖
	merchantRepo := persistence.NewMerchantRepository()
	blackappRepo := persistence.NewblackappRepository()

	merchantService := impl.NewMerchantService(merchantRepo)
	blackappService := impl.NewblackappService(blackappRepo)

	merchantHandler := handler.NewMerchantHandler(merchantService)
	blackappHandler := handler.NewblackappHandler(blackappService)

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
		blackapps := authorized.Group("/blackapps")
		blackapps.Use(middleware.RateLimit())
		{
			blackapps.POST("", blackappHandler.Create)
			blackapps.PUT("/:id", blackappHandler.Update)
			blackapps.DELETE("/:id", blackappHandler.Delete)
			blackapps.GET("/:id", blackappHandler.GetByID)
			blackapps.GET("", blackappHandler.List)
			blackapps.PUT("/:id/status", blackappHandler.UpdateStatus)
			blackapps.POST("/check", blackappHandler.Check)
		}
	}

	return r
}
