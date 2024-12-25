package mapper

import (
	"merchants/domain/entity"
	"pkgs/database"

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
