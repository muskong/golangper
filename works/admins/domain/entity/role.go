package entity

// Role represents a role in the system
type Role struct {
	RoleID          int    `gorm:"primaryKey;column:role_id;autoIncrement"`
	RoleName        string `gorm:"column:role_name;size:100;not null;unique"`
	RoleCode        string `gorm:"column:role_code;size:50;not null;unique"`
	RoleDescription string `gorm:"column:role_description;type:text"`
	RoleStatus      int8   `gorm:"column:role_status;default:1"`
	BaseModel
}

// TableName returns the table name for the Role model
func (Role) TableName() string {
	return "sys_roles"
}

// RoleMenu represents the many-to-many relationship between roles and menus
type RoleMenu struct {
	RoleID int `gorm:"primaryKey;column:role_id"`
	MenuID int `gorm:"primaryKey;column:menu_id"`
}

// TableName returns the table name for the RoleMenu model
func (RoleMenu) TableName() string {
	return "sys_role_menus"
}

// RoleDepartment represents the many-to-many relationship between roles and departments
type RoleDepartment struct {
	RoleID       int `gorm:"primaryKey;column:role_id"`
	DepartmentID int `gorm:"primaryKey;column:department_id"`
}

// TableName returns the table name for the RoleDepartment model
func (RoleDepartment) TableName() string {
	return "sys_role_departments"
}
