package entity

import "time"

type Merchant struct {
	ID              uint   `gorm:"primarykey"`
	Name            string `gorm:"type:varchar(100);not null"`
	Address         string `gorm:"type:varchar(255)"`
	ContactPerson   string `gorm:"type:varchar(50)"`
	ContactPhone    string `gorm:"type:varchar(20)"`
	Remark          string `gorm:"type:text"`
	Status          int    `gorm:"type:int;default:1"`
	IPWhitelist     string `gorm:"type:text"`
	APIKey          string `gorm:"type:varchar(64);uniqueIndex"`
	APISecret       string `gorm:"type:varchar(64)"`
	APIToken        string `gorm:"type:varchar(255)"`
	TokenExpireTime time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time `gorm:"index"`
}
