package impl

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"

	"merchants/domain/constants"
	"merchants/domain/entity"
	"merchants/domain/repository"
	"merchants/service/dto"
	"pkgs/logger"
	"pkgs/middleware"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type merchantService struct {
	repo            repository.MerchantRepository
	loginLogRepo    repository.LoginLogRepository
	jwtSecret       string
	tokenExpireTime time.Duration
}

func NewMerchantService(repo repository.MerchantRepository, loginLogRepo repository.LoginLogRepository, jwtSecret string, tokenExpireTime time.Duration) *merchantService {
	return &merchantService{
		repo:            repo,
		loginLogRepo:    loginLogRepo,
		jwtSecret:       jwtSecret,
		tokenExpireTime: tokenExpireTime,
	}
}

func (s *merchantService) Create(ctx *gin.Context, req *dto.CreateMerchantDTO) error {
	merchant := &entity.Merchant{
		Name:          req.Name,
		Address:       req.Address,
		ContactPerson: req.ContactPerson,
		ContactPhone:  req.ContactPhone,
		Remark:        req.Remark,
		IPWhitelist:   req.IPWhitelist,
		Status:        constants.StatusEnabled,
	}

	if err := s.repo.Create(ctx, merchant); err != nil {
		logger.Logger.Error("创建商户失败", zap.Error(err))
		return err
	}

	return s.GenerateAPICredentials(ctx, merchant.ID)
}

func (s *merchantService) Update(ctx *gin.Context, req *dto.UpdateMerchantDTO) error {
	merchant, err := s.repo.FindByID(ctx, req.ID)
	if err != nil {
		return err
	}

	merchant.Name = req.Name
	merchant.Address = req.Address
	merchant.ContactPerson = req.ContactPerson
	merchant.ContactPhone = req.ContactPhone
	merchant.Remark = req.Remark
	merchant.IPWhitelist = req.IPWhitelist

	return s.repo.Update(ctx, merchant)
}

