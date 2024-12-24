package repository

import (
	"blackapp/internal/domain/entity"

	"github.com/gin-gonic/gin"
)

type LoginLogRepository interface {
	Create(ctx *gin.Context, log *entity.LoginLog) error
	List(ctx *gin.Context, userType int, page, size int) ([]*entity.LoginLog, int64, error)
}

type QueryLogRepository interface {
	Create(ctx *gin.Context, log *entity.QueryLog) error
	List(ctx *gin.Context, merchantID int, page, size int) ([]*entity.QueryLog, int64, error)
}
