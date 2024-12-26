package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"admins/service"
	"admins/service/dto"
	"pkgs/response"
)

type AdminHandler struct {
	adminService service.AdminService
}

func NewAdminHandler(adminService service.AdminService) *AdminHandler {
	return &AdminHandler{adminService: adminService}
}

// @Summary 管理员登录
// @Description 管理员使用用户名和密码登录
// @Tags 管理员管理
// @Accept json
// @Produce json
// @Param request body dto.AdminLoginDTO true "登录请求"
// @Success 200 {object} map[string]string{"token": "string"}
// @Router /admins/login [post]
func (h *AdminHandler) AdminLogin(c *gin.Context) {
	var req dto.AdminLoginDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "无效的请求参数")
		return
	}

	token, err := h.adminService.AdminLogin(c, &req)
	if err != nil {
		response.Error(c, 401, "认证失败")
		return
	}

	response.Success(c, gin.H{"token": token})
}

// CreateAdmin 创建管理员
func (h *AdminHandler) CreateAdmin(c *gin.Context) {
	var req dto.CreateAdminDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "无效的请求参数")
		return
	}

	if err := h.adminService.CreateAdmin(c, &req); err != nil {
		response.ServerError(c)
		return
	}

	response.Success(c, nil)
}

// UpdateAdmin 更新管理员
func (h *AdminHandler) UpdateAdmin(c *gin.Context) {
	var req dto.UpdateAdminDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "无效的请求参数")
		return
	}

	if err := h.adminService.UpdateAdmin(c, &req); err != nil {
		response.ServerError(c)
		return
	}

	response.Success(c, nil)
}

// ListAdmins 管理员列表
func (h *AdminHandler) ListAdmins(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))

	admins, total, err := h.adminService.ListAdmins(c, page, size)
	if err != nil {
		response.ServerError(c)
		return
	}

	response.Success(c, gin.H{
		"list":  admins,
		"total": total,
	})
}

// UpdateAdminStatus 更新管理员状态
func (h *AdminHandler) UpdateAdminStatus(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效的ID")
		return
	}

	status, err := strconv.Atoi(c.Query("status"))
	if err != nil {
		response.BadRequest(c, "无效的状态值")
		return
	}

	err = h.adminService.UpdateAdmin(c, &dto.UpdateAdminDTO{
		AdminID:     int(id),
		AdminStatus: status,
	})

	if err != nil {
		response.ServerError(c)
		return
	}

	response.Success(c, nil)
}

// CreateRole 创建角色
func (h *AdminHandler) CreateRole(c *gin.Context) {
	var req dto.CreateRoleDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "无效的请求参数")
		return
	}

	if err := h.adminService.CreateRole(c, &req); err != nil {
		response.ServerError(c)
		return
	}

	response.Success(c, nil)
}

// UpdateRole 更新角色
func (h *AdminHandler) UpdateRole(c *gin.Context) {
	var req dto.UpdateRoleDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "无效的请求参数")
		return
	}

	if err := h.adminService.UpdateRole(c, &req); err != nil {
		response.ServerError(c)
		return
	}

	response.Success(c, nil)
}

// ListRoles 角色列表
func (h *AdminHandler) ListRoles(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))

	roles, total, err := h.adminService.ListRoles(c, page, size)
	if err != nil {
		response.ServerError(c)
		return
	}

	response.Success(c, gin.H{
		"list":  roles,
		"total": total,
	})
}

// DeleteRole 删除角色
func (h *AdminHandler) DeleteRole(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效的ID")
		return
	}

	if err := h.adminService.DeleteRole(c, int(id)); err != nil {
		response.ServerError(c)
		return
	}

	response.Success(c, nil)
}

// CreateDepartment 创建部门
func (h *AdminHandler) CreateDepartment(c *gin.Context) {
	var req dto.CreateDepartmentDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "无效的请求参数")
		return
	}

	if err := h.adminService.CreateDepartment(c, &req); err != nil {
		response.ServerError(c)
		return
	}

	response.Success(c, nil)
}

// UpdateDepartment 更新部门
func (h *AdminHandler) UpdateDepartment(c *gin.Context) {
	var req dto.UpdateDepartmentDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "无效的请求参数")
		return
	}

	if err := h.adminService.UpdateDepartment(c, &req); err != nil {
		response.ServerError(c)
		return
	}

	response.Success(c, nil)
}

// GetDepartmentTree 获取部门树
func (h *AdminHandler) GetDepartmentTree(c *gin.Context) {
	tree, err := h.adminService.GetDepartmentTree(c)
	if err != nil {
		response.ServerError(c)
		return
	}

	response.Success(c, tree)
}

// DeleteDepartment 删除部门
func (h *AdminHandler) DeleteDepartment(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效的ID")
		return
	}

	if err := h.adminService.DeleteDepartment(c, int(id)); err != nil {
		response.ServerError(c)
		return
	}

	response.Success(c, nil)
}

// GetAdminInfo 获取管理员信息
func (h *AdminHandler) GetAdminInfo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.BadRequest(c, "无效的ID")
		return
	}

	info, err := h.adminService.GetAdminInfo(c, id)
	if err != nil {
		response.ServerError(c)
		return
	}

	response.Success(c, info)
}
