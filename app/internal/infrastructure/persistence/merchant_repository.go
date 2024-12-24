package persistence

import (
	"context"
	"time"

	"blackapp/internal/domain/entity"
	"blackapp/pkg/database"
)

type merchantRepository struct{}

func NewMerchantRepository() *merchantRepository {
	return &merchantRepository{}
}

func (r *merchantRepository) Create(ctx context.Context, merchant *entity.Merchant) error {
	return database.DB.Create(merchant).Error
}

func (r *merchantRepository) Update(ctx context.Context, merchant *entity.Merchant) error {
	return database.DB.Save(merchant).Error
}

func (r *merchantRepository) Delete(ctx context.Context, id int) error {
	return database.DB.Delete(&entity.Merchant{}, id).Error
}

func (r *merchantRepository) FindByID(ctx context.Context, id int) (*entity.Merchant, error) {
	var merchant entity.Merchant
	err := database.DB.First(&merchant, id).Error
	return &merchant, err
}

func (r *merchantRepository) FindByAPIKey(ctx context.Context, apiKey string) (*entity.Merchant, error) {
	var merchant entity.Merchant
	err := database.DB.Where("api_key = ?", apiKey).First(&merchant).Error
	return &merchant, err
}

func (r *merchantRepository) List(ctx context.Context, page, size int) ([]*entity.Merchant, int64, error) {
	var merchants []*entity.Merchant
	var total int64

	offset := (page - 1) * size

	err := database.DB.Model(&entity.Merchant{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = database.DB.Offset(offset).Limit(size).Find(&merchants).Error
	return merchants, total, err
}

func (r *merchantRepository) UpdateStatus(ctx context.Context, id int, status int) error {
	return database.DB.Model(&entity.Merchant{}).Where("id = ?", id).Update("status", status).Error
}

func (r *merchantRepository) UpdateToken(ctx context.Context, id int, token string, expireTime time.Time) error {
	updates := map[string]interface{}{
		"api_token":         token,
		"token_expire_time": expireTime,
	}
	return database.DB.Model(&entity.Merchant{}).Where("id = ?", id).Updates(updates).Error
}
