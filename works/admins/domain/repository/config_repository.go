package repository

import (
	"github.com/muskong/gopermission/works/admins/domain/entity"

	"github.com/gin-gonic/gin"
)

type ConfigRepository interface {
	Create(ctx *gin.Context, config *entity.Config) error
	Update(ctx *gin.Context, config *entity.Config) error
	Delete(ctx *gin.Context, configID int) error
	FindByID(ctx *gin.Context, configID int) (*entity.Config, error)
	FindByKey(ctx *gin.Context, configKey string) (*entity.Config, error)
	List(ctx *gin.Context, offset, limit int) ([]*entity.Config, int64, error)
}
