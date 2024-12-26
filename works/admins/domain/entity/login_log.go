package entity

import "time"

// LoginLog represents a user login log
type LoginLog struct {
	LogID         int       `gorm:"column:logID;primaryKey;autoIncrement"`
	AdminID       int       `gorm:"column:adminID;not null"`
	AdminName     string    `gorm:"column:adminName;size:50"`
	LoginIP       string    `gorm:"column:loginIP;size:50"`
	LoginLocation string    `gorm:"column:loginLocation;size:100"`
	LoginBrowser  string    `gorm:"column:loginBrowser;size:50"`
	LoginOS       string    `gorm:"column:loginOS;size:50"`
	LoginStatus   int8      `gorm:"column:loginStatus"`
	LoginMessage  string    `gorm:"column:loginMessage;size:200"`
	CreatedAt     time.Time `gorm:"column:createdAt;not null"`
}

// TableName returns the table name for the LoginLog model
func (LoginLog) TableName() string {
	return "sys_login_logs"
}
