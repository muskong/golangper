package service

import (
	"blackapp/internal/service/dto"

	"github.com/gin-gonic/gin"
)

type SystemService interface {
	// 系统监控
	GetSystemMetrics(ctx *gin.Context) (*dto.SystemMetrics, error)

	// 管理员管理
	AdminLogin(ctx *gin.Context, req *dto.AdminLoginDTO) (string, error)
	CreateAdmin(ctx *gin.Context, req *dto.CreateAdminDTO) error
	UpdateAdmin(ctx *gin.Context, req *dto.UpdateAdminDTO) error
	DeleteAdmin(ctx *gin.Context, id int) error
	GetAdminByID(ctx *gin.Context, id int) (*dto.AdminDTO, error)
	ListAdmins(ctx *gin.Context, page, size int) ([]*dto.AdminDTO, int64, error)
	UpdateAdminStatus(ctx *gin.Context, id int, status int) error

	// 日志查询
	ListLoginLogs(ctx *gin.Context, userType int, page, size int) ([]*dto.LoginLogDTO, int64, error)
	ListQueryLogs(ctx *gin.Context, merchantID int, page, size int) ([]*dto.QueryLogDTO, int64, error)
}
