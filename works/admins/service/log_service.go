package service

import (
	"admins/api/dto"
	"context"
)

type LogService interface {
	// 操作日志
	CreateOperationLog(ctx context.Context, log *dto.OperationLogCreateRequest) error
	ListOperationLogs(ctx context.Context, query dto.LogQueryRequest) (*dto.PageResponse, error)
	DeleteOperationLog(ctx context.Context, logID int) error
	ClearOperationLogs(ctx context.Context) error

	// 登录日志
	CreateLoginLog(ctx context.Context, log *dto.LoginLogCreateRequest) error
	ListLoginLogs(ctx context.Context, query dto.LogQueryRequest) (*dto.PageResponse, error)
	DeleteLoginLog(ctx context.Context, logID int) error
	ClearLoginLogs(ctx context.Context) error
}
