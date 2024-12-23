package repository

import (
	"context"
)

type blackappRepository interface {
	Create(ctx context.Context, blackapp *entity.blackapp) error
	Update(ctx context.Context, blackapp *entity.blackapp) error
	Delete(ctx context.Context, id uint) error
	FindByID(ctx context.Context, id uint) (*entity.blackapp, error)
	List(ctx context.Context, page, size int) ([]*entity.blackapp, int64, error)
	UpdateStatus(ctx context.Context, id uint, status int) error

	// 验证用户是否在黑名单中
	CheckByPhone(ctx context.Context, phone string) (*entity.blackapp, error)
	CheckByIDCard(ctx context.Context, idCard string) (*entity.blackapp, error)
	CheckByName(ctx context.Context, name string) (*entity.blackapp, error)
}
