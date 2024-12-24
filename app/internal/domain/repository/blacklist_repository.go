package repository

import (
	"blackapp/internal/domain/entity"
	"context"
)

type BlacklistRepository interface {
	Create(ctx context.Context, blacklist *entity.Blacklist) error
	Update(ctx context.Context, blacklist *entity.Blacklist) error
	Delete(ctx context.Context, id uint) error
	FindByID(ctx context.Context, id uint) (*entity.Blacklist, error)
	List(ctx context.Context, page, size int) ([]*entity.Blacklist, int64, error)
	UpdateStatus(ctx context.Context, id uint, status int) error

	// 验证用户是否在黑名单中
	CheckByPhone(ctx context.Context, phone string) (*entity.Blacklist, error)
	CheckByIDCard(ctx context.Context, idCard string) (*entity.Blacklist, error)
	CheckByName(ctx context.Context, name string) (*entity.Blacklist, error)
}
