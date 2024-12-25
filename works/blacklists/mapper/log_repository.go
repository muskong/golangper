package mapper

import (
	"blacklists/domain/entity"
	"pkgs/database"

	"github.com/gin-gonic/gin"
)

type QueryLogRepository struct{}

func NewQueryLogRepository() *QueryLogRepository {
	return &QueryLogRepository{}
}

func (r *QueryLogRepository) Create(ctx *gin.Context, log *entity.QueryLog) error {
	return database.DB.Create(log).Error
}

func (r *QueryLogRepository) List(ctx *gin.Context, merchantID int, page, size int) ([]*entity.QueryLog, int64, error) {
	var logs []*entity.QueryLog
	var total int64

	query := database.DB.Model(&entity.QueryLog{})
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
