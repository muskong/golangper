package repository

import (
	"blackapp/internal/domain/entity"

	"github.com/gin-gonic/gin"
)

type BlacklistRepository interface {
	Create(ctx *gin.Context, blacklist *entity.Blacklist) error
	Update(ctx *gin.Context, blacklist *entity.Blacklist) error
	Delete(ctx *gin.Context, id int) error
	FindByID(ctx *gin.Context, id int) (*entity.Blacklist, error)
	List(ctx *gin.Context, page, size int) ([]*entity.Blacklist, int64, error)
	UpdateStatus(ctx *gin.Context, id int, status int) error

	// 验证用户是否在黑名单中
	CheckByPhone(ctx *gin.Context, phone string) (*entity.Blacklist, error)
	CheckByIDCard(ctx *gin.Context, idCard string) (*entity.Blacklist, error)
	CheckByName(ctx *gin.Context, name string) (*entity.Blacklist, error)
}
