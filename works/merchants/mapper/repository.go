package mapper

import (
	"gorm.io/gorm"
)

func NewRepository(db *gorm.DB) (*merchantRepository, *loginLogRepository) {
	return NewMerchantRepository(db), NewLoginLogRepository(db)
}
