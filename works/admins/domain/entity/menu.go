package entity

import "time"

// Menu represents a menu item in the system
type Menu struct {
	MenuID         int       `gorm:"column:menu_id;primaryKey;autoIncrement"`
	ParentID       int       `gorm:"column:parent_id;default:0"`
	MenuName       string    `gorm:"column:menu_name;size:50;not null"`
	MenuCode       string    `gorm:"column:menu_code;size:50;not null;uniqueIndex"`
	MenuType       int8      `gorm:"column:menu_type;not null"` // 1:目录 2:菜单 3:按钮
	MenuPath       string    `gorm:"column:menu_path;size:200"`
	MenuComponent  string    `gorm:"column:menu_component;size:200"`
	MenuPermission string    `gorm:"column:menu_permission;size:100"`
	MenuIcon       string    `gorm:"column:menu_icon;size:100"`
	MenuSort       int       `gorm:"column:menu_sort;default:0"`
	MenuStatus     int8      `gorm:"column:menu_status;default:1"`
	MenuVisible    int8      `gorm:"column:menu_visible;default:1"`
	MenuCache      int8      `gorm:"column:menu_cache;default:0"`
	CreatedAt      time.Time `gorm:"column:created_at;not null"`
	UpdatedAt      time.Time `gorm:"column:updated_at"`
	DeletedAt      time.Time `gorm:"column:deleted_at;index"`
}

// TableName returns the table name for the Menu model
func (Menu) TableName() string {
	return "sys_menus"
}
