package entity

import (
	"time"
)

type Admin struct {
	AdminID       int    `gorm:"primaryKey;column:admin_id;autoIncrement"`
	DepartmentID  *int   `gorm:"column:department_id"`
	AdminName     string `gorm:"column:admin_name;size:100;not null;unique"`
	AdminEmail    string `gorm:"column:admin_email;size:100;unique"`
	AdminPhone    string `gorm:"column:admin_phone;size:20"`
	AdminSex      int8   `gorm:"column:admin_sex;default:0"`
	AdminAvatar   string `gorm:"column:admin_avatar;size:255"`
	AdminPassword string `gorm:"column:admin_password;size:100;not null"`
	AdminStatus   int    `gorm:"column:admin_status;type:int;default:1"` // constants.StatusEnabled
	LastLogin     time.Time
	Roles         []Role `gorm:"many2many:sys_user_roles;foreignKey:UserID;joinForeignKey:user_id;References:RoleID;joinReferences:role_id"`
	Posts         []Post `gorm:"many2many:sys_user_posts;foreignKey:UserID;joinForeignKey:user_id;References:PostID;joinReferences:post_id"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
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
