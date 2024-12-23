package impl

import (
	"context"

	"blackapp/internal/service/dto"
	"blackapp/pkg/logger"
)

type blackappService struct {
	repo repository.blackappRepository
}

func NewblackappService(repo repository.blackappRepository) *blackappService {
	return &blackappService{repo: repo}
}

func (s *blackappService) Create(ctx context.Context, req *dto.CreateblackappDTO) error {
	blackapp := &entity.blackapp{
		Name:       req.Name,
		Phone:      req.Phone,
		IDCard:     req.IDCard,
		Email:      req.Email,
		Address:    req.Address,
		Remark:     req.Remark,
		Status:     0, // 待审核状态
		MerchantID: req.MerchantID,
	}

	if err := s.repo.Create(ctx, blackapp); err != nil {
		logger.Logger.Error("创建黑名单记录失败", err)
		return err
	}

	return nil
}

func (s *blackappService) Update(ctx context.Context, req *dto.UpdateblackappDTO) error {
	blackapp, err := s.repo.FindByID(ctx, req.ID)
	if err != nil {
		return err
	}

	blackapp.Name = req.Name
	blackapp.Phone = req.Phone
	blackapp.IDCard = req.IDCard
	blackapp.Email = req.Email
	blackapp.Address = req.Address
	blackapp.Remark = req.Remark
	blackapp.Status = req.Status

	return s.repo.Update(ctx, blackapp)
}

func (s *blackappService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

func (s *blackappService) GetByID(ctx context.Context, id uint) (*dto.blackappDTO, error) {
	blackapp, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return toblackappDTO(blackapp), nil
}

func (s *blackappService) List(ctx context.Context, page, size int) ([]*dto.blackappDTO, int64, error) {
	blackapps, total, err := s.repo.List(ctx, page, size)
	if err != nil {
		return nil, 0, err
	}

	dtos := make([]*dto.blackappDTO, len(blackapps))
	for i, blackapp := range blackapps {
		dtos[i] = toblackappDTO(blackapp)
	}

	return dtos, total, nil
}

func (s *blackappService) UpdateStatus(ctx context.Context, id uint, status int) error {
	return s.repo.UpdateStatus(ctx, id, status)
}

func (s *blackappService) Check(ctx context.Context, req *dto.CheckblackappDTO) (bool, error) {
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

func toblackappDTO(blackapp *entity.blackapp) *dto.blackappDTO {
	return &dto.blackappDTO{
		ID:         blackapp.ID,
		Name:       blackapp.Name,
		Phone:      blackapp.Phone,
		IDCard:     blackapp.IDCard,
		Email:      blackapp.Email,
		Address:    blackapp.Address,
		Remark:     blackapp.Remark,
		Status:     blackapp.Status,
		MerchantID: blackapp.MerchantID,
		CreatedAt:  blackapp.CreatedAt,
		UpdatedAt:  blackapp.UpdatedAt,
	}
}
