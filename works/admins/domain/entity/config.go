package entity

// Config represents a system configuration
type Config struct {
	ConfigID    int    `gorm:"primaryKey;column:config_id;autoIncrement"`
	ConfigName  string `gorm:"column:config_name;size:100"`
	ConfigKey   string `gorm:"column:config_key;size:100"`
	ConfigValue string `gorm:"column:config_value;size:500"`
	ConfigType  string `gorm:"column:config_type;size:1;default:'N'"`
	BaseModel
}

// TableName returns the table name for the Config model
func (Config) TableName() string {
	return "sys_configs"
}
