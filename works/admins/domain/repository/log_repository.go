package repository

import (
	"admins/api/dto"
	"admins/domain/entity"
	"context"
)

type LogRepository interface {
	// 操作日志
	CreateOperationLog(ctx context.Context, log *entity.OperationLog) error
	ListOperationLogs(ctx context.Context, query dto.LogQueryRequest) ([]*entity.OperationLog, int64, error)
	DeleteOperationLog(ctx context.Context, logID int) error
	ClearOperationLogs(ctx context.Context) error

	// 登录日志
	CreateLoginLog(ctx context.Context, log *entity.LoginLog) error
	ListLoginLogs(ctx context.Context, query dto.LogQueryRequest) ([]*entity.LoginLog, int64, error)
	DeleteLoginLog(ctx context.Context, logID int) error
	ClearLoginLogs(ctx context.Context) error
}
