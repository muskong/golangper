package persistence

import (
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"blackapp/internal/domain/constants"
	"blackapp/internal/domain/entity"
	"blackapp/pkg/database"
)

type BlacklistRepository struct{}

func NewBlacklistRepository() *BlacklistRepository {
	return &BlacklistRepository{}
}

func (r *BlacklistRepository) Create(ctx *gin.Context, blacklist *entity.Blacklist) error {
	return database.DB.Create(blacklist).Error
}

func (r *BlacklistRepository) Update(ctx *gin.Context, blacklist *entity.Blacklist) error {
	return database.DB.Save(blacklist).Error
}

func (r *BlacklistRepository) Delete(ctx *gin.Context, id int) error {
	return database.DB.Delete(&entity.Blacklist{}, id).Error
}

func (r *BlacklistRepository) FindByID(ctx *gin.Context, id int) (*entity.Blacklist, error) {
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

func (r *BlacklistRepository) List(ctx *gin.Context, page, size int) ([]*entity.Blacklist, int64, error) {
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

func (r *BlacklistRepository) UpdateStatus(ctx *gin.Context, id int, status int) error {
	return database.DB.Model(&entity.Blacklist{}).Where("id = ?", id).Update("status", status).Error
}

func (r *BlacklistRepository) CheckByPhone(ctx *gin.Context, phone string) (*entity.Blacklist, error) {
	var blacklist entity.Blacklist
	err := database.DB.Where("phone = ? AND status = ?", phone, constants.BlacklistStatusApproved).First(&blacklist).Error
	if err != nil {
		return nil, err
	}
	return &blacklist, nil
}

func (r *BlacklistRepository) CheckByIDCard(ctx *gin.Context, idCard string) (*entity.Blacklist, error) {
	var blacklist entity.Blacklist
	err := database.DB.Where("id_card = ? AND status = ?", idCard, constants.BlacklistStatusApproved).First(&blacklist).Error
	if err != nil {
		return nil, err
	}
	return &blacklist, nil
}

func (r *BlacklistRepository) CheckByName(ctx *gin.Context, name string) (*entity.Blacklist, error) {
	var blacklist entity.Blacklist
	err := database.DB.Where("name = ? AND status = ?", name, constants.BlacklistStatusApproved).First(&blacklist).Error
	if err != nil {
		return nil, err
	}
	return &blacklist, nil
}
