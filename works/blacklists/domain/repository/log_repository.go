package repository

import (
	"github.com/muskong/gopermission/works/blacklists/domain/entity"

	"github.com/gin-gonic/gin"
)

type QueryLogRepository interface {
	Create(ctx *gin.Context, log *entity.QueryLog) error
	List(ctx *gin.Context, merchantID int, page, size int) ([]*entity.QueryLog, int64, error)
}
