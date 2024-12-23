package service

import (
	"blacklist/internal/model"
	"blacklist/internal/repository"
	"blacklist/pkg/utils"
	"fmt"
)

type AdminService struct {
	repo repository.AdminRepository
}

func NewAdminService(repo repository.AdminRepository) *AdminService {
	return &AdminService{repo: repo}
}

func (s *AdminService) Create(admin *model.Admin) error {
	// 检查用户名是否已存在
	exists, _ := s.repo.GetByUsername(admin.Username)
	if exists != nil {
		return fmt.Errorf("用户名已存在")
	}

	// 加密密码
	hashedPassword, err := utils.HashPassword(admin.Password)
	if err != nil {
		return err
	}
	admin.Password = hashedPassword

	return s.repo.Create(admin)
}

func (s *AdminService) Update(admin *model.Admin) error {
	return s.repo.Update(admin)
}

func (s *AdminService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *AdminService) GetByID(id uint) (*model.Admin, error) {
	return s.repo.GetByID(id)
}

func (s *AdminService) List(page, pageSize int) ([]model.Admin, int64, error) {
	return s.repo.List(page, pageSize)
}
