package handler

import (
	"blacklist/internal/service"
	"blacklist/internal/service/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	adminService service.AdminService
}

func (h *AdminHandler) Create(c *gin.Context) {
	// 处理用户创建请求
	adminDTO := &dto.AdminDTO{}
	if err := c.ShouldBindJSON(adminDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.adminService.Create(adminDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Admin created successfully"})
}

func (h *AdminHandler) Get(c *gin.Context) {
	id := c.Param("id")
	admin, err := h.adminService.FindByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, admin)
}
