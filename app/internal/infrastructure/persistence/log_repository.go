package persistence

import (
	"blackapp/internal/domain/entity"
	"blackapp/pkg/database"

	"github.com/gin-gonic/gin"
)

type LoginLogRepository struct{}

func NewLoginLogRepository() *LoginLogRepository {
	return &LoginLogRepository{}
}

func (r *LoginLogRepository) Create(ctx *gin.Context, log *entity.LoginLog) error {
	return database.DB.Create(log).Error
}

func (r *LoginLogRepository) List(ctx *gin.Context, userType int, page, size int) ([]*entity.LoginLog, int64, error) {
	var logs []*entity.LoginLog
	var total int64

	query := database.DB.Model(&entity.LoginLog{})
	if userType > 0 {
		query = query.Where("type = ?", userType)
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
