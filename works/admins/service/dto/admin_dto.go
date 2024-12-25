package dto

type AdminLoginDTO struct {
	AdminName     string `json:"adminName" binding:"required"`
	AdminPassword string `json:"adminPassword" binding:"required"`
}

type CreateAdminDTO struct {
	AdminName     string `json:"adminName" binding:"required"`
	AdminPassword string `json:"adminPassword" binding:"required"`
	AdminEmail    string `json:"adminEmail" binding:"required"`
	AdminPhone    string `json:"adminPhone" binding:"required"`
	AdminSex      int8   `json:"adminSex" binding:"required"`
	AdminAvatar   string `json:"adminAvatar"`
	AdminStatus   int    `json:"adminStatus" binding:"required"`
}

type UpdateAdminDTO struct {
	AdminID       int    `json:"adminID" binding:"required"`
	AdminName     string `json:"adminName"`
	AdminPassword string `json:"adminPassword"`
	AdminEmail    string `json:"adminEmail"`
	AdminPhone    string `json:"adminPhone"`
	AdminSex      int8   `json:"adminSex"`
	AdminAvatar   string `json:"adminAvatar"`
	AdminStatus   int    `json:"adminStatus"`
}

type AdminDTO struct {
	AdminID      int    `json:"adminID"`
	DepartmentID *int   `json:"departmentID"`
	AdminName    string `json:"adminName"`
	AdminEmail   string `json:"adminEmail"`
	AdminPhone   string `json:"adminPhone"`
	AdminSex     int8   `json:"adminSex"`
	AdminAvatar  string `json:"adminAvatar"`
	AdminStatus  int    `json:"adminStatus"`
	AdminLogin   string `json:"adminLogin"`
}
