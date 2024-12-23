package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Database struct {
		Host     string `mapstructure:"host" validate:"required"`
		Port     int    `mapstructure:"port" validate:"required,min=1,max=65535"`
		User     string `mapstructure:"username" validate:"required"`
		Password string `mapstructure:"password"`
		DBName   string `mapstructure:"dbname" validate:"required"`
		SSLMode  string `mapstructure:"sslmode"`
	} `mapstructure:"database"`
	Redis struct {
		Host     string `mapstructure:"host" validate:"required"`
		Port     int    `mapstructure:"port" validate:"required,min=1,max=65535"`
		Password string `mapstructure:"password"`
		DB       int    `mapstructure:"db" validate:"min=0"`
	} `mapstructure:"redis"`
	Server struct {
		Port string `mapstructure:"port" validate:"required"`
		Mode string `mapstructure:"mode" validate:"required,oneof=debug release"`
	} `mapstructure:"server"`
}

func LoadConfig() (*Config, error) {
	// 设置默认值
	setDefaults()

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")
	viper.AddConfigPath(".") // 添加多个配置路径
	viper.AutomaticEnv()     // 支持环境变量覆盖

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %w", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("解析配置文件失败: %w", err)
	}

	// 验证配置
	if err := validateConfig(&config); err != nil {
		return nil, fmt.Errorf("配置验证失败: %w", err)
	}

	return &config, nil
}

func setDefaults() {
	viper.SetDefault("database.port", 5432)
	viper.SetDefault("database.sslmode", "disable")
	viper.SetDefault("redis.port", 6379)
	viper.SetDefault("redis.db", 0)
}

func validateConfig(cfg *Config) error {
	// 这里可以添加自定义验证逻辑
	return nil
}