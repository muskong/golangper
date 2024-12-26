package mapper

import (
	"gorm.io/gorm"
)

func NewRepository(db *gorm.DB) (*adminRepository, *logRepository, *configRepository) {
	return NewAdminRepository(db), NewLogRepository(db), NewConfigRepository(db)
}
