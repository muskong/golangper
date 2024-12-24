package persistence

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"blackapp/internal/domain/entity"
	"blackapp/pkg/database"
)

type BlacklistRepository struct{}

func NewBlacklistRepository() *BlacklistRepository {
	return &BlacklistRepository{}
}

func (r *BlacklistRepository) Create(ctx context.Context, blacklist *entity.Blacklist) error {
	return database.DB.Create(blacklist).Error
}

func (r *BlacklistRepository) Update(ctx context.Context, blacklist *entity.Blacklist) error {
	return database.DB.Save(blacklist).Error
}

func (r *BlacklistRepository) Delete(ctx context.Context, id int) error {
	return database.DB.Delete(&entity.Blacklist{}, id).Error
}

func (r *BlacklistRepository) FindByID(ctx context.Context, id int) (*entity.Blacklist, error) {
	var blacklist entity.Blacklist
	err := database.DB.First(&blacklist, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &blacklist, nil
}

func (r *BlacklistRepository) List(ctx context.Context, page, size int) ([]*entity.Blacklist, int64, error) {
	var blacklists []*entity.Blacklist
	var total int64

	offset := (page - 1) * size

	err := database.DB.Model(&entity.Blacklist{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = database.DB.Offset(offset).Limit(size).Find(&blacklists).Error
	return blacklists, total, err
}

func (r *BlacklistRepository) UpdateStatus(ctx context.Context, id int, status int) error {
	return database.DB.Model(&entity.Blacklist{}).Where("id = ?", id).Update("status", status).Error
}

func (r *BlacklistRepository) CheckByPhone(ctx context.Context, phone string) (*entity.Blacklist, error) {
	var blacklist entity.Blacklist
	err := database.DB.Where("phone = ? AND status = ?", phone, 1).First(&blacklist).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &blacklist, nil
}

func (r *BlacklistRepository) CheckByIDCard(ctx context.Context, idCard string) (*entity.Blacklist, error) {
	var blacklist entity.Blacklist
	err := database.DB.Where("id_card = ? AND status = ?", idCard, 1).First(&blacklist).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &blacklist, nil
}

func (r *BlacklistRepository) CheckByName(ctx context.Context, name string) (*entity.Blacklist, error) {
	var blacklist entity.Blacklist
	err := database.DB.Where("name = ? AND status = ?", name, 1).First(&blacklist).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &blacklist, nil
}
