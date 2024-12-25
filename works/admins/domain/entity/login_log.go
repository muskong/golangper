package entity

import "time"

// LoginLog represents a user login log
type LoginLog struct {
	LoginID     int64     `gorm:"primaryKey;column:login_id;autoIncrement"`
	UserID      int       `gorm:"column:user_id"`
	UserName    string    `gorm:"column:user_name;size:50"`
	UserEmail   string    `gorm:"column:user_email;size:100"`
	LoginType   string    `gorm:"column:login_type;size:10"`
	LoginIP     string    `gorm:"column:login_ip;size:50"`
	LoginStatus string    `gorm:"column:login_status"`
	LoginAgent  string    `gorm:"column:login_agent;size:255"`
	CreatedAt   time.Time `gorm:"column:created_at"`
}

// TableName returns the table name for the LoginLog model
func (LoginLog) TableName() string {
	return "sys_login_logs"
}
