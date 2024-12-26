package mapper

import (
	"merchants/domain/entity"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginLogRepository struct {
	db *gorm.DB
}

func NewLoginLogRepository(db *gorm.DB) *LoginLogRepository {
	return &LoginLogRepository{db: db}
}

func (r *LoginLogRepository) Create(ctx *gin.Context, log *entity.LoginLog) error {
	return r.db.Create(log).Error
}

func (r *LoginLogRepository) List(ctx *gin.Context, userType int, page, size int) ([]*entity.LoginLog, int64, error) {
	var logs []*entity.LoginLog
	var total int64

	query := r.db.Model(&entity.LoginLog{})
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
