package service

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"

	"blacklist/internal/model"
	"blacklist/internal/repository"
	"blacklist/pkg/utils"
)

type MerchantService struct {
	repo *repository.MerchantRepository
}

func NewMerchantService(repo *repository.MerchantRepository) *MerchantService {
	return &MerchantService{repo: repo}
}

// generateAPICredentials 生成API凭证
func generateAPICredentials() (string, string, error) {
	// 生成16字节的随机API Key
	keyBytes := make([]byte, 16)
	if _, err := rand.Read(keyBytes); err != nil {
		return "", "", err
	}
	apiKey := hex.EncodeToString(keyBytes)

	// 生成32字节的随机API Secret
	secretBytes := make([]byte, 32)
	if _, err := rand.Read(secretBytes); err != nil {
		return "", "", err
	}
	apiSecret := hex.EncodeToString(secretBytes)

	return apiKey, apiSecret, nil
}

// Create 创建商户
func (s *MerchantService) Create(merchant *model.Merchant) error {
	// 生成API凭证
	apiKey, apiSecret, err := generateAPICredentials()
	if err != nil {
		return fmt.Errorf("生成API凭证失败: %w", err)
	}

	merchant.APIKey = apiKey
	merchant.APISecret = apiSecret
	merchant.Status = model.MerchantStatusNormal

	return s.repo.Create(merchant)
}

// Update 更新商户
func (s *MerchantService) Update(merchant *model.Merchant) error {
	return s.repo.Update(merchant)
}

// Delete 删除商户
func (s *MerchantService) Delete(id uint) error {
	return s.repo.Delete(id)
}

// GetByID 获取商户详情
func (s *MerchantService) GetByID(id uint) (*model.Merchant, error) {
	return s.repo.GetByID(id)
}

// List 获取商户列表
type MerchantQuery struct {
	Name   string
	Status int
	Page   int
	Size   int
}

func (s *MerchantService) List(query *MerchantQuery) ([]model.Merchant, int64, error) {
	if query.Page <= 0 {
		query.Page = 1
	}
	if query.Size <= 0 {
		query.Size = 10
	}

	repoQuery := &repository.MerchantQuery{
		Name:   query.Name,
		Status: query.Status,
		Page:   query.Page,
		Size:   query.Size,
	}

	return s.repo.List(repoQuery)
}

// UpdateStatus 更新商户状态
func (s *MerchantService) UpdateStatus(id uint, status int) error {
	if status != model.MerchantStatusNormal && status != model.MerchantStatusDisabled {
		return fmt.Errorf("无效的状态值")
	}
	return s.repo.UpdateStatus(id, status)
}

// RegenerateAPICredentials 重新生成API凭证
func (s *MerchantService) RegenerateAPICredentials(id uint) (*model.Merchant, error) {
	merchant, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	apiKey, apiSecret, err := generateAPICredentials()
	if err != nil {
		return nil, fmt.Errorf("生成API凭证失败: %w", err)
	}

	merchant.APIKey = apiKey
	merchant.APISecret = apiSecret

	if err := s.repo.Update(merchant); err != nil {
		return nil, err
	}

	return merchant, nil
}

// Login 商户登录
func (s *MerchantService) Login(apiKey, apiSecret string) (string, error) {
	merchant, err := s.repo.GetByAPIKey(apiKey)
	if err != nil {
		return "", fmt.Errorf("商户不存在")
	}

	if merchant.Status != model.MerchantStatusNormal {
		return "", fmt.Errorf("商户已禁用")
	}

	if merchant.APISecret != apiSecret {
		return "", fmt.Errorf("API Secret不正确, %s != %s", merchant.APISecret, apiSecret)
	}

	// 生成Token
	token, err := utils.GenerateToken(merchant.ID)
	if err != nil {
		return "", fmt.Errorf("生成Token失败: %w", err)
	}

	// 更新Token和过期时间
	expiry := time.Now().Add(24 * time.Hour)
	if err := s.repo.UpdateToken(merchant.ID, token, expiry); err != nil {
		return "", fmt.Errorf("更新Token失败: %w", err)
	}

	return token, nil
}