func (s *merchantService) Delete(ctx *gin.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

func (s *merchantService) GetByID(ctx *gin.Context, id int) (*dto.MerchantDTO, error) {
	merchant, err := s.repo.FindByID(ctx, id)
	if err != nil {
		logger.Logger.Error("查询商户失败", zap.Int("id", id), zap.Error(err))
		return nil, err
	}

	return toMerchantDTO(merchant), nil
}

func (s *merchantService) List(ctx *gin.Context, page, size int) ([]*dto.MerchantDTO, int64, error) {
	merchants, total, err := s.repo.List(ctx, page, size)
	if err != nil {
		logger.Logger.Error("查询商户失败", zap.Int("page", page), zap.Int("size", size), zap.Error(err))
		return nil, 0, err
	}

	dtos := make([]*dto.MerchantDTO, len(merchants))
	for i, merchant := range merchants {
		dtos[i] = toMerchantDTO(merchant)
	}

	return dtos, total, nil
}

func (s *merchantService) UpdateStatus(ctx *gin.Context, id int, status int) error {
	return s.repo.UpdateStatus(ctx, id, status)
}

func (s *merchantService) GenerateAPICredentials(ctx *gin.Context, id int) error {
	merchant, err := s.repo.FindByID(ctx, id)
	if err != nil {
		logger.Logger.Error("查询商户失败", zap.Int("id", id), zap.Error(err))
		return err
	}

	// 生成 API Key
	apiKeyBytes := make([]byte, 32)
	if _, err := rand.Read(apiKeyBytes); err != nil {
		logger.Logger.Error("生成API Key失败", zap.Int("id", id), zap.Error(err))
		return err
	}
	merchant.APIKey = hex.EncodeToString(apiKeyBytes)

	// 生成 API Secret
	apiSecretBytes := make([]byte, 32)
	if _, err := rand.Read(apiSecretBytes); err != nil {
		logger.Logger.Error("生成API Secret失败", zap.Int("id", id), zap.Error(err))
		return err
	}
	merchant.APISecret = hex.EncodeToString(apiSecretBytes)

	return s.repo.Update(ctx, merchant)
}

func (s *merchantService) Login(ctx *gin.Context, req *dto.MerchantLoginDTO) (string, error) {
	merchant, err := s.repo.FindByAPIKey(ctx, req.APIKey)
	if err != nil {
		logger.Logger.Error("查询商户失败", zap.String("apiKey", req.APIKey), zap.Error(err))
		return "", err
	}

	if merchant.APISecret != req.APISecret {
		logger.Logger.Error("商户登录失败", zap.String("apiKey", req.APIKey), zap.String("apiSecret", req.APISecret), zap.Error(fmt.Errorf("invalid credentials")))
		return "", fmt.Errorf("invalid credentials")
	}

	tokenString, err := middleware.GenerateToken(merchant.ID)
	if err != nil {
		logger.Logger.Error("生成商户token失败", zap.Int("merchantID", merchant.ID), zap.Error(err))
		return "", err
	}

	loginLog := &entity.LoginLog{
		Type:      constants.UserTypeMerchant,
		UserID:    merchant.ID,
		IP:        ctx.ClientIP(),
		UserAgent: ctx.Request.UserAgent(),
		Status:    constants.LogStatusSuccess,
	}

	if err := s.loginLogRepo.Create(ctx, loginLog); err != nil {
		logger.Logger.Error("创建商户登录日志失败", zap.Int("merchantID", merchant.ID), zap.Error(err))
		return "", err
	}

	merchant.APIToken = tokenString
	merchant.TokenExpireTime = time.Now().Add(s.tokenExpireTime)

	if err := s.repo.UpdateToken(ctx, merchant.ID, tokenString, merchant.TokenExpireTime); err != nil {
		logger.Logger.Error("更新商户token失败", zap.Int("merchantID", merchant.ID), zap.Error(err))
		return "", err
	}

	logger.Logger.Info("商户登录成功", zap.String("token", tokenString), zap.Int("merchantID", merchant.ID))

	return tokenString, nil
}

func (s *merchantService) ListLoginLogs(ctx *gin.Context, userType int, page, size int) ([]*dto.LoginLogDTO, int64, error) {
	logs, total, err := s.loginLogRepo.List(ctx, userType, page, size)
	if err != nil {
		logger.Logger.Error("查询登录日志失败", zap.Int("userType", userType), zap.Int("page", page), zap.Int("size", size), zap.Error(err))
		return nil, 0, err
	}

	dtos := make([]*dto.LoginLogDTO, len(logs))
	for i, log := range logs {
		dtos[i] = toLoginLogDTO(log)
	}
	return dtos, total, nil
}

func toMerchantDTO(merchant *entity.Merchant) *dto.MerchantDTO {
	return &dto.MerchantDTO{
		ID:              merchant.ID,
		Name:            merchant.Name,
		Address:         merchant.Address,
		ContactPerson:   merchant.ContactPerson,
		ContactPhone:    merchant.ContactPhone,
		Remark:          merchant.Remark,
		Status:          merchant.Status,
		IPWhitelist:     merchant.IPWhitelist,
		APIKey:          merchant.APIKey,
		APISecret:       merchant.APISecret,
		APIToken:        merchant.APIToken,
		TokenExpireTime: merchant.TokenExpireTime,
		CreatedAt:       merchant.CreatedAt,
		UpdatedAt:       merchant.UpdatedAt,
	}
}

func toLoginLogDTO(log *entity.LoginLog) *dto.LoginLogDTO {
	return &dto.LoginLogDTO{
		ID:        log.ID,
		Type:      log.Type,
		UserID:    log.UserID,
		IP:        log.IP,
		UserAgent: log.UserAgent,
		Status:    log.Status,
		CreatedAt: log.CreatedAt,
	}
}
