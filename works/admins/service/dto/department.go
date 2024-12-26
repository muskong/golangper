package dto

type CreateDepartmentDTO struct {
	ParentID              int    `json:"parentId"`
	DepartmentName        string `json:"departmentName" binding:"required"`
	DepartmentCode        string `json:"departmentCode" binding:"required"`
	DepartmentDescription string `json:"departmentDescription"`
	DepartmentLeader      string `json:"departmentLeader"`
	DepartmentPhone       string `json:"departmentPhone"`
	DepartmentEmail       string `json:"departmentEmail" binding:"omitempty,email"`
	DepartmentSort        int    `json:"departmentSort"`
}

type UpdateDepartmentDTO struct {
	DepartmentID          int    `json:"departmentId" binding:"required"`
	ParentID              int    `json:"parentId"`
	DepartmentName        string `json:"departmentName" binding:"required"`
	DepartmentCode        string `json:"departmentCode" binding:"required"`
	DepartmentDescription string `json:"departmentDescription"`
	DepartmentLeader      string `json:"departmentLeader"`
	DepartmentPhone       string `json:"departmentPhone"`
	DepartmentEmail       string `json:"departmentEmail" binding:"omitempty,email"`
	DepartmentSort        int    `json:"departmentSort"`
	DepartmentStatus      int8   `json:"departmentStatus" binding:"oneof=0 1"`
}

type DepartmentTreeDTO struct {
	DepartmentID          int                  `json:"departmentId"`
	ParentID              int                  `json:"parentId"`
	DepartmentName        string               `json:"departmentName"`
	DepartmentCode        string               `json:"departmentCode"`
	DepartmentDescription string               `json:"departmentDescription"`
	DepartmentLeader      string               `json:"departmentLeader"`
	DepartmentPhone       string               `json:"departmentPhone"`
	DepartmentEmail       string               `json:"departmentEmail"`
	DepartmentSort        int                  `json:"departmentSort"`
	DepartmentStatus      int8                 `json:"departmentStatus"`
	Children              []*DepartmentTreeDTO `json:"children"`
}
