package service

import (
	"blackapp/internal/service/dto"
	"context"
)

type MerchantService interface {
	Create(ctx context.Context, req *dto.CreateMerchantDTO) error
	Update(ctx context.Context, req *dto.UpdateMerchantDTO) error
	Delete(ctx context.Context, id int) error
	GetByID(ctx context.Context, id int) (*dto.MerchantDTO, error)
	List(ctx context.Context, page, size int) ([]*dto.MerchantDTO, int64, error)
	UpdateStatus(ctx context.Context, id int, status int) error
	GenerateAPICredentials(ctx context.Context, id int) error
	Login(ctx context.Context, apiKey, apiSecret string) (string, error)
}
