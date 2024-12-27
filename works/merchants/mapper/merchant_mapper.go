package mapper

import (
	"github.com/muskong/gopermission/works/merchants/domain/entity"

	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type merchantRepository struct {
	db *gorm.DB
}

func NewMerchantRepository(db *gorm.DB) *merchantRepository {
	return &merchantRepository{db: db}
}

func (r *merchantRepository) Create(ctx *gin.Context, merchant *entity.Merchant) error {
	return r.db.Create(merchant).Error
}

func (r *merchantRepository) Update(ctx *gin.Context, merchant *entity.Merchant) error {
	return r.db.Save(merchant).Error
}

func (r *merchantRepository) Delete(ctx *gin.Context, id int) error {
	return r.db.Delete(&entity.Merchant{}, id).Error
}

func (r *merchantRepository) FindByID(ctx *gin.Context, id int) (*entity.Merchant, error) {
	var merchant entity.Merchant
	err := r.db.First(&merchant, id).Error
	return &merchant, err
}

func (r *merchantRepository) FindByAPIKey(ctx *gin.Context, apiKey string) (*entity.Merchant, error) {
	var merchant entity.Merchant
	err := r.db.Where("api_key = ?", apiKey).First(&merchant).Error
	return &merchant, err
}

func (r *merchantRepository) List(ctx *gin.Context, page, size int) ([]*entity.Merchant, int64, error) {
	var merchants []*entity.Merchant
	var total int64

	offset := (page - 1) * size

	err := r.db.Model(&entity.Merchant{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.Offset(offset).Limit(size).Find(&merchants).Error
	return merchants, total, err
}

func (r *merchantRepository) UpdateStatus(ctx *gin.Context, id int, status int) error {
	return r.db.Model(&entity.Merchant{}).Where("id = ?", id).Update("status", status).Error
}

func (r *merchantRepository) UpdateToken(ctx *gin.Context, id int, token string, expireTime time.Time) error {
	updates := map[string]interface{}{
		"api_token":         token,
		"token_expire_time": expireTime,
	}
	return r.db.Model(&entity.Merchant{}).Where("id = ?", id).Updates(updates).Error
}
