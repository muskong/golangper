package model

import (
	"time"

	"gorm.io/gorm"
)

type BlacklistUser struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Name      string         `gorm:"size:100;not null" json:"name"`
	Phone     string         `gorm:"size:20;uniqueIndex" json:"phone"`
	IDCard    string         `gorm:"size:18;uniqueIndex" json:"id_card"`
	Email     string         `gorm:"size:100" json:"email"`
	Address   string         `gorm:"size:200" json:"address"`
	Remark    string         `gorm:"size:500" json:"remark"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
