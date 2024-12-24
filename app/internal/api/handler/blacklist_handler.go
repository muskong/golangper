package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"blackapp/internal/api/response"
	"blackapp/internal/service"
	"blackapp/internal/service/dto"
)

type BlacklistHandler struct {
	blacklistService service.BlacklistService
}

func NewBlacklistHandler(blacklistService service.BlacklistService) *BlacklistHandler {
	return &BlacklistHandler{blacklistService: blacklistService}
}

func (h *BlacklistHandler) Create(c *gin.Context) {
	var req dto.CreateBlacklistDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "无效的请求参数")
		return
	}

	if err := h.blacklistService.Create(c.Request.Context(), &req); err != nil {
		response.ServerError(c)
		return
	}

	response.Success(c, nil)
}

func (h *BlacklistHandler) Update(c *gin.Context) {
	var req dto.UpdateBlacklistDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "无效的请求参数")
		return
	}

	if err := h.blacklistService.Update(c.Request.Context(), &req); err != nil {
		response.ServerError(c)
		return
	}

	response.Success(c, nil)
}

func (h *BlacklistHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效的ID")
		return
	}

	if err := h.blacklistService.Delete(c.Request.Context(), int(id)); err != nil {
		response.ServerError(c)
		return
	}

	response.Success(c, nil)
}

func (h *BlacklistHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效的ID")
		return
	}

	blacklist, err := h.blacklistService.GetByID(c.Request.Context(), int(id))
	if err != nil {
		response.ServerError(c)
		return
	}

	response.Success(c, blacklist)
}

func (h *BlacklistHandler) List(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))

	blacklists, total, err := h.blacklistService.List(c.Request.Context(), page, size)
	if err != nil {
		response.ServerError(c)
		return
	}

	response.Success(c, gin.H{
		"list":  blacklists,
		"total": total,
	})
}

func (h *BlacklistHandler) UpdateStatus(c *gin.Context) {
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

	if err := h.blacklistService.UpdateStatus(c.Request.Context(), int(id), status); err != nil {
		response.ServerError(c)
		return
	}

	response.Success(c, nil)
}

func (h *BlacklistHandler) Check(c *gin.Context) {
	var req dto.CheckBlacklistDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "无效的请求参数")
		return
	}

	exists, err := h.blacklistService.Check(c.Request.Context(), &req)
	if err != nil {
		response.ServerError(c)
		return
	}

	response.Success(c, gin.H{"exists": exists})
}
