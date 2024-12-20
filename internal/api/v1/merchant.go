package v1

import (
	"net/http"
	"strconv"

	"blacklist/internal/model"
	"blacklist/internal/service"
	"blacklist/pkg/utils"

	"github.com/gin-gonic/gin"
)

type MerchantHandler struct {
	service *service.MerchantService
}

func NewMerchantHandler(s *service.MerchantService) *MerchantHandler {
	return &MerchantHandler{service: s}
}

// CreateMerchant 创建商户
func (h *MerchantHandler) CreateMerchant(c *gin.Context) {
	var merchant model.Merchant
	if err := c.ShouldBindJSON(&merchant); err != nil {
		utils.Error(c, http.StatusBadRequest, "无效的请求参数")
		return
	}

	// 验证必填字段
	if merchant.Name == "" {
		utils.Error(c, http.StatusBadRequest, "商户名称不能为空")
		return
	}

	// 验证手机号格式
	if merchant.Phone != "" && !utils.IsPhoneNumber(merchant.Phone) {
		utils.Error(c, http.StatusBadRequest, "手机号格式错误")
		return
	}

	if err := h.service.Create(&merchant); err != nil {
		utils.Error(c, http.StatusInternalServerError, "创建商户失败")
		return
	}

	utils.Success(c, merchant)
}

// GetMerchant 获取商户详情
func (h *MerchantHandler) GetMerchant(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "无效的商户ID")
		return
	}

	merchant, err := h.service.GetByID(uint(id))
	if err != nil {
		utils.Error(c, http.StatusNotFound, "商户不存在")
		return
	}

	utils.Success(c, merchant)
}

// ListMerchants 获取商户列表
func (h *MerchantHandler) ListMerchants(c *gin.Context) {
	query := &service.MerchantQuery{
		Name:   c.Query("name"),
		Status: 0,
		Page:   1,
		Size:   10,
	}

	if status, err := strconv.Atoi(c.Query("status")); err == nil {
		query.Status = status
	}
	if page, err := strconv.Atoi(c.DefaultQuery("page", "1")); err == nil {
		query.Page = page
	}
	if size, err := strconv.Atoi(c.DefaultQuery("page_size", "10")); err == nil {
		query.Size = size
	}

	merchants, total, err := h.service.List(query)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, "获取商户列表失败")
		return
	}

	utils.SuccessWithPagination(c, merchants, total, query.Page, query.Size)
}

// UpdateMerchant 更新商户信息
func (h *MerchantHandler) UpdateMerchant(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "无效的商户ID")
		return
	}

	var merchant model.Merchant
	if err := c.ShouldBindJSON(&merchant); err != nil {
		utils.Error(c, http.StatusBadRequest, "无效的请求参数")
		return
	}

	merchant.ID = uint(id)
	if err := h.service.Update(&merchant); err != nil {
		utils.Error(c, http.StatusInternalServerError, "更新商户失败")
		return
	}

	utils.Success(c, merchant)
}

// DeleteMerchant 删除商户
func (h *MerchantHandler) DeleteMerchant(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "无效的商户ID")
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		utils.Error(c, http.StatusInternalServerError, "删除商户失败")
		return
	}

	utils.Success(c, nil)
}

// UpdateMerchantStatus 更新商户状态
func (h *MerchantHandler) UpdateMerchantStatus(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "无效的商户ID")
		return
	}

	var req struct {
		Status int `json:"status"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, http.StatusBadRequest, "无效的请求参数")
		return
	}

	if err := h.service.UpdateStatus(uint(id), req.Status); err != nil {
		utils.Error(c, http.StatusInternalServerError, "更新商户状态失败")
		return
	}

	utils.Success(c, nil)
}

// RegenerateAPICredentials 重新生成API凭证
func (h *MerchantHandler) RegenerateAPICredentials(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "无效的商户ID")
		return
	}

	merchant, err := h.service.RegenerateAPICredentials(uint(id))
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, "重新生成API凭证失败")
		return
	}

	utils.Success(c, merchant)
}

// Login 商户登录
func (h *MerchantHandler) Login(c *gin.Context) {
	var req struct {
		APIKey    string `json:"api_key" binding:"required"`
		APISecret string `json:"api_secret" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, http.StatusBadRequest, "无效的请求参数")
		return
	}

	merchant, token, err := h.service.Login(
		c.Request.Context(),
		req.APIKey,
		req.APISecret,
		c.ClientIP(),
		c.Request.UserAgent(),
	)

	if err != nil {
		utils.Error(c, http.StatusUnauthorized, err.Error())
		return
	}

	utils.Success(c, gin.H{
		"merchant": merchant,
		"token":   token,
	})
}

// GetLoginLogs 获取商户登录日志
func (h *MerchantHandler) GetLoginLogs(c *gin.Context) {
	merchantID := utils.GetCurrentMerchantID(c)
	if merchantID == 0 {
		utils.Error(c, http.StatusUnauthorized, "未授权访问")
		return
	}

	page := utils.GetPage(c)
	pageSize := utils.GetPageSize(c)

	logs, total, err := h.service.GetLoginLogs(c.Request.Context(), merchantID, page, pageSize)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, "获取登录日志失败")
		return
	}

	utils.Success(c, gin.H{
		"list":  logs,
		"total": total,
		"page":  page,
		"size":  pageSize,
	})
}
