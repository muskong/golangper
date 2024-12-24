package dto

type SystemMetrics struct {
	CPU struct {
		Usage     float64 `json:"usage"`
		Cores     int     `json:"cores"`
		ModelName string  `json:"model_name"`
	} `json:"cpu"`

	Memory struct {
		Total     uint64  `json:"total"`
		Used      uint64  `json:"used"`
		Free      uint64  `json:"free"`
		UsageRate float64 `json:"usage_rate"`
	} `json:"memory"`

	Redis struct {
		Connected  bool   `json:"connected"`
		UsedMemory uint64 `json:"used_memory"`
		Keys       int    `json:"keys"`
		Clients    int    `json:"clients"`
	} `json:"redis"`

	Postgres struct {
		Connected   bool   `json:"connected"`
		Version     string `json:"version"`
		Connections int    `json:"connections"`
		DBSize      uint64 `json:"db_size"`
	} `json:"postgres"`
}

type AdminLoginDTO struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CreateAdminDTO struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
}

type UpdateAdminDTO struct {
	ID       int    `json:"id" binding:"required"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Status   int    `json:"status"`
}

type AdminDTO struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Name      string `json:"name"`
	Status    int    `json:"status"`
	LastLogin string `json:"last_login"`
}
