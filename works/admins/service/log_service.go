package service

import (
	"github.com/muskong/gopermission/works/admins/service/dto"

	"github.com/gin-gonic/gin"
)

type LogService interface {
	// 操作日志
	CreateOperationLog(ctx *gin.Context, log *dto.OperationLogCreateDTO) error
	ListOperationLogs(ctx *gin.Context, page, pageSize int) ([]*dto.OperationLogInfo, int64, error)
	DeleteOperationLog(ctx *gin.Context, logID int) error
	ClearOperationLogs(ctx *gin.Context) error
}
