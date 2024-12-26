package entity

import "time"

// LoginLog represents a user login log
type LoginLog struct {
	LogID         int       `gorm:"column:log_id;primaryKey;autoIncrement"`
	UserID        int       `gorm:"column:user_id;not null"`
	UserName      string    `gorm:"column:user_name;size:50"`
	LoginIP       string    `gorm:"column:login_ip;size:50"`
	LoginLocation string    `gorm:"column:login_location;size:100"`
	LoginBrowser  string    `gorm:"column:login_browser;size:50"`
	LoginOS       string    `gorm:"column:login_os;size:50"`
	LoginStatus   int8      `gorm:"column:login_status"`
	LoginMessage  string    `gorm:"column:login_message;size:200"`
	CreatedAt     time.Time `gorm:"column:created_at;not null"`
}

// TableName returns the table name for the LoginLog model
func (LoginLog) TableName() string {
	return "sys_login_logs"
}
