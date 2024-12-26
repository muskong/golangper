package impl

import (
	"admins/domain/constants"
	"admins/domain/entity"
	"admins/domain/repository"
	"admins/service/dto"
	"command-server/middleware"
	"pkgs/logger"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type adminService struct {
	adminRepo   repository.AdminRepository
	logRepo     repository.LogRepository
	jwtSecret   string
	tokenExpire time.Duration
}

func NewAdminService(
	adminRepo repository.AdminRepository,
	logRepo repository.LogRepository,
	jwtSecret string,
	tokenExpire time.Duration,
) *adminService {
	return &adminService{
		adminRepo:   adminRepo,
		logRepo:     logRepo,
		jwtSecret:   jwtSecret,
		tokenExpire: tokenExpire,
	}
}

func (s *adminService) AdminLogin(ctx *gin.Context, req *dto.AdminLoginDTO) (string, error) {
	admin, err := s.adminRepo.FindByAdminName(ctx, req.AdminName)
	if err != nil {
		logger.Logger.Error("查询管理员失败", zap.String("adminName", req.AdminName), zap.Error(err))
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(admin.AdminPassword), []byte(req.AdminPassword)); err != nil {
		logger.Logger.Error("管理员登录失败", zap.String("adminName", req.AdminName), zap.String("adminPassword", req.AdminPassword))
		return "", err
	}

	// 生成JWT token
	token, err := middleware.AdminGenerateToken(admin.AdminID, admin.AdminName)
	if err != nil {
		logger.Logger.Error("生成JWT token失败", zap.String("adminName", req.AdminName), zap.Error(err))
		return "", err
	}

	// 更新最后登录时间
	admin.LastLogin = time.Now()
	if err := s.adminRepo.Update(ctx, admin); err != nil {
		logger.Logger.Error("更新管理员最后登录时间失败", zap.Error(err))
	}

	// 记录登录日志
	if err := s.logRepo.CreateOperationLog(ctx, &entity.OperationLog{
		AdminID:           admin.AdminID,
		OperationIP:       ctx.ClientIP(),
		OperationLocation: "管理员登录成功",
		OperationBrowser:  ctx.Request.UserAgent(),
		OperationOS:       ctx.Request.UserAgent(),
		OperationMethod:   ctx.Request.Method,
		OperationPath:     ctx.Request.URL.Path,
		OperationModule:   "管理员登录",
		OperationContent:  "管理员登录成功",
		OperationStatus:   constants.LogStatusSuccess,
		CreatedAt:         time.Now(),
	}); err != nil {
		logger.Logger.Error("记录登录日志失败", zap.Error(err))
	}

	return token, nil
}

func (s *adminService) CreateAdmin(ctx *gin.Context, req *dto.CreateAdminDTO) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.AdminPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	admin := &entity.Admin{
		DepartmentID:  req.DepartmentID,
		AdminName:     req.AdminName,
		AdminPassword: string(hashedPassword),
		AdminEmail:    req.AdminEmail,
		AdminPhone:    req.AdminPhone,
		AdminSex:      req.AdminSex,
		AdminStatus:   constants.StatusEnabled, // 默认启用
		CreatedAt:     time.Now(),
	}

	if err := s.adminRepo.Create(ctx, admin); err != nil {
		return err
	}

	// 更新角色关联
	if err := s.adminRepo.UpdateAdminRoles(ctx, admin.AdminID, req.RoleIDs); err != nil {
		return err
	}

	// 更新岗位关联
	if len(req.PostIDs) > 0 {
		if err := s.adminRepo.UpdateAdminPosts(ctx, admin.AdminID, req.PostIDs); err != nil {
			return err
		}
	}

	return nil
}

func (s *adminService) UpdateAdmin(ctx *gin.Context, req *dto.UpdateAdminDTO) error {
	admin, err := s.adminRepo.FindByID(ctx, req.AdminID)
	if err != nil {
		logger.Logger.Error("查询管理员失败", zap.Int("id", req.AdminID), zap.Error(err))
		return err
	}

	admin.DepartmentID = req.DepartmentID
	admin.AdminName = req.AdminName
	if req.AdminPassword != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.AdminPassword), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		admin.AdminPassword = string(hashedPassword)
	}
	admin.AdminEmail = req.AdminEmail
	admin.AdminPhone = req.AdminPhone
	admin.AdminSex = req.AdminSex
	admin.AdminAvatar = req.AdminAvatar
	admin.AdminStatus = req.AdminStatus
	admin.UpdatedAt = time.Now()

	if err := s.adminRepo.Update(ctx, admin); err != nil {
		return err
	}

	// 更新角色关联
	if err := s.adminRepo.UpdateAdminRoles(ctx, admin.AdminID, req.RoleIDs); err != nil {
		return err
	}

	// 更新岗位关联
	if len(req.PostIDs) > 0 {
		if err := s.adminRepo.UpdateAdminPosts(ctx, admin.AdminID, req.PostIDs); err != nil {
			return err
		}
	}

	return nil
}

