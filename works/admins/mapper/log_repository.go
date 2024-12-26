package mapper

import (
	"admins/domain/entity"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type logRepository struct {
	db *gorm.DB
}

func NewLogRepository(db *gorm.DB) *logRepository {
	return &logRepository{db: db}
}

func (r *logRepository) CreateOperationLog(ctx *gin.Context, log *entity.OperationLog) error {
	return r.db.Create(log).Error
}

func (r *logRepository) ListOperationLogs(ctx *gin.Context, page, pageSize int) ([]*entity.OperationLog, int64, error) {
	var logs []*entity.OperationLog
	var total int64
	if err := r.db.Model(&entity.OperationLog{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&logs).Error; err != nil {
		return nil, 0, err
	}

	return logs, total, nil
}
