package repository

import (
	"admins/domain/entity"
	"context"
)

type ConfigRepository interface {
	Create(ctx context.Context, config *entity.Config) error
	Update(ctx context.Context, config *entity.Config) error
	Delete(ctx context.Context, configID int) error
	FindByID(ctx context.Context, configID int) (*entity.Config, error)
	FindByKey(ctx context.Context, configKey string) (*entity.Config, error)
	List(ctx context.Context, offset, limit int) ([]*entity.Config, int64, error)
}
