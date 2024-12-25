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

	if err := h.adminService.UpdateAdminStatus(c, int(id), status); err != nil {
		response.ServerError(c)
		return
	}

	response.Success(c, nil)
}
