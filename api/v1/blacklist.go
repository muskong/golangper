package v1

import (
	"net/http"
	"strconv"

	"blacklist/internal/model"
	"blacklist/internal/service"

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
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求参数"})
		return
	}

	if err := h.service.Create(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建用户失败"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// GetBlacklistUser 获取黑名单用户详情
func (h *BlacklistHandler) GetBlacklistUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的用户ID"})
		return
	}

	user, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
		return
	}

	c.JSON(http.StatusOK, user)
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
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "获取用户列表失败",
			"msg":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"items": users,
		"page":  query.Page,
		"size":  query.Size,
	})
}

// UpdateBlacklistUser 更新黑名单用户
func (h *BlacklistHandler) UpdateBlacklistUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的用户ID"})
		return
	}

	var user model.BlacklistUser
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求参数"})
		return
	}

	user.ID = uint(id)
	if err := h.service.Update(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新用户失败"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// DeleteBlacklistUser 删除黑名单用户
func (h *BlacklistHandler) DeleteBlacklistUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的用户ID"})
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除用户失败"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// CheckPhoneExists 检查手机号是否在黑名单中
func (h *BlacklistHandler) CheckPhoneExists(c *gin.Context) {
	phone := c.Query("phone")
	if phone == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "手机号不能为空",
		})
		return
	}

	exists, err := h.service.CheckPhoneExists(phone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "检查失败",
			"msg":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"exists": exists,
	})
}

// GetByPhone 根据手机号获取黑名单用户信息
func (h *BlacklistHandler) GetByPhone(c *gin.Context) {
	phone := c.Query("phone")
	if phone == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "手机号不能为空",
		})
		return
	}

	user, err := h.service.GetByPhone(phone)
	if err != nil {
		status := http.StatusInternalServerError
		if err == gorm.ErrRecordNotFound {
			status = http.StatusNotFound
		}
		c.JSON(status, gin.H{
			"error": "获取用户信息失败",
			"msg":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}
