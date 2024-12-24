package service

import (
	"blackapp/internal/service/dto"

	"github.com/gin-gonic/gin"
)

type MerchantService interface {
	Create(ctx *gin.Context, req *dto.CreateMerchantDTO) error
	Update(ctx *gin.Context, req *dto.UpdateMerchantDTO) error
	Delete(ctx *gin.Context, id int) error
	GetByID(ctx *gin.Context, id int) (*dto.MerchantDTO, error)
	List(ctx *gin.Context, page, size int) ([]*dto.MerchantDTO, int64, error)
	UpdateStatus(ctx *gin.Context, id int, status int) error
	GenerateAPICredentials(ctx *gin.Context, id int) error
	Login(ctx *gin.Context, req *dto.MerchantLoginDTO) (string, error)
}
