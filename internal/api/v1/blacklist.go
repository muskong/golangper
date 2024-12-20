package v1

import (
	"net/http"
	"strconv"

	"blacklist/internal/model"
	"blacklist/internal/service"
	"blacklist/pkg/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BlacklistHandler struct {
	service *service.BlacklistService
}

func NewBlacklistHandler(s *service.BlacklistService) *BlacklistHandler {
	return &BlacklistHandler{service: s}
}

// CreateBlacklistUser 创建黑名单用户
func (h *BlacklistHandler) CreateBlacklistUser(c *gin.Context) {
	var user model.BlacklistUser
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.Error(c, http.StatusBadRequest, "无效的请求参数")
		return
	}

	// 验证手机号格式
	if !utils.IsPhoneNumber(user.Phone) {
		utils.Error(c, http.StatusBadRequest, "手机号格式错误")
		return
	}

	// 验证身份证号格式
	if !utils.IsIDCard(user.IDCard) {
		utils.Error(c, http.StatusBadRequest, "身份证号格式错误")
		return
	}

	// 验证邮箱格式
	if user.Email != "" && !utils.IsEmail(user.Email) {
		utils.Error(c, http.StatusBadRequest, "邮箱格式错误")
		return
	}

	if err := h.service.Create(&user); err != nil {
		utils.Error(c, http.StatusInternalServerError, "创建用户失败")
		return
	}

	utils.Success(c, user)
}

// GetBlacklistUser 获取黑名单用户详情
func (h *BlacklistHandler) GetBlacklistUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "无效的用户ID")
		return
	}

	user, err := h.service.GetByID(uint(id))
	if err != nil {
		utils.Error(c, http.StatusNotFound, "用户不存在")
		return
	}

	utils.Success(c, user)
}

// ListBlacklistUsers 获取黑名单用户列表
func (h *BlacklistHandler) ListBlacklistUsers(c *gin.Context) {
	// 获取查询参数
	query := &service.BlacklistUserQuery{
		Name:    c.Query("name"),
		Phone:   c.Query("phone"),
		IDCard:  c.Query("id_card"),
		Email:   c.Query("email"),
		Address: c.Query("address"),
		Remark:  c.Query("remark"),
		Page:    1,
		Size:    10,
	}

	// 获取分页参数
	if page, err := strconv.Atoi(c.DefaultQuery("page", "1")); err == nil {
		query.Page = page
	}
	if size, err := strconv.Atoi(c.DefaultQuery("page_size", "10")); err == nil {
		query.Size = size
	}

	// 调用服务层查询
	users, total, err := h.service.List(query)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, "获取用户列表失败")
		return
	}

	utils.SuccessWithPagination(c, users, total, query.Page, query.Size)
}

// UpdateBlacklistUser 更新黑名单用户
func (h *BlacklistHandler) UpdateBlacklistUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "无效的用户ID")
		return
	}

	var user model.BlacklistUser
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.Error(c, http.StatusBadRequest, "无效的请求参数")
		return
	}

	user.ID = uint(id)
	if err := h.service.Update(&user); err != nil {
		utils.Error(c, http.StatusInternalServerError, "更新用户失败")
		return
	}

	utils.Success(c, user)
}

// DeleteBlacklistUser 删除黑名单用户
func (h *BlacklistHandler) DeleteBlacklistUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "无效的用户ID")
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		utils.Error(c, http.StatusInternalServerError, "删除用户失败")
		return
	}

	utils.Success(c, nil)
}

// CheckPhoneExists 检查手机号是否在黑名单中
func (h *BlacklistHandler) CheckPhoneExists(c *gin.Context) {
	phone := c.Query("phone")
	if phone == "" {
		utils.Error(c, http.StatusBadRequest, "手机号不能为空")
		return
	}

	// 验证手机号格式
	if !utils.IsPhoneNumber(phone) {
		utils.Error(c, http.StatusBadRequest, "手机号格式错误")
		return
	}

	exists, err := h.service.CheckPhoneExists(phone)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, "检查失败")
		return
	}

	utils.Success(c, gin.H{"exists": exists})
}

