package main

import (
	"blacklist/config"
	"blacklist/internal/pkg/database"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"blacklist/internal/middleware"
	"blacklist/internal/pkg/redis"
	"blacklist/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化日志
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// 加载配置
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	// 初始化数据库连接
	db, err := database.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("初始化数据库失败: %v", err)
	}
	defer db.Close()

	// 优雅关闭
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		<-quit
		log.Println("正在关闭服务...")
		cancel()
	}()

	// 初始化Redis连接
	if err := redis.InitRedis(cfg); err != nil {
		log.Fatalf("Redis初始化失败: %v", err)
	}

	// 设置运行模式
	gin.SetMode(cfg.Server.Mode)

	// 创建Gin引擎
	app := gin.New()

	// 注册中间件
	app.Use(gin.Recovery())
	app.Use(middleware.Logger())
	app.Use(middleware.Cors())

	// 注册路由
	router.RegisterRoutes(app, db)

	// 启动服务器
	if err := app.Run(":" + cfg.Server.Port); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}

	<-ctx.Done()
	log.Println("服务已关闭")
}
