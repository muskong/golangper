package router

import (
	"command-server/middleware"
	"log"
	"time"

	"github.com/gin-gonic/gin"

	"pkgs/config"
	"pkgs/database"

	adminMapper "admins/mapper"
	blacklistMapper "blacklists/mapper"
	merchantMapper "merchants/mapper"

	adminHandler "admins/handler"
	blacklistHandler "blacklists/handler"
	merchantHandler "merchants/handler"
	systemHandler "systems/handler"

	adminService "admins/service/impl"
	blacklistService "blacklists/service/impl"
	merchantService "merchants/service/impl"
	systemService "systems/service/impl"
)

func InitServerRouter() *gin.Engine {
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
	adminRepo, logRepo, _ := adminMapper.NewRepository(database.DB)
	merchantRepo, loginLogRepo := merchantMapper.NewRepository(database.DB)
	queryLogRepo, blacklistRepo := blacklistMapper.NewRepository(database.DB)

	adminService := adminService.NewAdminService(adminRepo, logRepo, jwtSecret, tokenExpire)
	merchantService := merchantService.NewMerchantService(merchantRepo, loginLogRepo, jwtSecret, tokenExpire)
	blacklistService := blacklistService.NewBlacklistService(blacklistRepo, queryLogRepo)
	systemService := systemService.NewSystemService(database.RDB, database.DB)

	adminHandler := adminHandler.NewAdminHandler(adminService)
	merchantHandler := merchantHandler.NewMerchantHandler(merchantService)
	blacklistHandler := blacklistHandler.NewBlacklistHandler(blacklistService)
	systemHandler := systemHandler.NewSystemHandler(systemService)

	// 公开接口
	public := app.Group("/api/v1")
	{
		public.POST("/admins/login", adminHandler.AdminLogin)
	}

	// 需要认证的接口
	authorized := app.Group("/api/v1")
	authorized.Use(middleware.JWTAuthAdmin())
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
			merchants.GET("/logs", merchantHandler.ListLoginLogs)
		}

		// 黑名单管理
		blacklists := authorized.Group("/blacklists")
		{
			blacklists.POST("", blacklistHandler.Create)
			blacklists.PUT("/:id", blacklistHandler.Update)
			blacklists.DELETE("/:id", blacklistHandler.Delete)
			blacklists.GET("/:id", blacklistHandler.GetByID)
			blacklists.GET("", blacklistHandler.List)
			blacklists.PUT("/:id/status", blacklistHandler.UpdateStatus)
			blacklists.GET("/logs", blacklistHandler.ListQueryLogs)
		}

		// 系统监控
		authorized.GET("/system/metrics", systemHandler.GetSystemMetrics)

		// 管理员管理
		admins := authorized.Group("/admins")
		{
			admins.POST("", adminHandler.CreateAdmin)
			admins.PUT("/:id", adminHandler.UpdateAdmin)
			admins.GET("", adminHandler.ListAdmins)
			admins.PUT("/:id/status", adminHandler.UpdateAdminStatus)
		}

		// 角色管理
		role := authorized.Group("/role")
		{
			role.POST("", adminHandler.CreateRole)
			role.PUT("", adminHandler.UpdateRole)
			role.GET("/list", adminHandler.ListRoles)
			role.DELETE("/:id", adminHandler.DeleteRole)
		}

		// 部门管理
		dept := authorized.Group("/department")
		{
			dept.POST("", adminHandler.CreateDepartment)
			dept.PUT("", adminHandler.UpdateDepartment)
			dept.GET("/tree", adminHandler.GetDepartmentTree)
			dept.DELETE("/:id", adminHandler.DeleteDepartment)
		}
	}

	return app
}
