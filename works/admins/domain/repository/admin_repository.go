package repository

import (
	"admins/domain/entity"

	"github.com/gin-gonic/gin"
)

type AdminRepository interface {
	// Admin相关
	FindByAdminName(c *gin.Context, adminName string) (*entity.Admin, error)
	Create(c *gin.Context, admin *entity.Admin) error
	Update(c *gin.Context, admin *entity.Admin) error
	Delete(c *gin.Context, adminID int) error
	FindByID(c *gin.Context, adminID int) (*entity.Admin, error)
	List(c *gin.Context, offset, limit int) ([]*entity.Admin, int64, error)

	// Role相关
	CreateRole(c *gin.Context, role *entity.Role) error
	UpdateRole(c *gin.Context, role *entity.Role) error
	DeleteRole(c *gin.Context, roleID int) error
	FindRoleByID(c *gin.Context, roleID int) (*entity.Role, error)
	ListRoles(c *gin.Context, offset, limit int) ([]*entity.Role, int64, error)

	// Department相关
	CreateDepartment(c *gin.Context, dept *entity.Department) error
	UpdateDepartment(c *gin.Context, dept *entity.Department) error
	DeleteDepartment(c *gin.Context, deptID int) error
	FindDepartmentByID(c *gin.Context, deptID int) (*entity.Department, error)
	ListAllDepartments(c *gin.Context) ([]*entity.Department, error)

	// 关联关系
	UpdateAdminRoles(c *gin.Context, adminID int, roleIDs []int) error
	UpdateAdminPosts(c *gin.Context, adminID int, postIDs []int) error
	UpdateRoleMenus(c *gin.Context, roleID int, menuIDs []int) error
	UpdateRoleDepartments(c *gin.Context, roleID int, deptIDs []int) error
}
