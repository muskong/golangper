package mapper

import (
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"blacklists/domain/constants"
	"blacklists/domain/entity"
)

type BlacklistRepository struct {
	db *gorm.DB
}

func NewBlacklistRepository(db *gorm.DB) *BlacklistRepository {
	return &BlacklistRepository{db: db}
}

func (r *BlacklistRepository) Create(ctx *gin.Context, blacklist *entity.Blacklist) error {
	return r.db.Create(blacklist).Error
}

func (r *BlacklistRepository) Update(ctx *gin.Context, blacklist *entity.Blacklist) error {
	return r.db.Save(blacklist).Error
}

func (r *BlacklistRepository) Delete(ctx *gin.Context, id int) error {
	return r.db.Delete(&entity.Blacklist{}, id).Error
}

func (r *BlacklistRepository) FindByID(ctx *gin.Context, id int) (*entity.Blacklist, error) {
	var blacklist entity.Blacklist
	err := r.db.First(&blacklist, id).Error
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

	err := r.db.Model(&entity.Blacklist{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.Offset(offset).Limit(size).Find(&blacklists).Error
	return blacklists, total, err
}

func (r *BlacklistRepository) UpdateStatus(ctx *gin.Context, id int, status int) error {
	return r.db.Model(&entity.Blacklist{}).Where("id = ?", id).Update("status", status).Error
}

func (r *BlacklistRepository) CheckByPhone(ctx *gin.Context, phone string) (*entity.Blacklist, error) {
	var blacklist entity.Blacklist
	err := r.db.Where("phone = ? AND status = ?", phone, constants.BlacklistStatusApproved).First(&blacklist).Error
	if err != nil {
		return nil, err
	}
	return &blacklist, nil
}

func (r *BlacklistRepository) CheckByIDCard(ctx *gin.Context, idCard string) (*entity.Blacklist, error) {
	var blacklist entity.Blacklist
	err := r.db.Where("id_card = ? AND status = ?", idCard, constants.BlacklistStatusApproved).First(&blacklist).Error
	if err != nil {
		return nil, err
	}
	return &blacklist, nil
}

func (r *BlacklistRepository) CheckByName(ctx *gin.Context, name string) (*entity.Blacklist, error) {
	var blacklist entity.Blacklist
	err := r.db.Where("name = ? AND status = ?", name, constants.BlacklistStatusApproved).First(&blacklist).Error
	if err != nil {
		return nil, err
	}
	return &blacklist, nil
}
