package repository

import (
	"admins/domain/entity"

	"github.com/gin-gonic/gin"
)

type LogRepository interface {
	// 操作日志
	CreateOperationLog(ctx *gin.Context, log *entity.OperationLog) error
	ListOperationLogs(ctx *gin.Context, page, pageSize int) ([]*entity.OperationLog, int64, error)
}
