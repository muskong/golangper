package model

import (
	"time"

	"gorm.io/gorm"
)

// Merchant 商户模型
type Merchant struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	Name      string         `json:"name" gorm:"type:varchar(100);not null;comment:商户名称"`
	Address   string         `json:"address" gorm:"type:varchar(255);comment:商户地址"`
	Contact   string         `json:"contact" gorm:"type:varchar(50);comment:商户联系人"`
	Phone     string         `json:"phone" gorm:"type:varchar(20);comment:商户联系电话"`
	Remark    string         `json:"remark" gorm:"type:text;comment:商户备注"`
	Status    int            `json:"status" gorm:"type:smallint;default:1;comment:商户状态 1:正常 2:禁用"`
	APIKey    string         `json:"api_key" gorm:"type:varchar(32);uniqueIndex;comment:商户API Key"`
	APISecret string         `json:"-" gorm:"type:varchar(64);comment:商户API Secret"` // 不返回给前端
	Token     string         `json:"-" gorm:"type:varchar(255);comment:商户API Token"`
	TokenExp  time.Time      `json:"token_exp" gorm:"comment:Token过期时间"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
	LoginLogs []MerchantLoginLog `json:"-" gorm:"foreignKey:MerchantID"`
}

// TableName 指定表名
func (Merchant) TableName() string {
	return "merchants"
}

// MerchantStatus 商户状态
const (
	MerchantStatusNormal   = 1 // 正常
	MerchantStatusDisabled = 2 // 禁用
)
