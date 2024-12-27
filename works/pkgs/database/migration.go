package database

import (
	blacklists "github.com/muskong/gopermission/works/blacklists/domain/entity"

	merchants "github.com/muskong/gopermission/works/merchants/domain/entity"

	"github.com/muskong/gopermission/works/pkgs/logger"

	admins "github.com/muskong/gopermission/works/admins/domain/entity"

	"go.uber.org/zap"
)

// AutoMigrate 自动迁移数据库表结构
func AutoMigrate() error {
	// 需要迁移的实体列表
	entities := []interface{}{
		&admins.Admin{},
		&admins.AdminRole{},
		&admins.AdminPost{},
		&admins.Department{},
		&admins.Role{},
		&admins.Post{},
		&admins.Menu{},
		&admins.RoleMenu{},
		&admins.Config{},
		&admins.Job{},
		&admins.JobLog{},
		&merchants.Merchant{},
		&blacklists.Blacklist{},
		&merchants.LoginLog{},
		&blacklists.QueryLog{},
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
	if err := DB.Model(&admins.Admin{}).Count(&count).Error; err != nil {
		return err
	}

	// 如果没有管理员账号，则创建默认管理员
	if count == 0 {
		admin := &admins.Admin{
			AdminName:     "admin",
			AdminPassword: "21232f297a57a5a743894a0e4a801fc3", // admin的MD5值
			AdminEmail:    "admin@example.com",
			AdminPhone:    "1234567890",
			AdminSex:      0,
			AdminAvatar:   "",
			AdminStatus:   1,
		}
		if err := DB.Create(admin).Error; err != nil {
			logger.Logger.Error("创建默认管理员失败", zap.Error(err))
			return err
		}
		logger.Logger.Info("创建默认管理员成功")
	}

	return nil
}
