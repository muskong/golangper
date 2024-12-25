package entity

// Menu represents a menu item in the system
type Menu struct {
	MenuID          int    `gorm:"primaryKey;column:menu_id;autoIncrement"`
	MenuName        string `gorm:"column:menu_name;size:100;not null"`
	ParentID        int    `gorm:"column:parent_id"`
	MenuPath        string `gorm:"column:menu_path;size:255"`
	MenuComponent   string `gorm:"column:menu_component;size:255"`
	MenuSort        int    `gorm:"column:menu_sort;default:0"`
	MenuType        string `gorm:"column:menu_type;size:1"`
	MenuStatus      int8   `gorm:"column:menu_status;default:1"`
	MenuVisible     bool   `gorm:"column:menu_visible;default:true"`
	MenuPermissions string `gorm:"column:menu_permissions;size:100"`
	MenuIcon        string `gorm:"column:menu_icon;size:100;default:'#'"`
	Children        []Menu `gorm:"-"`
	BaseModel
}

// TableName returns the table name for the Menu model
func (Menu) TableName() string {
	return "sys_menus"
}
