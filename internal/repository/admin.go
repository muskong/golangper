package repository

import (
	"blacklist/internal/model"
	"blacklist/internal/pkg/database"
)

type AdminRepository interface {
	Create(admin *model.Admin) error
	Update(admin *model.Admin) error
	Delete(id uint) error
	GetByID(id uint) (*model.Admin, error)
	GetByUsername(username string) (*model.Admin, error)
	List(page, pageSize int) ([]model.Admin, int64, error)
}

type adminRepository struct {
	db *database.PostgresDB
}

func NewAdminRepository(db *database.PostgresDB) AdminRepository {
	return &adminRepository{db: db}
}

func (r *adminRepository) Create(admin *model.Admin) error {
	return r.db.Create(admin).Error
}

func (r *adminRepository) Update(admin *model.Admin) error {
	return r.db.Save(admin).Error
}

func (r *adminRepository) Delete(id uint) error {
	return r.db.Delete(&model.Admin{}, id).Error
}

func (r *adminRepository) GetByID(id uint) (*model.Admin, error) {
	var admin model.Admin
	err := r.db.First(&admin, id).Error
	return &admin, err
}

func (r *adminRepository) GetByUsername(username string) (*model.Admin, error) {
	var admin model.Admin
	err := r.db.Where("username = ?", username).First(&admin).Error
	return &admin, err
}

func (r *adminRepository) List(page, pageSize int) ([]model.Admin, int64, error) {
	var admins []model.Admin
	var total int64

	err := r.db.Model(&model.Admin{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err = r.db.Offset(offset).Limit(pageSize).Find(&admins).Error
	return admins, total, err
}
