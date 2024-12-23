package repository

import "blacklist/internal/domain/entity"

type AdminRepository interface {
	Create(admin *entity.Admin) error
	FindByID(id string) (*entity.Admin, error)
	// ... 其他仓储接口方法
}
