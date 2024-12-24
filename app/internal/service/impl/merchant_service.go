package impl

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"

	"blackapp/internal/domain/entity"
	"blackapp/internal/domain/repository"
	"blackapp/internal/service/dto"
	"blackapp/pkg/logger"

	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
)

type merchantService struct {
	repo            repository.MerchantRepository
	jwtSecret       string
	tokenExpireTime time.Duration
}

func NewMerchantService(repo repository.MerchantRepository, jwtSecret string, tokenExpireTime time.Duration) *merchantService {
	return &merchantService{
		repo:            repo,
		jwtSecret:       jwtSecret,
		tokenExpireTime: tokenExpireTime,
	}
}

func (s *merchantService) Create(ctx context.Context, req *dto.CreateMerchantDTO) error {
	merchant := &entity.Merchant{
		Name:          req.Name,
		Address:       req.Address,
		ContactPerson: req.ContactPerson,
		ContactPhone:  req.ContactPhone,
		Remark:        req.Remark,
		IPWhitelist:   req.IPWhitelist,
		Status:        1,
	}

	if err := s.repo.Create(ctx, merchant); err != nil {
		logger.Logger.Error("创建商户失败", zap.Error(err))
		return err
	}

	return s.GenerateAPICredentials(ctx, merchant.ID)
}

func (s *merchantService) Update(ctx context.Context, req *dto.UpdateMerchantDTO) error {
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

func (s *merchantService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

func (s *merchantService) GetByID(ctx context.Context, id int) (*dto.MerchantDTO, error) {
	merchant, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return toMerchantDTO(merchant), nil
}

func (s *merchantService) List(ctx context.Context, page, size int) ([]*dto.MerchantDTO, int64, error) {
	merchants, total, err := s.repo.List(ctx, page, size)
	if err != nil {
		return nil, 0, err
	}

	dtos := make([]*dto.MerchantDTO, len(merchants))
	for i, merchant := range merchants {
		dtos[i] = toMerchantDTO(merchant)
	}

	return dtos, total, nil
}

func (s *merchantService) UpdateStatus(ctx context.Context, id int, status int) error {
	return s.repo.UpdateStatus(ctx, id, status)
}

func (s *merchantService) GenerateAPICredentials(ctx context.Context, id int) error {
	merchant, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	// 生成 API Key
	apiKeyBytes := make([]byte, 32)
	if _, err := rand.Read(apiKeyBytes); err != nil {
		return err
	}
	merchant.APIKey = hex.EncodeToString(apiKeyBytes)

	// 生成 API Secret
	apiSecretBytes := make([]byte, 32)
	if _, err := rand.Read(apiSecretBytes); err != nil {
		return err
	}
	merchant.APISecret = hex.EncodeToString(apiSecretBytes)

	return s.repo.Update(ctx, merchant)
}

func (s *merchantService) Login(ctx context.Context, apiKey, apiSecret string) (string, error) {
	merchant, err := s.repo.FindByAPIKey(ctx, apiKey)
	if err != nil {
		return "", err
	}

	if merchant.APISecret != apiSecret {
		return "", fmt.Errorf("invalid credentials")
	}

	// 使用配置的过期时间
	expireTime := time.Now().Add(s.tokenExpireTime)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"merchant_id": merchant.ID,
		"exp":         expireTime.Unix(),
	})

	tokenString, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", err
	}

	merchant.APIToken = tokenString
	merchant.TokenExpireTime = expireTime

	if err := s.repo.UpdateToken(ctx, merchant.ID, tokenString, merchant.TokenExpireTime); err != nil {
		return "", err
	}

	return tokenString, nil
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
