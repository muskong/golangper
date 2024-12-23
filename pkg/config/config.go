package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	Redis    RedisConfig    `mapstructure:"redis"`
	JWT      JWTConfig      `mapstructure:"jwt"`
	Log      LogConfig      `mapstructure:"log"`
}

type ServerConfig struct {
	Port string `mapstructure:"port" validate:"required"`
	Mode string `mapstructure:"mode" validate:"required,oneof=debug release"`
}

type DatabaseConfig struct {
	Driver   string `mapstructure:"driver"`
	Host     string `mapstructure:"host" validate:"required"`
	Port     int    `mapstructure:"port" validate:"required,min=1,max=65535"`
	User     string `mapstructure:"username" validate:"required"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname" validate:"required"`
	SSLMode  string `mapstructure:"sslmode"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host" validate:"required"`
	Port     int    `mapstructure:"port" validate:"required,min=1,max=65535"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db" validate:"min=0"`
}

type JWTConfig struct {
	Secret string `mapstructure:"secret"`
	Expire string `mapstructure:"expire"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"maxsize"`
	MaxAge     int    `mapstructure:"maxage"`
	MaxBackups int    `mapstructure:"maxbackups"`
}

var Cfg *Config

func Init() error {
	// 设置默认值
	setDefaults()

	// 读取配置文件
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.AddConfigPath(".")

	// 配置 Viper 读取环境变量
	viper.SetEnvPrefix("APP") // 环境变量前缀，可选
	viper.AutomaticEnv()      // 启用环境变量读取

	// 设置环境变量与配置键的映射关系
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	// 绑定特定的环境变量
	viper.BindEnv("server.port", "SERVER_PORT")
	viper.BindEnv("server.mode", "SERVER_MODE")
	viper.BindEnv("database.host", "DB_HOST")
	viper.BindEnv("database.port", "DB_PORT")
	viper.BindEnv("database.username", "DB_USER")
	viper.BindEnv("database.password", "DB_PASSWORD")
	viper.BindEnv("database.dbname", "DB_NAME")
	viper.BindEnv("database.sslmode", "DB_SSLMODE")
	viper.BindEnv("redis.host", "REDIS_HOST")
	viper.BindEnv("redis.port", "REDIS_PORT")
	viper.BindEnv("redis.password", "REDIS_PASSWORD")
	viper.BindEnv("redis.db", "REDIS_DB")
	viper.BindEnv("jwt.secret", "JWT_SECRET")
	viper.BindEnv("jwt.expire", "JWT_EXPIRE")
	viper.BindEnv("log.level", "LOG_LEVEL")
	viper.BindEnv("log.filename", "LOG_FILENAME")
	viper.BindEnv("log.maxsize", "LOG_MAXSIZE")
	viper.BindEnv("log.maxage", "LOG_MAXAGE")
	viper.BindEnv("log.maxbackups", "LOG_MAXBACKUPS")

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("读取配置文件失败: %w", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return fmt.Errorf("解析配置文件失败: %w", err)
	}

	// 验证配置
	if err := validateConfig(&config); err != nil {
		return fmt.Errorf("配置验证失败: %w", err)
	}

	Cfg = &config
	return nil
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
