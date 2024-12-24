package service

import (
	"blackapp/internal/service/dto"
	"context"
)

type SystemService interface {
	// 系统监控
	GetSystemMetrics(ctx context.Context) (*dto.SystemMetrics, error)

	// 管理员管理
	AdminLogin(ctx context.Context, req *dto.AdminLoginDTO) (string, error)
	CreateAdmin(ctx context.Context, req *dto.CreateAdminDTO) error
	UpdateAdmin(ctx context.Context, req *dto.UpdateAdminDTO) error
	DeleteAdmin(ctx context.Context, id int) error
	GetAdminByID(ctx context.Context, id int) (*dto.AdminDTO, error)
	ListAdmins(ctx context.Context, page, size int) ([]*dto.AdminDTO, int64, error)
	UpdateAdminStatus(ctx context.Context, id int, status int) error

	// 日志查询
	ListLoginLogs(ctx context.Context, userType int, page, size int) ([]*dto.LoginLogDTO, int64, error)
	ListQueryLogs(ctx context.Context, merchantID int, page, size int) ([]*dto.QueryLogDTO, int64, error)
}
