package service

import (
	"admins/service/dto"

	"github.com/gin-gonic/gin"
)

type AdminService interface {
	// 管理员管理
	AdminLogin(ctx *gin.Context, req *dto.AdminLoginDTO) (string, error)
	CreateAdmin(ctx *gin.Context, req *dto.CreateAdminDTO) error
	UpdateAdmin(ctx *gin.Context, req *dto.UpdateAdminDTO) error
	DeleteAdmin(ctx *gin.Context, id int) error
	GetAdminByID(ctx *gin.Context, id int) (*dto.AdminDTO, error)
	ListAdmins(ctx *gin.Context, page, size int) ([]*dto.AdminDTO, int64, error)
	UpdateAdminStatus(ctx *gin.Context, id int, status int) error
}
