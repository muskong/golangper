package entity

// Job represents a scheduled job in the system
type Job struct {
	JobID             int    `gorm:"primaryKey;column:job_id;autoIncrement"`
	JobName           string `gorm:"column:job_name;size:100;not null"`
	JobCronExpression string `gorm:"column:job_cron_expression;size:100;not null"`
	JobCommand        string `gorm:"column:job_command;size:255;not null"`
	JobStatus         int8   `gorm:"column:job_status;default:1"`
	JobDescription    string `gorm:"column:job_description;type:text"`
	BaseModel
}

// TableName returns the table name for the Job model
func (Job) TableName() string {
	return "sys_jobs"
}

type JobLog struct {
	JobLogID      int64  `gorm:"primaryKey;column:job_log_id;autoIncrement"`
	JobName       string `gorm:"column:job_name;size:64;not null"`
	JobGroup      string `gorm:"column:job_group;size:64;not null"`
	JobTarget     string `gorm:"column:job_target;size:500;not null"`
	JobMessage    string `gorm:"column:job_message;size:500"`
	Status        string `gorm:"column:status;size:1;default:'0'"`
	ExceptionInfo string `gorm:"column:exception_info;size:2000"`
	BaseModel
}

// TableName returns the table name for the JobLog model
func (JobLog) TableName() string {
	return "sys_job_logs"
}
