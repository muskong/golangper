package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/muskong/gopermission/works/merchants/service"
	"github.com/muskong/gopermission/works/merchants/service/dto"

	"github.com/muskong/gopermission/works/pkgs/response"
)

type MerchantHandler struct {
	merchantService service.MerchantService
}

func NewMerchantHandler(merchantService service.MerchantService) *MerchantHandler {
	return &MerchantHandler{merchantService: merchantService}
}

func (h *MerchantHandler) Create(c *gin.Context) {
	var req dto.CreateMerchantDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "无效的请求参数")
		return
	}

	if err := h.merchantService.Create(c, &req); err != nil {
		response.ServerError(c)
		return
	}

	response.Success(c, nil)
}

func (h *MerchantHandler) Update(c *gin.Context) {
	var req dto.UpdateMerchantDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "无效的请求参数")
		return
	}

	if err := h.merchantService.Update(c, &req); err != nil {
		response.ServerError(c)
		return
	}

	response.Success(c, nil)
}

func (h *MerchantHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效的ID")
		return
	}

	if err := h.merchantService.Delete(c, int(id)); err != nil {
		response.ServerError(c)
		return
	}

	response.Success(c, nil)
}

func (h *MerchantHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效的ID")
		return
	}

	merchant, err := h.merchantService.GetByID(c, int(id))
	if err != nil {
		response.ServerError(c)
		return
	}

	response.Success(c, merchant)
}

func (h *MerchantHandler) List(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))

	merchants, total, err := h.merchantService.List(c, page, size)
	if err != nil {
		response.ServerError(c)
		return
	}

	response.Success(c, gin.H{
		"list":  merchants,
		"total": total,
	})
}

func (h *MerchantHandler) UpdateStatus(c *gin.Context) {
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

	if err := h.merchantService.UpdateStatus(c, int(id), status); err != nil {
		response.ServerError(c)
		return
	}

	response.Success(c, nil)
}

func (h *MerchantHandler) Login(c *gin.Context) {
	var req dto.MerchantLoginDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "无效的请求参数")
		return
	}

	if req.APIKey == "" || req.APISecret == "" {
		response.BadRequest(c, "API Key和Secret不能为空")
		return
	}

	token, err := h.merchantService.Login(c, &req)
	if err != nil {
		response.Error(c, 401, "认证失败")
		return
	}

	response.Success(c, gin.H{"token": token})
}

// ListLoginLogs 登录日志列表
func (h *MerchantHandler) ListLoginLogs(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))
	userType, _ := strconv.Atoi(c.DefaultQuery("type", "0"))

	logs, total, err := h.merchantService.ListLoginLogs(c, userType, page, size)
	if err != nil {
		response.ServerError(c)
		return
	}

	response.Success(c, gin.H{
		"list":  logs,
		"total": total,
	})
}
