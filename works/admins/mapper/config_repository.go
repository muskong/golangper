package mapper

import (
	"admins/domain/entity"
	"admins/domain/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type configRepository struct {
	db *gorm.DB
}

func NewConfigRepository(db *gorm.DB) repository.ConfigRepository {
	return &configRepository{db: db}
}

func (r *configRepository) Create(ctx *gin.Context, config *entity.Config) error {
	return r.db.Create(config).Error
}

func (r *configRepository) Update(ctx *gin.Context, config *entity.Config) error {
	return r.db.Save(config).Error
}

func (r *configRepository) Delete(ctx *gin.Context, configID int) error {
	return r.db.Delete(&entity.Config{}, configID).Error
}

func (r *configRepository) FindByID(ctx *gin.Context, configID int) (*entity.Config, error) {
	var config entity.Config
	err := r.db.Where("id = ?", configID).First(&config).Error
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func (r *configRepository) FindByKey(ctx *gin.Context, configKey string) (*entity.Config, error) {
	var config entity.Config
	err := r.db.Where("config_key = ?", configKey).First(&config).Error
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func (r *configRepository) List(ctx *gin.Context, offset, limit int) ([]*entity.Config, int64, error) {
	var configs []*entity.Config
	var total int64
	err := r.db.Offset(offset).Limit(limit).Find(&configs).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	return configs, total, nil
}
