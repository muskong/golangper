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
	admin, err := s.adminRepo.FindByUsername(ctx, req.Username)
	if err != nil {
		logger.Logger.Error("查询管理员失败", zap.String("username", req.Username), zap.Error(err))
		return "", err
	}

	if admin == nil || admin.Password != hashPassword(req.Password) {
		logger.Logger.Error("管理员登录失败", zap.String("username", req.Username), zap.String("password", req.Password))
		return "", fmt.Errorf("invalid credentials")
	}

	tokenString, err := middleware.GenerateToken(admin.ID)
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
		Username:  req.Username,
		Password:  hashPassword(req.Password),
		Name:      req.Name,
		Status:    constants.StatusEnabled,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return s.adminRepo.Create(ctx, admin)
}

func (s *adminService) UpdateAdmin(ctx *gin.Context, req *dto.UpdateAdminDTO) error {
	admin, err := s.adminRepo.FindByID(ctx, req.ID)
	if err != nil {
		logger.Logger.Error("查询管理员失败", zap.Int("id", req.ID), zap.Error(err))
		return err
	}

	admin.Username = req.Username
	if req.Password != "" {
		admin.Password = hashPassword(req.Password)
	}
	admin.Name = req.Name
	admin.Status = req.Status
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
		ID:        admin.ID,
		Username:  admin.Username,
		Name:      admin.Name,
		Status:    admin.Status,
		LastLogin: admin.LastLogin.Format("2006-01-02 15:04:05"),
	}
}

func hashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}
