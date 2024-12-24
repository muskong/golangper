package repository

import (
	"blackapp/internal/domain/entity"

	"github.com/gin-gonic/gin"
)

type AdminRepository interface {
	Create(ctx *gin.Context, admin *entity.Admin) error
	Update(ctx *gin.Context, admin *entity.Admin) error
	Delete(ctx *gin.Context, id int) error
	FindByID(ctx *gin.Context, id int) (*entity.Admin, error)
	FindByUsername(ctx *gin.Context, username string) (*entity.Admin, error)
	List(ctx *gin.Context, page, size int) ([]*entity.Admin, int64, error)
	UpdateStatus(ctx *gin.Context, id int, status int) error
}
