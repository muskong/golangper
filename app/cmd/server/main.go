package main

import (
	"log"

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

	// 初始化日志
	if err := logger.Init(); err != nil {
		log.Fatalf("初始化日志失败: %v", err)
	}

	// 初始化数据库
	if err := database.Init(); err != nil {
		log.Fatalf("初始化数据库失败: %v", err)
	}
	// 初始化路由
	r := router.InitRouter()

	// 启动服务器
	if err := r.Run(config.GetString("server.port")); err != nil {
		log.Fatalf("启动服务器失败: %v", err)
	}
}
