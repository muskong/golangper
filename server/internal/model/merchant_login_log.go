package model

import (
	"time"
)

// MerchantLoginLog 商户登录日志
type MerchantLoginLog struct {
	ID         uint      `json:"id" gorm:"primarykey"`
	MerchantID uint      `json:"merchant_id" gorm:"index;not null;comment:商户ID"`
	Merchant   Merchant  `json:"merchant" gorm:"foreignKey:MerchantID"`
	IP         string    `json:"ip" gorm:"type:varchar(50);not null;comment:登录IP"`
	UserAgent  string    `json:"user_agent" gorm:"type:varchar(255);comment:用户代理"`
	LoginTime  time.Time `json:"login_time" gorm:"not null;comment:登录时间"`
	Status     int       `json:"status" gorm:"type:smallint;not null;comment:登录状态 1:成功 2:失败"`
	Remark     string    `json:"remark" gorm:"type:varchar(255);comment:备注信息"`
}

// TableName 指定表名
func (MerchantLoginLog) TableName() string {
	return "merchant_login_logs"
}

// LoginStatus 登录状态常量
const (
	LoginStatusSuccess = 1 // 登录成功
	LoginStatusFailed  = 2 // 登录失败
)