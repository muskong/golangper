package entity

// OperationLog represents a system operation log
type OperationLog struct {
	OperationID       int64  `gorm:"primaryKey;column:operation_id;autoIncrement"`
	UserID            int    `gorm:"column:user_id"`
	OperationIP       string `gorm:"column:operation_ip;size:50"`
	OperationMethod   string `gorm:"column:operation_method;size:10"`
	OperationPath     string `gorm:"column:operation_path;size:255"`
	OperationStatus   int    `gorm:"column:operation_status"`
	OperationLatency  int    `gorm:"column:operation_latency"`
	OperationAgent    string `gorm:"column:operation_agent;size:255"`
	OperationRequest  string `gorm:"column:operation_request;type:text"`
	OperationResponse string `gorm:"column:operation_response;type:text"`
	BaseModel
}

// TableName returns the table name for the OperationLog model
func (OperationLog) TableName() string {
	return "sys_operation_logs"
}
