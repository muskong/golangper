package impl

import (
	"blacklists/domain/constants"
	"blacklists/domain/entity"
	"blacklists/domain/repository"
	"blacklists/service/dto"
	"pkgs/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BlacklistService struct {
	repo         repository.BlacklistRepository
	queryLogRepo repository.QueryLogRepository
}

func NewBlacklistService(repo repository.BlacklistRepository, queryLogRepo repository.QueryLogRepository) *BlacklistService {
	return &BlacklistService{repo: repo, queryLogRepo: queryLogRepo}
}

func (s *BlacklistService) Create(ctx *gin.Context, req *dto.CreateBlacklistDTO) error {
	blacklist := &entity.Blacklist{
		Name:    req.Name,
		Phone:   req.Phone,
		IDCard:  req.IDCard,
		Email:   req.Email,
		Address: req.Address,
		Remark:  req.Remark,
		Status:  constants.BlacklistStatusPending, // 待审核状态
	}

	if err := s.repo.Create(ctx, blacklist); err != nil {
		logger.Logger.Error("创建黑名单记录失败", zap.Error(err))
		return err
	}

	return nil
}

func (s *BlacklistService) Update(ctx *gin.Context, req *dto.UpdateBlacklistDTO) error {
	blacklist, err := s.repo.FindByID(ctx, req.ID)
	if err != nil {
		return err
	}

	blacklist.Name = req.Name
	blacklist.Phone = req.Phone
	blacklist.IDCard = req.IDCard
	blacklist.Email = req.Email
	blacklist.Address = req.Address
	blacklist.Remark = req.Remark
	blacklist.Status = req.Status

	return s.repo.Update(ctx, blacklist)
}

func (s *BlacklistService) Delete(ctx *gin.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

func (s *BlacklistService) GetByID(ctx *gin.Context, id int) (*dto.BlacklistDTO, error) {
	blacklist, err := s.repo.FindByID(ctx, id)
	if err != nil {
		logger.Logger.Error("查询黑名单记录失败", zap.Int("id", id), zap.Error(err))
		return nil, err
	}

	return toBlacklistDTO(blacklist), nil
}

func (s *BlacklistService) List(ctx *gin.Context, page, size int) ([]*dto.BlacklistDTO, int64, error) {
	blacklists, total, err := s.repo.List(ctx, page, size)
	if err != nil {
		logger.Logger.Error("查询黑名单记录失败", zap.Int("page", page), zap.Int("size", size), zap.Error(err))
		return nil, 0, err
	}

	dtos := make([]*dto.BlacklistDTO, len(blacklists))
	for i, blacklist := range blacklists {
		dtos[i] = toBlacklistDTO(blacklist)
	}

	return dtos, total, nil
}

func (s *BlacklistService) UpdateStatus(ctx *gin.Context, id int, status int) error {
	return s.repo.UpdateStatus(ctx, id, status)
}

func (s *BlacklistService) Check(ctx *gin.Context, req *dto.CheckBlacklistDTO) (bool, error) {
	var err error
	var exists bool

	// 按照优先级依次检查手机号、身份证、姓名
	if req.Phone != "" {
		_, err = s.repo.CheckByPhone(ctx, req.Phone)
	}

	if req.IDCard != "" {
		_, err = s.repo.CheckByIDCard(ctx, req.IDCard)
	}

	if req.Name != "" {
		_, err = s.repo.CheckByName(ctx, req.Name)
	}

	exists = err == nil

	merchantID := ctx.GetInt("merchantID")
	queryLog := &entity.QueryLog{
		MerchantID: merchantID,
		Phone:      req.Phone,
		IDCard:     req.IDCard,
		Name:       req.Name,
		IP:         ctx.ClientIP(),
		UserAgent:  ctx.Request.UserAgent(),
		Exists:     exists,
	}

	if err := s.queryLogRepo.Create(ctx, queryLog); err != nil {
		logger.Logger.Error("创建查询日志失败", zap.Error(err))
		return false, err
	}

	logger.Logger.Info("黑名单检查结果", zap.Bool("exists", exists))
	return exists, nil
}

func (s *BlacklistService) ListQueryLogs(ctx *gin.Context, merchantID int, page, size int) ([]*dto.QueryLogDTO, int64, error) {
	logs, total, err := s.queryLogRepo.List(ctx, merchantID, page, size)
	if err != nil {
		logger.Logger.Error("查询查询日志失败", zap.Int("merchantID", merchantID), zap.Int("page", page), zap.Int("size", size), zap.Error(err))
		return nil, 0, err
	}

	dtos := make([]*dto.QueryLogDTO, len(logs))
	for i, log := range logs {
		dtos[i] = toQueryLogDTO(log)
	}
	return dtos, total, nil
}

func toBlacklistDTO(blacklist *entity.Blacklist) *dto.BlacklistDTO {
	return &dto.BlacklistDTO{
		ID:        blacklist.ID,
		Name:      blacklist.Name,
		Phone:     blacklist.Phone,
		IDCard:    blacklist.IDCard,
		Email:     blacklist.Email,
		Address:   blacklist.Address,
		Remark:    blacklist.Remark,
		Status:    blacklist.Status,
		CreatedAt: blacklist.CreatedAt,
		UpdatedAt: blacklist.UpdatedAt,
	}
}

func toQueryLogDTO(log *entity.QueryLog) *dto.QueryLogDTO {
	return &dto.QueryLogDTO{
		ID:         log.ID,
		MerchantID: log.MerchantID,
		Phone:      log.Phone,
		IDCard:     log.IDCard,
		Name:       log.Name,
		IP:         log.IP,
		UserAgent:  log.UserAgent,
		Exists:     log.Exists,
		CreatedAt:  log.CreatedAt,
	}
}
