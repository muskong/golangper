package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"blacklist/internal/api"
	"blacklist/internal/middleware"
	"blacklist/internal/pkg/database"
	"blacklist/internal/pkg/redis"
	"blacklist/pkg/config"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化日志
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// 加载配置
	err := config.Init()
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	// 优雅关闭
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		sig := <-quit
		switch sig {
		case syscall.SIGINT:
			log.Println("接收到SIGINT信号，正在关闭服务...")
		case syscall.SIGTERM:
			log.Println("接收到SIGTERM信号，正在关闭服务...")
		default:
			log.Printf("接收到%v信号，正在关闭服务...", sig)
		}
		cancel()
		<-ctx.Done()
		os.Exit(0)
		log.Println("服务已关闭")
	}()

	// 初始化数据库连接
	db, err := database.NewPostgresDB(&config.Cfg.Database)
	if err != nil {
		log.Fatalf("初始化数据库失败: %v", err)
	}
	defer db.Close()

	// 初始化Redis连接
	if err := redis.InitRedis(&config.Cfg.Redis); err != nil {
		log.Fatalf("Redis初始化失败: %v", err)
	}

	// 设置运行模式
	gin.SetMode(config.Cfg.Server.Mode)

	// 创建Gin引擎
	app := gin.New()

	// 注册中间件
	app.Use(gin.Recovery())
	app.Use(middleware.Logger())
	app.Use(middleware.Cors())

	// 注册路由
	api.RegisterRoutes(app, db)

	// 启动服务器
	if err := app.Run(":" + config.Cfg.Server.Port); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}

}
