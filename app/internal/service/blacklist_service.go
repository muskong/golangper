package service

import (
	"blackapp/internal/service/dto"
	"context"
)

type BlacklistService interface {
	Create(ctx context.Context, req *dto.CreateBlacklistDTO) error
	Update(ctx context.Context, req *dto.UpdateBlacklistDTO) error
	Delete(ctx context.Context, id uint) error
	GetByID(ctx context.Context, id uint) (*dto.BlacklistDTO, error)
	List(ctx context.Context, page, size int) ([]*dto.BlacklistDTO, int64, error)
	UpdateStatus(ctx context.Context, id uint, status int) error
	Check(ctx context.Context, req *dto.CheckBlacklistDTO) (bool, error)
}