func (s *adminService) DeleteAdmin(ctx *gin.Context, adminID int) error {
	// 删除管理员
	if err := s.adminRepo.Delete(ctx, adminID); err != nil {
		return err
	}

	s.logRepo.CreateOperationLog(ctx, &entity.OperationLog{
		AdminID:           adminID,
		OperationIP:       ctx.ClientIP(),
		OperationLocation: "管理员删除成功",
		OperationBrowser:  ctx.Request.UserAgent(),
		OperationOS:       ctx.Request.UserAgent(),
		OperationMethod:   ctx.Request.Method,
		OperationPath:     ctx.Request.URL.Path,
		OperationModule:   "管理员删除",
		OperationContent:  "管理员删除成功",
		OperationStatus:   constants.LogStatusSuccess,
		CreatedAt:         time.Now(),
	})

	return nil
}

func (s *adminService) GetAdminInfo(ctx *gin.Context, adminID int) (*dto.AdminDTO, error) {
	admin, err := s.adminRepo.FindByID(ctx, adminID)
	if err != nil {
		return nil, err
	}

	info := &dto.AdminDTO{
		AdminID:      admin.AdminID,
		DepartmentID: admin.DepartmentID,
		AdminName:    admin.AdminName,
		AdminEmail:   admin.AdminEmail,
		AdminPhone:   admin.AdminPhone,
		AdminSex:     admin.AdminSex,
		AdminAvatar:  admin.AdminAvatar,
		AdminStatus:  admin.AdminStatus,
		AdminLogin:   admin.LastLogin.Format(time.DateTime),
	}

	// 转换角色信息
	info.Roles = make([]dto.RoleDTO, len(admin.Roles))
	for i, role := range admin.Roles {
		info.Roles[i] = dto.RoleDTO{
			RoleID:          role.RoleID,
			RoleName:        role.RoleName,
			RoleCode:        role.RoleCode,
			RoleDescription: role.RoleDescription,
			RoleStatus:      role.RoleStatus,
			CreatedAt:       role.CreatedAt.Format(time.DateTime),
		}
	}

	// 转换岗位信息
	info.Posts = make([]dto.PostDTO, len(admin.Posts))
	for i, post := range admin.Posts {
		info.Posts[i] = dto.PostDTO{
			PostID:   post.PostID,
			PostName: post.PostName,
			PostCode: post.PostCode,
		}
	}

	return info, nil
}

func (s *adminService) ListAdmins(ctx *gin.Context, page, size int) ([]*dto.AdminDTO, int64, error) {
	admins, total, err := s.adminRepo.List(ctx, page, size)
	if err != nil {
		logger.Logger.Error("查询管理员失败", zap.Int("page", page), zap.Int("size", size), zap.Error(err))
		return nil, 0, err
	}

	items := make([]*dto.AdminDTO, len(admins))
	for i, admin := range admins {
		items[i] = &dto.AdminDTO{
			AdminID:      admin.AdminID,
			DepartmentID: admin.DepartmentID,
			AdminName:    admin.AdminName,
			AdminEmail:   admin.AdminEmail,
			AdminPhone:   admin.AdminPhone,
			AdminSex:     admin.AdminSex,
			AdminAvatar:  admin.AdminAvatar,
			AdminStatus:  admin.AdminStatus,
			AdminLogin:   admin.LastLogin.Format(time.DateTime),
		}
	}

	return items, total, nil
}

// Role相关方法实现
func (s *adminService) CreateRole(ctx *gin.Context, req *dto.CreateRoleDTO) error {
	role := &entity.Role{
		RoleName:        req.RoleName,
		RoleCode:        req.RoleCode,
		RoleDescription: req.RoleDescription,
		RoleStatus:      1, // 默认启用
		CreatedAt:       time.Now(),
	}

	if err := s.adminRepo.CreateRole(ctx, role); err != nil {
		return err
	}

	// 更新菜单关联
	if err := s.adminRepo.UpdateRoleMenus(ctx, role.RoleID, req.MenuIDs); err != nil {
		return err
	}

	// 更新部门关联
	if len(req.DepartmentIDs) > 0 {
		if err := s.adminRepo.UpdateRoleDepartments(ctx, role.RoleID, req.DepartmentIDs); err != nil {
			return err
		}
	}

	return nil
}

