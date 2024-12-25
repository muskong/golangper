package entity

import (
	"time"
)

// Department represents a department in the system
type Department struct {
	DepartmentID          int          `gorm:"primaryKey;column:department_id;autoIncrement"`
	ParentID              int          `gorm:"column:parent_id"`
	DepartmentName        string       `gorm:"column:department_name;size:100;not null"`
	DepartmentCode        string       `gorm:"column:department_code;size:50;not null"`
	DepartmentDescription string       `gorm:"column:department_description;type:text"`
	DepartmentLeader      string       `gorm:"column:department_leader;size:100"`
	DepartmentPhone       string       `gorm:"column:department_phone;size:20"`
	DepartmentEmail       string       `gorm:"column:department_email;size:100"`
	DepartmentSort        int          `gorm:"column:department_sort;default:0"`
	DepartmentStatus      int8         `gorm:"column:department_status;default:1"`
	Children              []Department `gorm:"-"`
	CreatedAt             time.Time
}

// TableName returns the table name for the Department model
func (Department) TableName() string {
	return "sys_departments"
}
