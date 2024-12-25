package entity

import "time"

type QueryLog struct {
	ID         int    `gorm:"primarykey"`
	MerchantID int    `gorm:"index"`
	Phone      string `gorm:"type:varchar(20)"`
	IDCard     string `gorm:"type:varchar(18)"`
	Name       string `gorm:"type:varchar(50)"`
	IP         string `gorm:"type:varchar(50)"`
	UserAgent  string `gorm:"type:varchar(255)"`
	Exists     bool   `gorm:"type:boolean"`
	CreatedAt  time.Time
}
