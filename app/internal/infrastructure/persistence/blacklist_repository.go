package persistence

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"blackapp/pkg/database"
)

type blackappRepository struct{}

func NewblackappRepository() *blackappRepository {
	return &blackappRepository{}
}

func (r *blackappRepository) Create(ctx context.Context, blackapp *entity.blackapp) error {
	return database.DB.Create(blackapp).Error
}

func (r *blackappRepository) Update(ctx context.Context, blackapp *entity.blackapp) error {
	return database.DB.Save(blackapp).Error
}

func (r *blackappRepository) Delete(ctx context.Context, id uint) error {
	return database.DB.Delete(&entity.blackapp{}, id).Error
}

func (r *blackappRepository) FindByID(ctx context.Context, id uint) (*entity.blackapp, error) {
	var blackapp entity.blackapp
	err := database.DB.First(&blackapp, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &blackapp, nil
}

func (r *blackappRepository) List(ctx context.Context, page, size int) ([]*entity.blackapp, int64, error) {
	var blackapps []*entity.blackapp
	var total int64

	offset := (page - 1) * size

	err := database.DB.Model(&entity.blackapp{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = database.DB.Offset(offset).Limit(size).Find(&blackapps).Error
	return blackapps, total, err
}

func (r *blackappRepository) UpdateStatus(ctx context.Context, id uint, status int) error {
	return database.DB.Model(&entity.blackapp{}).Where("id = ?", id).Update("status", status).Error
}

func (r *blackappRepository) CheckByPhone(ctx context.Context, phone string) (*entity.blackapp, error) {
	var blackapp entity.blackapp
	err := database.DB.Where("phone = ? AND status = ?", phone, 1).First(&blackapp).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &blackapp, nil
}

func (r *blackappRepository) CheckByIDCard(ctx context.Context, idCard string) (*entity.blackapp, error) {
	var blackapp entity.blackapp
	err := database.DB.Where("id_card = ? AND status = ?", idCard, 1).First(&blackapp).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &blackapp, nil
}

func (r *blackappRepository) CheckByName(ctx context.Context, name string) (*entity.blackapp, error) {
	var blackapp entity.blackapp
	err := database.DB.Where("name = ? AND status = ?", name, 1).First(&blackapp).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &blackapp, nil
}
