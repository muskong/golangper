package impl

import (
	"admins/api/dto"
	"admins/domain/entity"
	"admins/domain/repository"
	"context"
	"time"
)

type configService struct {
	configRepo repository.ConfigRepository
}

func NewConfigService(configRepo repository.ConfigRepository) *configService {
	return &configService{configRepo: configRepo}
}

func (s *configService) Create(ctx context.Context, req dto.ConfigCreateRequest) error {
	config := &entity.Config{
		ConfigName:        req.ConfigName,
		ConfigKey:         req.ConfigKey,
		ConfigValue:       req.ConfigValue,
		ConfigType:        req.ConfigType,
		ConfigStatus:      1,
		ConfigDescription: req.ConfigDescription,
		CreatedAt:         time.Now(),
	}

	return s.configRepo.Create(ctx, config)
}

func (s *configService) Update(ctx context.Context, req dto.ConfigUpdateRequest) error {
	config, err := s.configRepo.FindByID(ctx, req.ConfigID)
	if err != nil {
		return err
	}

	config.ConfigName = req.ConfigName
	config.ConfigValue = req.ConfigValue
	config.ConfigType = req.ConfigType
	config.ConfigStatus = req.ConfigStatus
	config.ConfigDescription = req.ConfigDescription
	config.UpdatedAt = time.Now()

	return s.configRepo.Update(ctx, config)
}

func (s *configService) Delete(ctx context.Context, configID int) error {
	return s.configRepo.Delete(ctx, configID)
}

func (s *configService) GetByID(ctx context.Context, configID int) (*dto.ConfigInfo, error) {
	config, err := s.configRepo.FindByID(ctx, configID)
	if err != nil {
		return nil, err
	}

	return &dto.ConfigInfo{
		ConfigID:          config.ConfigID,
		ConfigName:        config.ConfigName,
		ConfigKey:         config.ConfigKey,
		ConfigValue:       config.ConfigValue,
		ConfigType:        config.ConfigType,
		ConfigStatus:      config.ConfigStatus,
		ConfigDescription: config.ConfigDescription,
		CreatedAt:         config.CreatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}

func (s *configService) GetByKey(ctx context.Context, configKey string) (*dto.ConfigInfo, error) {
	config, err := s.configRepo.FindByKey(ctx, configKey)
	if err != nil {
		return nil, err
	}

	return &dto.ConfigInfo{
		ConfigID:          config.ConfigID,
		ConfigName:        config.ConfigName,
		ConfigKey:         config.ConfigKey,
		ConfigValue:       config.ConfigValue,
		ConfigType:        config.ConfigType,
		ConfigStatus:      config.ConfigStatus,
		ConfigDescription: config.ConfigDescription,
		CreatedAt:         config.CreatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}

func (s *configService) List(ctx context.Context, query dto.PageQuery) (*dto.PageResponse, error) {
	configs, total, err := s.configRepo.List(ctx, (query.PageNum-1)*query.PageSize, query.PageSize)
	if err != nil {
		return nil, err
	}

	items := make([]dto.ConfigInfo, len(configs))
	for i, config := range configs {
		items[i] = dto.ConfigInfo{
			ConfigID:          config.ConfigID,
			ConfigName:        config.ConfigName,
			ConfigKey:         config.ConfigKey,
			ConfigValue:       config.ConfigValue,
			ConfigType:        config.ConfigType,
			ConfigStatus:      config.ConfigStatus,
			ConfigDescription: config.ConfigDescription,
			CreatedAt:         config.CreatedAt.Format("2006-01-02 15:04:05"),
		}
	}

	return &dto.PageResponse{
		Total:    total,
		List:     items,
		PageNum:  query.PageNum,
		PageSize: query.PageSize,
	}, nil
}

func (s *configService) RefreshCache(ctx context.Context) error {
	// TODO: 实现配置缓存刷新逻辑
	return nil
}
