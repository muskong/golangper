package repository

import (
	"github.com/muskong/gopermission/works/merchants/domain/entity"

	"github.com/gin-gonic/gin"
)

type LoginLogRepository interface {
	Create(ctx *gin.Context, log *entity.LoginLog) error
	List(ctx *gin.Context, userType int, page, size int) ([]*entity.LoginLog, int64, error)
}
