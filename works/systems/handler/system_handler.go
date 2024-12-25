package handler

import (
	"github.com/gin-gonic/gin"

	"pkgs/response"
	"systems/service"
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
	metrics, err := h.systemService.GetSystemMetrics(c)
	if err != nil {
		response.ServerError(c)
		return
	}
	response.Success(c, metrics)
}
