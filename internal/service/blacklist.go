package service

import (
	"blacklist/internal/model"
	"blacklist/internal/repository"
	"fmt"
	"context"
	"time"
	"log"
)

type BlacklistService struct {
	repo          repository.BlacklistRepository
	queryLogRepo  repository.BlacklistQueryLogRepository
}

func NewBlacklistService(
	repo repository.BlacklistRepository,
	queryLogRepo repository.BlacklistQueryLogRepository,
) *BlacklistService {
	return &BlacklistService{
		repo:         repo,
		queryLogRepo: queryLogRepo,
	}
}

func (s *BlacklistService) Create(user *model.BlacklistUser) error {
	return s.repo.Create(user)
}

func (s *BlacklistService) GetByID(id uint) (*model.BlacklistUser, error) {
	return s.repo.GetByID(id)
}

func (s *BlacklistService) Update(user *model.BlacklistUser) error {
	return s.repo.Update(user)
}

func (s *BlacklistService) Delete(id uint) error {
	return s.repo.Delete(id)
}

type BlacklistUserQuery struct {
	Name    string
	Phone   string
	IDCard  string
	Email   string
	Address string
	Remark  string
	Page    int
	Size    int
}

func (s *BlacklistService) List(query *BlacklistUserQuery) ([]model.BlacklistUser, int64, error) {
	if query.Page <= 0 {
		query.Page = 1
	}
	if query.Size <= 0 {
		query.Size = 10
	}

	repoQuery := &repository.BlacklistUserQuery{
		Name:    query.Name,
		Phone:   query.Phone,
		IDCard:  query.IDCard,
		Email:   query.Email,
		Address: query.Address,
		Remark:  query.Remark,
		Page:    query.Page,
		Size:    query.Size,
	}

	return s.repo.List(repoQuery)
}

// CheckPhoneExists 检查手机号是否已被列入黑名单
func (s *BlacklistService) CheckPhoneExists(phone string) (bool, error) {
	if phone == "" {
		return false, fmt.Errorf("手机号不能为空")
	}
	return s.repo.ExistsByPhone(phone)
}

// GetByPhone 根据手机号获取黑名单用户信息
func (s *BlacklistService) GetByPhone(phone string) (*model.BlacklistUser, error) {
	if phone == "" {
		return nil, fmt.Errorf("手机号不能为空")
	}
	return s.repo.GetByPhone(phone)
}

// ExistsQuery 存在性检查的查询参数
type ExistsQuery struct {
	Phone  string
	IDCard string
	Name   string
}

// CheckExists 检查用户是否存在
func (s *BlacklistService) CheckExists(ctx context.Context, phone string, merchantID uint, ip, userAgent string) (bool, error) {
	exists, err := s.repo.ExistsByPhone(ctx, phone)
	if err != nil {
		return false, err
	}

	// 记录查询日志
	if err := s.RecordQueryLog(ctx, merchantID, phone, ip, userAgent, exists); err != nil {
		// 这里只记录日志错误,不影响主流程
		log.Printf("记录查询日志失败: %v", err)
	}

	return exists, nil
}

// GetByIDCard 根据身份证号获取用户信息
func (s *BlacklistService) GetByIDCard(idCard string) (*model.BlacklistUser, error) {
	if idCard == "" {
		return nil, fmt.Errorf("身份证号不能为空")
	}
	return s.repo.GetByIDCard(idCard)
}

// GetByName 根据姓名获取用户信息列表
func (s *BlacklistService) GetByName(name string) ([]model.BlacklistUser, error) {
	if name == "" {
		return nil, fmt.Errorf("姓名不能为空")
	}
	return s.repo.GetByName(name)
}

// RecordQueryLog 记录查询日志
func (s *BlacklistService) RecordQueryLog(ctx context.Context, merchantID uint, phone, ip, userAgent string, result bool) error {
	log := &model.BlacklistQueryLog{
		MerchantID: merchantID,
		Phone:      phone,
		QueryTime:  time.Now(),
		IP:         ip,
		UserAgent:  userAgent,
		Result:     result,
	}
	return s.queryLogRepo.Create(ctx, log)
}

// GetQueryLogs 获取查询日志
func (s *BlacklistService) GetQueryLogs(ctx context.Context, merchantID uint, page, pageSize int) ([]model.BlacklistQueryLog, int64, error) {
	return s.queryLogRepo.FindByMerchantID(ctx, merchantID, page, pageSize)
}

// GetQueryLogsByPhone 获取指定手机号的查询日志
func (s *BlacklistService) GetQueryLogsByPhone(ctx context.Context, phone string, page, pageSize int) ([]model.BlacklistQueryLog, int64, error) {
	return s.queryLogRepo.FindByPhone(ctx, phone, page, pageSize)
}

// GetAllQueryLogs 获取所有查询日志(管理后台)
func (s *BlacklistService) GetAllQueryLogs(ctx context.Context, page, pageSize int) ([]model.BlacklistQueryLog, int64, error) {
	var logs []model.BlacklistQueryLog
	var total int64

	offset := (page - 1) * pageSize

	err := s.queryLogRepo.DB().WithContext(ctx).Model(&model.BlacklistQueryLog{}).
		Count(&total).
		Preload("Merchant"). // 预加载商户信息
		Order("query_time DESC").
		Offset(offset).
		Limit(pageSize).
		Find(&logs).Error

	return logs, total, err
}

// GetMerchantQueryLogs 获取指定商户的查询日志(管理后台)
func (s *BlacklistService) GetMerchantQueryLogs(ctx context.Context, merchantID uint, page, pageSize int) ([]model.BlacklistQueryLog, int64, error) {
	logs, total, err := s.queryLogRepo.FindByMerchantID(ctx, merchantID, page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	// 预加载商户信息
	if err := s.queryLogRepo.DB().Preload("Merchant").Find(&logs).Error; err != nil {
		return nil, 0, err
	}

	return logs, total, err
}
