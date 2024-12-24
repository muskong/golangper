package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"blackapp/internal/api/router"
	"blackapp/pkg/config"
	"blackapp/pkg/database"
	"blackapp/pkg/logger"
)

func main() {
	// 初始化配置
	if err := config.Init(); err != nil {
		log.Fatalf("初始化配置失败: %v", err)
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

	// 初始化日志
	if err := logger.Init(); err != nil {
		log.Fatalf("初始化日志失败: %v", err)
	}

	// 初始化数据库
	if err := database.Init(); err != nil {
		log.Fatalf("初始化数据库失败: %v", err)
	}

	// 数据库迁移
	if err := database.AutoMigrate(); err != nil {
		log.Fatalf("数据库迁移失败: %v", err)
	}

	// 初始化管理员账号
	if err := database.InitAdminUser(); err != nil {
		log.Fatalf("初始化管理员账号失败: %v", err)
	}

	// 初始化Redis
	if err := database.InitRedis(); err != nil {
		log.Fatalf("初始化Redis失败: %v", err)
	}

	// 初始化路由
	r := router.InitClientRouter()

	// 启动服务器
	if err := r.Run(config.GetString("app.address")); err != nil {
		log.Fatalf("启动服务器失败: %v", err)
	}
}
