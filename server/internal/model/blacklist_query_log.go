package model

import (
	"time"
)

// BlacklistQueryLog 黑名单查询日志
type BlacklistQueryLog struct {
	ID         uint      `json:"id" gorm:"primarykey"`
	MerchantID uint      `json:"merchant_id" gorm:"index;not null;comment:商户ID"`
	Merchant   Merchant  `json:"merchant" gorm:"foreignKey:MerchantID"`
	Phone      string    `json:"phone" gorm:"type:varchar(20);not null;comment:查询的手机号"`
	QueryTime  time.Time `json:"query_time" gorm:"not null;index;comment:查询时间"`
	IP         string    `json:"ip" gorm:"type:varchar(50);not null;comment:查询IP"`
	UserAgent  string    `json:"user_agent" gorm:"type:varchar(255);comment:用户代理"`
	Result     bool      `json:"result" gorm:"not null;comment:查询结果(是否存在)"`
}

// TableName 指定表名
func (BlacklistQueryLog) TableName() string {
	return "blacklist_query_logs"
}