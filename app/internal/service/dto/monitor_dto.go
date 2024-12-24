package dto

// CPUInfo CPU信息
type CPUInfo struct {
	Usage     float64 `json:"usage"`      // CPU使用率
	Cores     int     `json:"cores"`      // CPU核心数
	ModelName string  `json:"model_name"` // CPU型号
}

// MemoryInfo 内存信息
type MemoryInfo struct {
	Total     uint64  `json:"total"`      // 总内存
	Used      uint64  `json:"used"`       // 已使用内存
	Free      uint64  `json:"free"`       // 空闲内存
	UsageRate float64 `json:"usage_rate"` // 使用率
}

// RedisInfo Redis信息
type RedisInfo struct {
	Connected  bool   `json:"connected"`   // 连接状态
	UsedMemory uint64 `json:"used_memory"` // 使用内存
	Keys       int    `json:"keys"`        // 键数量
	Clients    int    `json:"clients"`     // 客户端数量
}

// PostgresInfo PostgreSQL信息
type PostgresInfo struct {
	Connected   bool   `json:"connected"`   // 连接状态
	Version     string `json:"version"`     // 版本信息
	Connections int    `json:"connections"` // 当前连接数
	DBSize      uint64 `json:"db_size"`     // 数据库大小(字节)
}

// SystemMetrics 系统指标
type SystemMetrics struct {
	CPU      CPUInfo      `json:"cpu"`
	Memory   MemoryInfo   `json:"memory"`
	Redis    RedisInfo    `json:"redis"`
	Postgres PostgresInfo `json:"postgres"`
}
