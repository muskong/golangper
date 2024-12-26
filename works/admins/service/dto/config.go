package dto

type ConfigCreateDTO struct {
	ConfigName        string `json:"configName" binding:"required"`
	ConfigKey         string `json:"configKey" binding:"required"`
	ConfigValue       string `json:"configValue" binding:"required"`
	ConfigType        string `json:"configType" binding:"required,oneof=text json yaml ini"`
	ConfigDescription string `json:"configDescription"`
}

type ConfigUpdateDTO struct {
	ConfigID          int    `json:"configId" binding:"required"`
	ConfigName        string `json:"configName" binding:"required"`
	ConfigValue       string `json:"configValue" binding:"required"`
	ConfigType        string `json:"configType" binding:"required,oneof=text json yaml ini"`
	ConfigStatus      int8   `json:"configStatus" binding:"oneof=0 1"`
	ConfigDescription string `json:"configDescription"`
}

type ConfigInfo struct {
	ConfigID          int    `json:"configId"`
	ConfigName        string `json:"configName"`
	ConfigKey         string `json:"configKey"`
	ConfigValue       string `json:"configValue"`
	ConfigType        string `json:"configType"`
	ConfigStatus      int8   `json:"configStatus"`
	ConfigDescription string `json:"configDescription"`
	CreatedAt         string `json:"createdAt"`
}
