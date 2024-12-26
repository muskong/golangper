package entity

import (
	"time"
)

type Admin struct {
	AdminID       int       `gorm:"column:admin_id;primaryKey;autoIncrement"`
	DepartmentID  *int      `gorm:"column:department_id"`
	AdminName     string    `gorm:"column:admin_name;size:50;not null;uniqueIndex"`
	AdminEmail    string    `gorm:"column:admin_email;size:100;not null"`
	AdminPhone    string    `gorm:"column:admin_phone;size:20;not null"`
	AdminSex      int8      `gorm:"column:admin_sex;default:0"`
	AdminAvatar   string    `gorm:"column:admin_avatar;size:255"`
	AdminStatus   int       `gorm:"column:admin_status;default:1"`
	AdminPassword string    `gorm:"column:admin_password;size:100;not null"`
	LastLogin     time.Time `gorm:"column:last_login"`
	CreatedAt     time.Time `gorm:"column:created_at;not null"`
	UpdatedAt     time.Time `gorm:"column:updated_at"`
	DeletedAt     time.Time `gorm:"column:deleted_at;index"`

	// 关联
	Roles      []Role      `gorm:"many2many:admin_roles"`
	Posts      []Post      `gorm:"many2many:admin_posts"`
	Department *Department `gorm:"foreignKey:DepartmentID"`
}

func (Admin) TableName() string {
	return "sys_admins"
}

type AdminRole struct {
	AdminID int `gorm:"primaryKey;column:admin_id"`
	RoleID  int `gorm:"primaryKey;column:role_id"`
}

func (AdminRole) TableName() string {
	return "sys_admin_roles"
}

type AdminPost struct {
	AdminID int `gorm:"primaryKey;column:admin_id"`
	PostID  int `gorm:"primaryKey;column:post_id"`
}

func (AdminPost) TableName() string {
	return "sys_admin_posts"
}
