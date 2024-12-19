package service

import (
	"blacklist/internal/model"
	"blacklist/internal/repository"
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

func (s *BlacklistService) List(page, pageSize int) ([]model.BlacklistUser, int64, error) {
	return s.repo.List(page, pageSize)
}