func (s *adminService) UpdateRole(ctx *gin.Context, req *dto.UpdateRoleDTO) error {
	role, err := s.adminRepo.FindRoleByID(ctx, req.RoleID)
	if err != nil {
		return err
	}

	role.RoleName = req.RoleName
	role.RoleCode = req.RoleCode
	role.RoleDescription = req.RoleDescription
	role.RoleStatus = req.RoleStatus

	if err := s.adminRepo.UpdateRole(ctx, role); err != nil {
		return err
	}

	// 更新菜单关联
	if err := s.adminRepo.UpdateRoleMenus(ctx, role.RoleID, req.MenuIDs); err != nil {
		return err
	}

	// 更新部门关联
	if len(req.DepartmentIDs) > 0 {
		if err := s.adminRepo.UpdateRoleDepartments(ctx, role.RoleID, req.DepartmentIDs); err != nil {
			return err
		}
	}

	return nil
}

func (s *adminService) DeleteRole(ctx *gin.Context, roleID int) error {
	return s.adminRepo.DeleteRole(ctx, roleID)
}

func (s *adminService) ListRoles(ctx *gin.Context, page, size int) ([]*dto.RoleDTO, int64, error) {
	roles, total, err := s.adminRepo.ListRoles(ctx, page, size)
	if err != nil {
		logger.Logger.Error("查询角色失败", zap.Int("page", page), zap.Int("size", size), zap.Error(err))
		return nil, 0, err
	}

	items := make([]*dto.RoleDTO, len(roles))
	for i, role := range roles {
		items[i] = &dto.RoleDTO{
			RoleID:          role.RoleID,
			RoleName:        role.RoleName,
			RoleCode:        role.RoleCode,
			RoleDescription: role.RoleDescription,
			RoleStatus:      role.RoleStatus,
			CreatedAt:       role.CreatedAt.Format("2006-01-02 15:04:05"),
		}
	}

	return items, total, nil
}

// Department相关方法实现
func (s *adminService) CreateDepartment(ctx *gin.Context, req *dto.CreateDepartmentDTO) error {
	dept := &entity.Department{
		ParentID:              req.ParentID,
		DepartmentName:        req.DepartmentName,
		DepartmentCode:        req.DepartmentCode,
		DepartmentDescription: req.DepartmentDescription,
		DepartmentLeader:      req.DepartmentLeader,
		DepartmentPhone:       req.DepartmentPhone,
		DepartmentEmail:       req.DepartmentEmail,
		DepartmentSort:        req.DepartmentSort,
		DepartmentStatus:      1, // 默认启用
		CreatedAt:             time.Now(),
	}

	return s.adminRepo.CreateDepartment(ctx, dept)
}

func (s *adminService) UpdateDepartment(ctx *gin.Context, req *dto.UpdateDepartmentDTO) error {
	dept, err := s.adminRepo.FindDepartmentByID(ctx, req.DepartmentID)
	if err != nil {
		return err
	}

	dept.ParentID = req.ParentID
	dept.DepartmentName = req.DepartmentName
	dept.DepartmentCode = req.DepartmentCode
	dept.DepartmentDescription = req.DepartmentDescription
	dept.DepartmentLeader = req.DepartmentLeader
	dept.DepartmentPhone = req.DepartmentPhone
	dept.DepartmentEmail = req.DepartmentEmail
	dept.DepartmentSort = req.DepartmentSort
	dept.DepartmentStatus = req.DepartmentStatus

	return s.adminRepo.UpdateDepartment(ctx, dept)
}

func (s *adminService) DeleteDepartment(ctx *gin.Context, deptID int) error {
	return s.adminRepo.DeleteDepartment(ctx, deptID)
}

func (s *adminService) GetDepartmentTree(ctx *gin.Context) ([]*dto.DepartmentTreeDTO, error) {
	depts, err := s.adminRepo.ListAllDepartments(ctx)
	if err != nil {
		return nil, err
	}

	// 构建部门树
	deptMap := make(map[int]*dto.DepartmentTreeDTO)
	var rootDepts []*dto.DepartmentTreeDTO

	// 首先转换所有部门
	for _, dept := range depts {
		deptTree := &dto.DepartmentTreeDTO{
			DepartmentID:          dept.DepartmentID,
			ParentID:              dept.ParentID,
			DepartmentName:        dept.DepartmentName,
			DepartmentCode:        dept.DepartmentCode,
			DepartmentDescription: dept.DepartmentDescription,
			DepartmentLeader:      dept.DepartmentLeader,
			DepartmentPhone:       dept.DepartmentPhone,
			DepartmentEmail:       dept.DepartmentEmail,
			DepartmentSort:        dept.DepartmentSort,
			DepartmentStatus:      dept.DepartmentStatus,
			Children:              make([]*dto.DepartmentTreeDTO, 0),
		}
		deptMap[dept.DepartmentID] = deptTree
	}

	// 构建树形结构
	for _, dept := range depts {
		if dept.ParentID == 0 {
			rootDepts = append(rootDepts, deptMap[dept.DepartmentID])
		} else {
			if parent, ok := deptMap[dept.ParentID]; ok {
				parent.Children = append(parent.Children, deptMap[dept.DepartmentID])
			}
		}
	}

	return rootDepts, nil
}

// 其他方法实现...
