package dto

type CreateRoleDTO struct {
	RoleName        string `json:"roleName" binding:"required"`
	RoleCode        string `json:"roleCode" binding:"required"`
	RoleDescription string `json:"roleDescription"`
	MenuIDs         []int  `json:"menuIds" binding:"required"`
	DepartmentIDs   []int  `json:"departmentIds"`
}

type UpdateRoleDTO struct {
	RoleID          int    `json:"roleId" binding:"required"`
	RoleName        string `json:"roleName" binding:"required"`
	RoleCode        string `json:"roleCode" binding:"required"`
	RoleDescription string `json:"roleDescription"`
	RoleStatus      int8   `json:"roleStatus" binding:"oneof=0 1"`
	MenuIDs         []int  `json:"menuIds" binding:"required"`
	DepartmentIDs   []int  `json:"departmentIds"`
}

type RoleDTO struct {
	RoleID          int    `json:"roleId"`
	RoleName        string `json:"roleName"`
	RoleCode        string `json:"roleCode"`
	RoleDescription string `json:"roleDescription"`
	RoleStatus      int8   `json:"roleStatus"`
	CreatedAt       string `json:"createdAt"`

	MenuIDs       []int `json:"menuIds"`
	DepartmentIDs []int `json:"departmentIds"`
}
