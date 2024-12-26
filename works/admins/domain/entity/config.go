package entity

import "time"

// Config represents a system configuration
type Config struct {
	ConfigID          int       `gorm:"column:config_id;primaryKey;autoIncrement"`
	ConfigName        string    `gorm:"column:config_name;size:100;not null"`
	ConfigKey         string    `gorm:"column:config_key;size:100;not null;uniqueIndex"`
	ConfigValue       string    `gorm:"column:config_value;size:500;not null"`
	ConfigType        string    `gorm:"column:config_type;size:50;not null"` // text/json/yaml/ini
	ConfigStatus      int8      `gorm:"column:config_status;default:1"`
	ConfigDescription string    `gorm:"column:config_description;size:200"`
	CreatedAt         time.Time `gorm:"column:created_at;not null"`
	UpdatedAt         time.Time `gorm:"column:updated_at"`
	DeletedAt         time.Time `gorm:"column:deleted_at;index"`
}

// TableName returns the table name for the Config model
func (Config) TableName() string {
	return "sys_configs"
}
