package entity

import "time"

type Admin struct {
	ID        int    `gorm:"primarykey"`
	Username  string `gorm:"type:varchar(50);uniqueIndex"`
	Password  string `gorm:"type:varchar(100)"`
	Name      string `gorm:"type:varchar(50)"`
	Status    int    `gorm:"type:int;default:1"` // constants.StatusEnabled
	LastLogin time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}
