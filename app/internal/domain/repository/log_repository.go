package repository

import (
	"blackapp/internal/domain/entity"
	"context"
)

type LoginLogRepository interface {
	Create(ctx context.Context, log *entity.LoginLog) error
	List(ctx context.Context, userType int, page, size int) ([]*entity.LoginLog, int64, error)
}

type QueryLogRepository interface {
	Create(ctx context.Context, log *entity.QueryLog) error
	List(ctx context.Context, merchantID int, page, size int) ([]*entity.QueryLog, int64, error)
}
