package router

import (
	v1 "blacklist/api/v1"
	"blacklist/internal/pkg/database"
	"blacklist/internal/repository"
	"blacklist/internal/service"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, db *database.PostgresDB) {
	// 初始化依赖
	blacklistRepo := repository.NewBlacklistRepository(db)
	blacklistService := service.NewBlacklistService(blacklistRepo)
	blacklistHandler := v1.NewBlacklistHandler(blacklistService)

	// API v1
	v1Group := r.Group("/api/v1")
	{
		// 黑名单用户管理
		blacklist := v1Group.Group("/blacklist")
		{
			blacklist.GET("", blacklistHandler.ListBlacklistUsers)
			blacklist.POST("", blacklistHandler.CreateBlacklistUser)
			blacklist.GET("/:id", blacklistHandler.GetBlacklistUser)
			blacklist.PUT("/:id", blacklistHandler.UpdateBlacklistUser)
			blacklist.DELETE("/:id", blacklistHandler.DeleteBlacklistUser)
			blacklist.GET("/check", blacklistHandler.CheckPhoneExists)
			blacklist.GET("/phone", blacklistHandler.GetByPhone)
			blacklist.GET("/exists", blacklistHandler.CheckExists)
			blacklist.GET("/id_card", blacklistHandler.GetByIDCard)
			blacklist.GET("/name", blacklistHandler.GetByName)
		}
	}
}
