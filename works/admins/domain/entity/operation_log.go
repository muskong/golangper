package entity

import "time"

// OperationLog represents a system operation log
type OperationLog struct {
	LogID             int       `gorm:"column:log_id;primaryKey;autoIncrement"`
	AdminID           int       `gorm:"column:admin_id;not null"`
	AdminName         string    `gorm:"column:admin_name;size:50"`
	OperationIP       string    `gorm:"column:operation_ip;size:50"`
	OperationLocation string    `gorm:"column:operation_location;size:100"`
	OperationBrowser  string    `gorm:"column:operation_browser;size:50"`
	OperationOS       string    `gorm:"column:operation_os;size:50"`
	OperationMethod   string    `gorm:"column:operation_method;size:10"`
	OperationPath     string    `gorm:"column:operation_path;size:200"`
	OperationModule   string    `gorm:"column:operation_module;size:50"`
	OperationContent  string    `gorm:"column:operation_content;size:1000"`
	OperationStatus   int       `gorm:"column:operation_status"`
	OperationLatency  int       `gorm:"column:operation_latency"` // 请求耗时(ms)
	OperationRequest  string    `gorm:"column:operation_request;type:text"`
	OperationResponse string    `gorm:"column:operation_response;type:text"`
	CreatedAt         time.Time `gorm:"column:created_at;not null"`
}

// TableName returns the table name for the OperationLog model
func (OperationLog) TableName() string {
	return "sys_operation_logs"
}
