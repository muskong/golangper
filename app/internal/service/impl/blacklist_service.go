package impl

import (
	"context"

	"blackapp/internal/domain/entity"
	"blackapp/internal/domain/repository"
	"blackapp/internal/service/dto"
	"blackapp/pkg/logger"

	"go.uber.org/zap"
)

type BlacklistService struct {
	repo repository.BlacklistRepository
}

func NewBlacklistService(repo repository.BlacklistRepository) *BlacklistService {
	return &BlacklistService{repo: repo}
}

func (s *BlacklistService) Create(ctx context.Context, req *dto.CreateBlacklistDTO) error {
	blacklist := &entity.Blacklist{
		Name:    req.Name,
		Phone:   req.Phone,
		IDCard:  req.IDCard,
		Email:   req.Email,
		Address: req.Address,
		Remark:  req.Remark,
		Status:  0, // 待审核状态
	}

	if err := s.repo.Create(ctx, blacklist); err != nil {
		logger.Logger.Error("创建黑名单记录失败", zap.Error(err))
		return err
	}

	return nil
}

func (s *BlacklistService) Update(ctx context.Context, req *dto.UpdateBlacklistDTO) error {
	blacklist, err := s.repo.FindByID(ctx, req.ID)
	if err != nil {
		return err
	}

	blacklist.Name = req.Name
	blacklist.Phone = req.Phone
	blacklist.IDCard = req.IDCard
	blacklist.Email = req.Email
	blacklist.Address = req.Address
	blacklist.Remark = req.Remark
	blacklist.Status = req.Status

	return s.repo.Update(ctx, blacklist)
}

func (s *BlacklistService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

func (s *BlacklistService) GetByID(ctx context.Context, id int) (*dto.BlacklistDTO, error) {
	blacklist, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return toBlacklistDTO(blacklist), nil
}

func (s *BlacklistService) List(ctx context.Context, page, size int) ([]*dto.BlacklistDTO, int64, error) {
	blacklists, total, err := s.repo.List(ctx, page, size)
	if err != nil {
		return nil, 0, err
	}

	dtos := make([]*dto.BlacklistDTO, len(blacklists))
	for i, blacklist := range blacklists {
		dtos[i] = toBlacklistDTO(blacklist)
	}

	return dtos, total, nil
}

func (s *BlacklistService) UpdateStatus(ctx context.Context, id int, status int) error {
	return s.repo.UpdateStatus(ctx, id, status)
}

func (s *BlacklistService) Check(ctx context.Context, req *dto.CheckBlacklistDTO) (bool, error) {
	// 按照优先级依次检查手机号、身份证、姓名
	if req.Phone != "" {
		if _, err := s.repo.CheckByPhone(ctx, req.Phone); err == nil {
			return true, nil
		}
	}

	if req.IDCard != "" {
		if _, err := s.repo.CheckByIDCard(ctx, req.IDCard); err == nil {
			return true, nil
		}
	}

	if req.Name != "" {
		if _, err := s.repo.CheckByName(ctx, req.Name); err == nil {
			return true, nil
		}
	}

	return false, nil
}

func toBlacklistDTO(blacklist *entity.Blacklist) *dto.BlacklistDTO {
	return &dto.BlacklistDTO{
		ID:        blacklist.ID,
		Name:      blacklist.Name,
		Phone:     blacklist.Phone,
		IDCard:    blacklist.IDCard,
		Email:     blacklist.Email,
		Address:   blacklist.Address,
		Remark:    blacklist.Remark,
		Status:    blacklist.Status,
		CreatedAt: blacklist.CreatedAt,
		UpdatedAt: blacklist.UpdatedAt,
	}
}
