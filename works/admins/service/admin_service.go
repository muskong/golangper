package service

import (
	"admins/service/dto"

	"github.com/gin-gonic/gin"
)

type AdminService interface {
	// 管理员相关
	AdminLogin(ctx *gin.Context, req *dto.AdminLoginDTO) (string, error)
	CreateAdmin(ctx *gin.Context, req *dto.CreateAdminDTO) error
	UpdateAdmin(ctx *gin.Context, req *dto.UpdateAdminDTO) error
	DeleteAdmin(ctx *gin.Context, adminID int) error
	GetAdminInfo(ctx *gin.Context, adminID int) (*dto.AdminDTO, error)
	ListAdmins(ctx *gin.Context, page, size int) ([]*dto.AdminDTO, int64, error)

	// 角色相关
	CreateRole(ctx *gin.Context, req *dto.CreateRoleDTO) error
	UpdateRole(ctx *gin.Context, req *dto.UpdateRoleDTO) error
	DeleteRole(ctx *gin.Context, roleID int) error
	ListRoles(ctx *gin.Context, page, size int) ([]*dto.RoleDTO, int64, error)

	// 部门相关
	CreateDepartment(ctx *gin.Context, req *dto.CreateDepartmentDTO) error
	UpdateDepartment(ctx *gin.Context, req *dto.UpdateDepartmentDTO) error
	DeleteDepartment(ctx *gin.Context, deptID int) error
	GetDepartmentTree(ctx *gin.Context) ([]*dto.DepartmentTreeDTO, error)
}
