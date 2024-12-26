package mapper

import (
	"blacklists/domain/entity"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type queryLogRepository struct {
	db *gorm.DB
}

func NewQueryLogRepository(db *gorm.DB) *queryLogRepository {
	return &queryLogRepository{db: db}
}

func (r *queryLogRepository) Create(ctx *gin.Context, log *entity.QueryLog) error {
	return r.db.Create(log).Error
}

func (r *queryLogRepository) List(ctx *gin.Context, merchantID int, page, size int) ([]*entity.QueryLog, int64, error) {
	var logs []*entity.QueryLog
	var total int64

	query := r.db.Model(&entity.QueryLog{})
	if merchantID > 0 {
		query = query.Where("merchant_id = ?", merchantID)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * size
	err = query.Order("created_at DESC").
		Offset(offset).
		Limit(size).
		Find(&logs).Error

	return logs, total, err
}
