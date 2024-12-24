package dto

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
	LastLogin string `json:"lastLogin"`
}
