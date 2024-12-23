package repository

import (
	"blackapp/internal/domain/entity"
	"context"
	"time"
)

type MerchantRepository interface {
	Create(ctx context.Context, merchant *entity.Merchant) error
	Update(ctx context.Context, merchant *entity.Merchant) error
	Delete(ctx context.Context, id uint) error
	FindByID(ctx context.Context, id uint) (*entity.Merchant, error)
	FindByAPIKey(ctx context.Context, apiKey string) (*entity.Merchant, error)
	List(ctx context.Context, page, size int) ([]*entity.Merchant, int64, error)
	UpdateStatus(ctx context.Context, id uint, status int) error
	UpdateToken(ctx context.Context, id uint, token string, expireTime time.Time) error
}
