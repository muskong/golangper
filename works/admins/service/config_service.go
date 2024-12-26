package service

import (
	"admins/service/dto"

	"github.com/gin-gonic/gin"
)

type ConfigService interface {
	Create(ctx *gin.Context, req dto.ConfigCreateDTO) error
	Update(ctx *gin.Context, req dto.ConfigUpdateDTO) error
	Delete(ctx *gin.Context, configID int) error
	GetByID(ctx *gin.Context, configID int) (*dto.ConfigInfo, error)
	GetByKey(ctx *gin.Context, configKey string) (*dto.ConfigInfo, error)
	List(ctx *gin.Context, page, pageSize int) ([]*dto.ConfigInfo, int64, error)
	RefreshCache(ctx *gin.Context) error
}
