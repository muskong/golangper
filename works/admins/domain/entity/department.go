package entity

import (
	"time"
)

// Department represents a department in the system
type Department struct {
	DepartmentID          int       `gorm:"column:department_id;primaryKey;autoIncrement"`
	ParentID              int       `gorm:"column:parent_id;default:0"`
	DepartmentName        string    `gorm:"column:department_name;size:50;not null"`
	DepartmentCode        string    `gorm:"column:department_code;size:50;not null;uniqueIndex"`
	DepartmentDescription string    `gorm:"column:department_description;size:200"`
	DepartmentLeader      string    `gorm:"column:department_leader;size:50"`
	DepartmentPhone       string    `gorm:"column:department_phone;size:20"`
	DepartmentEmail       string    `gorm:"column:department_email;size:100"`
	DepartmentSort        int       `gorm:"column:department_sort;default:0"`
	DepartmentStatus      int8      `gorm:"column:department_status;default:1"`
	CreatedAt             time.Time `gorm:"column:created_at;not null"`
	UpdatedAt             time.Time `gorm:"column:updated_at"`
	DeletedAt             time.Time `gorm:"column:deleted_at;index"`
}

// TableName returns the table name for the Department model
func (Department) TableName() string {
	return "sys_departments"
}
