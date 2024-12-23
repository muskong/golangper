package v1

import (
	"blacklist/internal/model"
	"blacklist/internal/service"
	"blacklist/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	service *service.AdminService
}

func NewAdminHandler(s *service.AdminService) *AdminHandler {
	return &AdminHandler{service: s}
}

func (h *AdminHandler) CreateAdmin(c *gin.Context) {
	var admin model.Admin
	if err := c.ShouldBindJSON(&admin); err != nil {
		utils.Error(c, http.StatusBadRequest, "无效的请求参数")
		return
	}

	if err := h.service.Create(&admin); err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(c, admin)
}

func (h *AdminHandler) UpdateAdmin(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "无效的管理员ID")
		return
	}

	var admin model.Admin
	if err := c.ShouldBindJSON(&admin); err != nil {
		utils.Error(c, http.StatusBadRequest, "无效的请求参数")
		return
	}

	admin.ID = uint(id)
	if err := h.service.Update(&admin); err != nil {
		utils.Error(c, http.StatusInternalServerError, "更新管理员失败")
		return
	}

	utils.Success(c, admin)
}

func (h *AdminHandler) DeleteAdmin(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "无效的管理员ID")
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		utils.Error(c, http.StatusInternalServerError, "删除管理员失败")
		return
	}

	utils.Success(c, nil)
}

func (h *AdminHandler) GetAdmin(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "无效的管理员ID")
		return
	}

	admin, err := h.service.GetByID(uint(id))
	if err != nil {
		utils.Error(c, http.StatusNotFound, "管理员不存在")
		return
	}

	utils.Success(c, admin)
}

func (h *AdminHandler) ListAdmins(c *gin.Context) {
	page := utils.GetPage(c)
	pageSize := utils.GetPageSize(c)

	admins, total, err := h.service.List(page, pageSize)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, "获取管理员列表失败")
		return
	}

	utils.SuccessWithPagination(c, admins, total, page, pageSize)
}
