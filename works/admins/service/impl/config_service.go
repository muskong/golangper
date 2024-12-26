package impl

import (
	"admins/domain/entity"
	"admins/domain/repository"
	"admins/service/dto"
	"time"

	"github.com/gin-gonic/gin"
)

type configService struct {
	configRepo repository.ConfigRepository
}

func NewConfigService(configRepo repository.ConfigRepository) *configService {
	return &configService{configRepo: configRepo}
}

func (s *configService) Create(ctx *gin.Context, req dto.ConfigCreateDTO) error {
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

func (s *configService) Update(ctx *gin.Context, req dto.ConfigUpdateDTO) error {
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

func (s *configService) Delete(ctx *gin.Context, configID int) error {
	return s.configRepo.Delete(ctx, configID)
}

func (s *configService) GetByID(ctx *gin.Context, configID int) (*dto.ConfigInfo, error) {
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

func (s *configService) GetByKey(ctx *gin.Context, configKey string) (*dto.ConfigInfo, error) {
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

func (s *configService) List(ctx *gin.Context, page, pageSize int) ([]*dto.ConfigInfo, int64, error) {
	configs, total, err := s.configRepo.List(ctx, (page-1)*pageSize, pageSize)
	if err != nil {
		return nil, 0, err
	}

	items := make([]*dto.ConfigInfo, len(configs))
	for i, config := range configs {
		items[i] = &dto.ConfigInfo{
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

	return items, total, nil
}

func (s *configService) RefreshCache(ctx *gin.Context) error {
	// TODO: 实现配置缓存刷新逻辑
	return nil
}
