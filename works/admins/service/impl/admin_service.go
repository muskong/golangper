package impl

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"go.uber.org/zap"

	"admins/domain/constants"
	"admins/domain/entity"
	"admins/domain/repository"
	"admins/service/dto"
	"pkgs/logger"
	"pkgs/middleware"

	"github.com/gin-gonic/gin"
)

type adminService struct {
	adminRepo   repository.AdminRepository
	jwtSecret   string
	tokenExpire time.Duration
}

func NewAdminService(
	adminRepo repository.AdminRepository,
	jwtSecret string,
	tokenExpire time.Duration,
) *adminService {
	return &adminService{
		adminRepo:   adminRepo,
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

	if admin == nil || admin.AdminPassword != hashPassword(req.AdminPassword) {
		logger.Logger.Error("管理员登录失败", zap.String("adminName", req.AdminName), zap.String("adminPassword", req.AdminPassword))
		return "", fmt.Errorf("invalid credentials")
	}

	tokenString, err := middleware.GenerateToken(admin.AdminID)
	if err != nil {
		logger.Logger.Error("生成管理员token失败", zap.Error(err))
		return "", err
	}

	// 更新最后登录时间
	admin.LastLogin = time.Now()
	if err := s.adminRepo.Update(ctx, admin); err != nil {
		logger.Logger.Error("更新管理员最后登录时间失败", zap.Error(err))
	}

	return tokenString, nil
}

func (s *adminService) CreateAdmin(ctx *gin.Context, req *dto.CreateAdminDTO) error {
	admin := &entity.Admin{
		AdminName:     req.AdminName,
		AdminPassword: hashPassword(req.AdminPassword),
		AdminEmail:    req.AdminEmail,
		AdminPhone:    req.AdminPhone,
		AdminSex:      req.AdminSex,
		AdminAvatar:   req.AdminAvatar,
		AdminStatus:   constants.StatusEnabled,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	return s.adminRepo.Create(ctx, admin)
}

func (s *adminService) UpdateAdmin(ctx *gin.Context, req *dto.UpdateAdminDTO) error {
	admin, err := s.adminRepo.FindByID(ctx, req.AdminID)
	if err != nil {
		logger.Logger.Error("查询管理员失败", zap.Int("id", req.AdminID), zap.Error(err))
		return err
	}

	admin.AdminName = req.AdminName
	if req.AdminPassword != "" {
		admin.AdminPassword = hashPassword(req.AdminPassword)
	}
	admin.AdminEmail = req.AdminEmail
	admin.AdminPhone = req.AdminPhone
	admin.AdminSex = req.AdminSex
	admin.AdminAvatar = req.AdminAvatar
	admin.AdminStatus = req.AdminStatus
	admin.UpdatedAt = time.Now()

	return s.adminRepo.Update(ctx, admin)
}

func (s *adminService) DeleteAdmin(ctx *gin.Context, id int) error {
	return s.adminRepo.Delete(ctx, id)
}

func (s *adminService) GetAdminByID(ctx *gin.Context, id int) (*dto.AdminDTO, error) {
	admin, err := s.adminRepo.FindByID(ctx, id)
	if err != nil {
		logger.Logger.Error("查询管理员失败", zap.Int("id", id), zap.Error(err))
		return nil, err
	}
	return toAdminDTO(admin), nil
}

func (s *adminService) ListAdmins(ctx *gin.Context, page, size int) ([]*dto.AdminDTO, int64, error) {
	admins, total, err := s.adminRepo.List(ctx, page, size)
	if err != nil {
		logger.Logger.Error("查询管理员失败", zap.Int("page", page), zap.Int("size", size), zap.Error(err))
		return nil, 0, err
	}

	dtos := make([]*dto.AdminDTO, len(admins))
	for i, admin := range admins {
		dtos[i] = toAdminDTO(admin)
	}
	return dtos, total, nil
}

func (s *adminService) UpdateAdminStatus(ctx *gin.Context, id int, status int) error {
	return s.adminRepo.UpdateStatus(ctx, id, status)
}

func toAdminDTO(admin *entity.Admin) *dto.AdminDTO {
	return &dto.AdminDTO{
		AdminID:     admin.AdminID,
		AdminName:   admin.AdminName,
		AdminEmail:  admin.AdminEmail,
		AdminPhone:  admin.AdminPhone,
		AdminSex:    admin.AdminSex,
		AdminAvatar: admin.AdminAvatar,
		AdminStatus: admin.AdminStatus,
		AdminLogin:  admin.LastLogin.Format("2006-01-02 15:04:05"),
	}
}

func hashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}
