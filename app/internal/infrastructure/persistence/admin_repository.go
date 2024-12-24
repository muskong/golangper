package persistence

import (
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"blackapp/internal/domain/entity"
	"blackapp/pkg/database"
)

type AdminRepository struct{}

func NewAdminRepository() *AdminRepository {
	return &AdminRepository{}
}

func (r *AdminRepository) Create(ctx *gin.Context, admin *entity.Admin) error {
	return database.DB.Create(admin).Error
}

func (r *AdminRepository) Update(ctx *gin.Context, admin *entity.Admin) error {
	return database.DB.Save(admin).Error
}

func (r *AdminRepository) Delete(ctx *gin.Context, id int) error {
	return database.DB.Delete(&entity.Admin{}, id).Error
}

func (r *AdminRepository) FindByID(ctx *gin.Context, id int) (*entity.Admin, error) {
	var admin entity.Admin
	err := database.DB.First(&admin, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &admin, nil
}

func (r *AdminRepository) FindByUsername(ctx *gin.Context, username string) (*entity.Admin, error) {
	var admin entity.Admin
	err := database.DB.Where("username = ?", username).First(&admin).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &admin, nil
}

func (r *AdminRepository) List(ctx *gin.Context, page, size int) ([]*entity.Admin, int64, error) {
	var admins []*entity.Admin
	var total int64

	offset := (page - 1) * size

	err := database.DB.Model(&entity.Admin{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = database.DB.Offset(offset).Limit(size).Find(&admins).Error
	return admins, total, err
}

func (r *AdminRepository) UpdateStatus(ctx *gin.Context, id int, status int) error {
	return database.DB.Model(&entity.Admin{}).Where("id = ?", id).Update("status", status).Error
}
