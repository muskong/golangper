package entity

import "time"

type LoginLog struct {
	ID        int    `gorm:"primarykey"`
	Type      int    `gorm:"type:int;not null"` // constants.UserTypeMerchant or constants.UserTypeAdmin
	UserID    int    `gorm:"index"`
	IP        string `gorm:"type:varchar(50)"`
	UserAgent string `gorm:"type:varchar(255)"`
	Status    int    `gorm:"type:int;not null"` // constants.LogStatusSuccess or constants.LogStatusFailed
	CreatedAt time.Time
}
