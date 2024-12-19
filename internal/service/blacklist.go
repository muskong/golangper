package service

import (
	"blacklist/internal/model"
	"blacklist/internal/repository"
	"fmt"
)

type BlacklistService struct {
	repo *repository.BlacklistRepository
}

func NewBlacklistService(repo *repository.BlacklistRepository) *BlacklistService {
	return &BlacklistService{repo: repo}
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
