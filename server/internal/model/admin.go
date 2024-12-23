package model

import (
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	Username  string         `json:"username" gorm:"size:50;uniqueIndex;not null;comment:用户名"`
	Password  string         `json:"-" gorm:"size:100;not null;comment:密码"`
	Name      string         `json:"name" gorm:"size:50;not null;comment:姓名"`
	Phone     string         `json:"phone" gorm:"size:20;comment:手机号"`
	Email     string         `json:"email" gorm:"size:100;comment:邮箱"`
	Status    int            `json:"status" gorm:"type:smallint;default:1;comment:状态 1:正常 2:禁用"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}
