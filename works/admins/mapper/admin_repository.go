package mapper

import (
	"admins/domain/entity"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *adminRepository {
	return &adminRepository{db: db}
}

// Admin相关实现
func (r *adminRepository) FindByAdminName(ctx *gin.Context, adminName string) (*entity.Admin, error) {
	var admin entity.Admin
	err := r.db.Where("admin_name = ?", adminName).First(&admin).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

func (r *adminRepository) Create(ctx *gin.Context, admin *entity.Admin) error {
	return r.db.Create(admin).Error
}

func (r *adminRepository) Update(ctx *gin.Context, admin *entity.Admin) error {
	return r.db.Save(admin).Error
}

func (r *adminRepository) Delete(ctx *gin.Context, adminID int) error {
	return r.db.Delete(&entity.Admin{}, adminID).Error
}

func (r *adminRepository) FindByID(ctx *gin.Context, adminID int) (*entity.Admin, error) {
	var admin entity.Admin
	err := r.db.Preload("Roles").Preload("Posts").First(&admin, adminID).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

func (r *adminRepository) List(ctx *gin.Context, offset, limit int) ([]*entity.Admin, int64, error) {
	var admins []*entity.Admin
	var total int64

	err := r.db.Model(&entity.Admin{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.Preload("Roles").Preload("Posts").
		Offset(offset).Limit(limit).Find(&admins).Error
	if err != nil {
		return nil, 0, err
	}

	return admins, total, nil
}

// Role相关实现
func (r *adminRepository) CreateRole(ctx *gin.Context, role *entity.Role) error {
	return r.db.Create(role).Error
}

func (r *adminRepository) UpdateRole(ctx *gin.Context, role *entity.Role) error {
	return r.db.Save(role).Error
}

func (r *adminRepository) DeleteRole(ctx *gin.Context, roleID int) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// 删除角色菜单关联
		if err := tx.Delete(&entity.RoleMenu{}, "role_id = ?", roleID).Error; err != nil {
			return err
		}
		// 删除角色部门关联
		if err := tx.Delete(&entity.RoleDepartment{}, "role_id = ?", roleID).Error; err != nil {
			return err
		}
		// 删除角色
		return tx.Delete(&entity.Role{}, roleID).Error
	})
}

func (r *adminRepository) FindRoleByID(ctx *gin.Context, roleID int) (*entity.Role, error) {
	var role entity.Role
	err := r.db.First(&role, roleID).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *adminRepository) ListRoles(ctx *gin.Context, offset, limit int) ([]*entity.Role, int64, error) {
	var roles []*entity.Role
	var total int64

	err := r.db.Model(&entity.Role{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.Offset(offset).Limit(limit).Find(&roles).Error
	if err != nil {
		return nil, 0, err
	}

	return roles, total, nil
}

// Department相关实现
func (r *adminRepository) CreateDepartment(ctx *gin.Context, dept *entity.Department) error {
	return r.db.Create(dept).Error
}

func (r *adminRepository) UpdateDepartment(ctx *gin.Context, dept *entity.Department) error {
	return r.db.Save(dept).Error
}

func (r *adminRepository) DeleteDepartment(ctx *gin.Context, deptID int) error {
	return r.db.Delete(&entity.Department{}, deptID).Error
}

func (r *adminRepository) FindDepartmentByID(ctx *gin.Context, deptID int) (*entity.Department, error) {
	var dept entity.Department
	err := r.db.First(&dept, deptID).Error
	if err != nil {
		return nil, err
	}
	return &dept, nil
}

func (r *adminRepository) ListAllDepartments(ctx *gin.Context) ([]*entity.Department, error) {
	var depts []*entity.Department
	err := r.db.Order("department_sort").Find(&depts).Error
	return depts, err
}

// 关联关系实现
func (r *adminRepository) UpdateAdminRoles(ctx *gin.Context, adminID int, roleIDs []int) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// 删除旧的关联
		if err := tx.Delete(&entity.AdminRole{}, "admin_id = ?", adminID).Error; err != nil {
			return err
		}

		// 创建新的关联
		for _, roleID := range roleIDs {
			if err := tx.Create(&entity.AdminRole{AdminID: adminID, RoleID: roleID}).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *adminRepository) UpdateAdminPosts(ctx *gin.Context, adminID int, postIDs []int) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// 删除旧的关联
		if err := tx.Delete(&entity.AdminPost{}, "admin_id = ?", adminID).Error; err != nil {
			return err
		}

		// 创建新的关联
		for _, postID := range postIDs {
			if err := tx.Create(&entity.AdminPost{AdminID: adminID, PostID: postID}).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *adminRepository) UpdateRoleMenus(ctx *gin.Context, roleID int, menuIDs []int) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// 删除旧的关联
		if err := tx.Delete(&entity.RoleMenu{}, "role_id = ?", roleID).Error; err != nil {
			return err
		}

		// 创建新的关联
		for _, menuID := range menuIDs {
			if err := tx.Create(&entity.RoleMenu{RoleID: roleID, MenuID: menuID}).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *adminRepository) UpdateRoleDepartments(ctx *gin.Context, roleID int, deptIDs []int) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// 删除旧的关联
		if err := tx.Delete(&entity.RoleDepartment{}, "role_id = ?", roleID).Error; err != nil {
			return err
		}

		// 创建新的关联
		for _, deptID := range deptIDs {
			if err := tx.Create(&entity.RoleDepartment{RoleID: roleID, DepartmentID: deptID}).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

// 其他方法实现...
