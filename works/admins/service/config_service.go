package service

import (
	"admins/api/dto"
	"context"
)

type ConfigService interface {
	Create(ctx context.Context, req dto.ConfigCreateRequest) error
	Update(ctx context.Context, req dto.ConfigUpdateRequest) error
	Delete(ctx context.Context, configID int) error
	GetByID(ctx context.Context, configID int) (*dto.ConfigInfo, error)
	GetByKey(ctx context.Context, configKey string) (*dto.ConfigInfo, error)
	List(ctx context.Context, query dto.PageQuery) (*dto.PageResponse, error)
	RefreshCache(ctx context.Context) error
}
