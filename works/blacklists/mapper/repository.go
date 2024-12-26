package mapper

import (
	"gorm.io/gorm"
)

func NewRepository(db *gorm.DB) (*queryLogRepository, *blacklistRepository) {
	return NewQueryLogRepository(db), NewBlacklistRepository(db)
}
