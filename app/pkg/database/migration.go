package database

import (
	"blackapp/internal/domain/entity"
	"blackapp/pkg/logger"
	"go.uber.org/zap"
)

// AutoMigrate 自动迁移数据库表结构
func AutoMigrate() error {
	// 需要迁移的实体列表
	entities := []interface{}{
		&entity.Admin{},
		&entity.Merchant{},
		&entity.Blacklist{},
		&entity.LoginLog{},
		&entity.QueryLog{},
	}

	// 执行迁移
	for _, entity := range entities {
		if err := DB.AutoMigrate(entity); err != nil {
			logger.Logger.Error("数据库迁移失败", zap.Error(err))
			return err
		}
	}

	return nil
}

// InitAdminUser 初始化管理员账号
func InitAdminUser() error {
	var count int64
	if err := DB.Model(&entity.Admin{}).Count(&count).Error; err != nil {
		return err
	}

	// 如果没有管理员账号，则创建默认管理员
	if count == 0 {
		admin := &entity.Admin{
			Username: "admin",
			Password: "21232f297a57a5a743894a0e4a801fc3", // admin的MD5值
			Name:     "系统管理员",
			Status:   1,
		}
		if err := DB.Create(admin).Error; err != nil {
			logger.Logger.Error("创建默认管理员失败", zap.Error(err))
			return err
		}
		logger.Logger.Info("创建默认管理员成功")
	}

	return nil
}
