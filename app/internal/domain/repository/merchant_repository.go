package repository

import (
	"blackapp/internal/domain/entity"
	"time"

	"github.com/gin-gonic/gin"
)

type MerchantRepository interface {
	Create(ctx *gin.Context, merchant *entity.Merchant) error
	Update(ctx *gin.Context, merchant *entity.Merchant) error
	Delete(ctx *gin.Context, id int) error
	FindByID(ctx *gin.Context, id int) (*entity.Merchant, error)
	FindByAPIKey(ctx *gin.Context, apiKey string) (*entity.Merchant, error)
	List(ctx *gin.Context, page, size int) ([]*entity.Merchant, int64, error)
	UpdateStatus(ctx *gin.Context, id int, status int) error
	UpdateToken(ctx *gin.Context, id int, token string, expireTime time.Time) error
}
