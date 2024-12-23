package impl

import (
	"blacklist/internal/domain/entity"
	"blacklist/internal/domain/repository"
	"blacklist/internal/service/dto"
)

type AdminService struct {
	adminRepo repository.AdminRepository
}

func (s *AdminService) CreateAdmin(dto *dto.AdminDTO) error {
	// 实现用户创建逻辑
	return s.adminRepo.Create(&entity.Admin{
		Username: dto.Username,
		Email:    dto.Email,
	})
}

func (s *AdminService) FindAdminByID(id string) (*dto.AdminDTO, error) {
	admin, err := s.adminRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return &dto.AdminDTO{
		ID:       admin.ID,
		Username: admin.Username,
		Email:    admin.Email,
	}, nil
}
