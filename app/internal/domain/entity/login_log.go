package entity

import "time"

type LoginLog struct {
	ID        int    `gorm:"primarykey"`
	Type      int    `gorm:"type:int"` // 1:商户 2:管理员
	UserID    int    `gorm:"index"`
	IP        string `gorm:"type:varchar(50)"`
	UserAgent string `gorm:"type:varchar(255)"`
	Status    int    `gorm:"type:int"` // 1:成功 2:失败
	CreatedAt time.Time
}
