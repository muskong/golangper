package entity

import "time"

// BaseModel contains common fields for all models
type BaseModel struct {
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
