package repository

import (
	"blackapp/internal/domain/entity"
	"context"
)

type AdminRepository interface {
	Create(ctx context.Context, admin *entity.Admin) error
	Update(ctx context.Context, admin *entity.Admin) error
	Delete(ctx context.Context, id int) error
	FindByID(ctx context.Context, id int) (*entity.Admin, error)
	FindByUsername(ctx context.Context, username string) (*entity.Admin, error)
	List(ctx context.Context, page, size int) ([]*entity.Admin, int64, error)
	UpdateStatus(ctx context.Context, id int, status int) error
}
