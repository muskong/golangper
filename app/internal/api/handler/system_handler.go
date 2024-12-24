package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"blackapp/internal/api/response"
	"blackapp/internal/service"
	"blackapp/internal/service/dto"
)

type SystemHandler struct {
	systemService service.SystemService
}

func NewSystemHandler(systemService service.SystemService) *SystemHandler {
	return &SystemHandler{systemService: systemService}
}

// @Summary 获取系统指标
// @Description 获取当前系统的CPU、内存、Redis和PostgreSQL信息
// @Tags 系统监控
// @Produce json
// @Success 200 {object} dto.SystemMetrics
// @Router /system/metrics [get]
func (h *SystemHandler) GetSystemMetrics(c *gin.Context) {
	metrics, err := h.systemService.GetSystemMetrics(c.Request.Context())
	if err != nil {
		response.ServerError(c)
		return
	}
	response.Success(c, metrics)
}

// @Summary 管理员登录
// @Description 管理员使用用户名和密码登录
// @Tags 管理员管理
// @Accept json
// @Produce json
// @Param request body dto.AdminLoginDTO true "登录请求"
// @Success 200 {object} map[string]string{"token": "string"}
// @Router /admins/login [post]
func (h *SystemHandler) AdminLogin(c *gin.Context) {
	var req dto.AdminLoginDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "无效的请求参数")
		return
	}

	token, err := h.systemService.AdminLogin(c.Request.Context(), &req)
	if err != nil {
		response.Error(c, 401, "认证失败")
		return
	}

	response.Success(c, gin.H{"token": token})
}

// CreateAdmin 创建管理员
func (h *SystemHandler) CreateAdmin(c *gin.Context) {
	var req dto.CreateAdminDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "无效的请求参数")
		return
	}

	if err := h.systemService.CreateAdmin(c.Request.Context(), &req); err != nil {
		response.ServerError(c)
		return
	}

	response.Success(c, nil)
}

// UpdateAdmin 更新管理员
func (h *SystemHandler) UpdateAdmin(c *gin.Context) {
	var req dto.UpdateAdminDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "无效的请求参数")
		return
	}

	if err := h.systemService.UpdateAdmin(c.Request.Context(), &req); err != nil {
		response.ServerError(c)
		return
	}

	response.Success(c, nil)
}

// ListAdmins 管理员列表
func (h *SystemHandler) ListAdmins(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))

	admins, total, err := h.systemService.ListAdmins(c.Request.Context(), page, size)
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
func (h *SystemHandler) UpdateAdminStatus(c *gin.Context) {
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

	if err := h.systemService.UpdateAdminStatus(c.Request.Context(), int(id), status); err != nil {
		response.ServerError(c)
		return
	}

	response.Success(c, nil)
}

// ListLoginLogs 登录日志列表
func (h *SystemHandler) ListLoginLogs(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))
	userType, _ := strconv.Atoi(c.DefaultQuery("type", "0"))

	logs, total, err := h.systemService.ListLoginLogs(c.Request.Context(), userType, page, size)
	if err != nil {
		response.ServerError(c)
		return
	}

	response.Success(c, gin.H{
		"list":  logs,
		"total": total,
	})
}

// ListQueryLogs 查询日志列表
func (h *SystemHandler) ListQueryLogs(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))
	merchantID, _ := strconv.ParseInt(c.DefaultQuery("merchant_id", "0"), 10, 64)

	logs, total, err := h.systemService.ListQueryLogs(c.Request.Context(), int(merchantID), page, size)
	if err != nil {
		response.ServerError(c)
		return
	}

	response.Success(c, gin.H{
		"list":  logs,
		"total": total,
	})
}