// GetByPhone 根据手机号获取黑名单用户信息
func (h *BlacklistHandler) GetByPhone(c *gin.Context) {
	phone := c.Query("phone")
	if phone == "" {
		utils.Error(c, http.StatusBadRequest, "手机号不能为空")
		return
	}

	// 验证手机号格式
	if !utils.IsPhoneNumber(phone) {
		utils.Error(c, http.StatusBadRequest, "手机号格式错误")
		return
	}

	user, err := h.service.GetByPhone(phone)
	if err != nil {
		status := http.StatusInternalServerError
		if err == gorm.ErrRecordNotFound {
			status = http.StatusNotFound
		}
		utils.Error(c, status, "获取用户信息失败")
		return
	}

	utils.Success(c, user)
}

// CheckExists 检查用户是否在黑名单中
func (h *BlacklistHandler) CheckExists(c *gin.Context) {
	query := &service.ExistsQuery{
		Phone:  c.Query("phone"),
		IDCard: c.Query("id_card"),
		Name:   c.Query("name"),
	}

	if query.Phone == "" && query.IDCard == "" && query.Name == "" {
		utils.Error(c, http.StatusBadRequest, "至少需要提供一个查询条件（手机号/身份证号/姓名）")
		return
	}

	// 检查手机号格式
	if query.Phone != "" && !utils.IsPhoneNumber(query.Phone) {
		utils.Error(c, http.StatusBadRequest, "手机号格式错误")
		return
	}

	exists, err := h.service.CheckExists(query)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, "检查失败")
		return
	}

	utils.Success(c, gin.H{"exists": exists})
}

// GetByIDCard 根据身份证号获取黑名单用户信息
func (h *BlacklistHandler) GetByIDCard(c *gin.Context) {
	idCard := c.Query("id_card")
	if idCard == "" {
		utils.Error(c, http.StatusBadRequest, "身份证号不能为空")
		return
	}

	// 验证身份证号格式
	if !utils.IsIDCard(idCard) {
		utils.Error(c, http.StatusBadRequest, "身份证号格式错误")
		return
	}

	user, err := h.service.GetByIDCard(idCard)
	if err != nil {
		status := http.StatusInternalServerError
		if err == gorm.ErrRecordNotFound {
			status = http.StatusNotFound
		}
		utils.Error(c, status, "获取用户信息失败")
		return
	}

	utils.Success(c, user)
}

// GetByName 根据姓名获取黑名单用户信息列表
func (h *BlacklistHandler) GetByName(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		utils.Error(c, http.StatusBadRequest, "姓名不能为空")
		return
	}

	users, err := h.service.GetByName(name)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, "获取用户信息失败")
		return
	}

	utils.Success(c, gin.H{
		"total": len(users),
		"items": users,
	})
}

// GetQueryLogs 获取黑名单查询日志
func (h *BlacklistHandler) GetQueryLogs(c *gin.Context) {
	merchantID := utils.GetCurrentMerchantID(c)
	if merchantID == 0 {
		utils.Error(c, http.StatusUnauthorized, "未授权访问")
		return
	}

	page := utils.GetPage(c)
	pageSize := utils.GetPageSize(c)

	logs, total, err := h.service.GetQueryLogs(c.Request.Context(), merchantID, page, pageSize)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, "获取查询日志失败")
		return
	}

	utils.Success(c, gin.H{
		"list":  logs,
		"total": total,
		"page":  page,
		"size":  pageSize,
	})
}

// GetQueryLogsByPhone 获取指定手机号的查询日志
func (h *BlacklistHandler) GetQueryLogsByPhone(c *gin.Context) {
	phone := c.Query("phone")
	if phone == "" {
		utils.Error(c, http.StatusBadRequest, "手机号不能为空")
		return
	}

	page := utils.GetPage(c)
	pageSize := utils.GetPageSize(c)

	logs, total, err := h.service.GetQueryLogsByPhone(c.Request.Context(), phone, page, pageSize)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, "获取查询日志失败")
		return
	}

	utils.Success(c, gin.H{
		"list":  logs,
		"total": total,
		"page":  page,
		"size":  pageSize,
	})
}
