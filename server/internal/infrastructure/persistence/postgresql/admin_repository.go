package mysql

import (
	"blacklist/internal/domain/entity"

	"gorm.io/gorm"
)

type AdminRepositoryImpl struct {
	db *gorm.DB
}

func (r *AdminRepositoryImpl) Create(admin *entity.Admin) error {
	// 实现用户创建的数据库操作
	return r.db.Create(admin).Error
}

func (r *AdminRepositoryImpl) FindByID(id string) (*entity.Admin, error) {
	var admin entity.Admin
	err := r.db.Where("id = ?", id).First(&admin).Error
	return &admin, err
}
