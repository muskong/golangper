package impl

import (
	"admins/domain/entity"
	"admins/domain/repository"
	"admins/service/dto"
	"time"

	"github.com/gin-gonic/gin"
)

type logService struct {
	logRepo repository.LogRepository
}

func NewLogService(logRepo repository.LogRepository) *logService {
	return &logService{logRepo: logRepo}
}

func (s *logService) CreateOperationLog(ctx *gin.Context, req *dto.OperationLogCreateDTO) error {
	log := &entity.OperationLog{
		AdminID:           req.AdminID,
		AdminName:         req.AdminName,
		OperationIP:       req.OperationIP,
		OperationLocation: req.OperationLocation,
		OperationBrowser:  req.OperationBrowser,
		OperationOS:       req.OperationOS,
		OperationMethod:   req.OperationMethod,
		OperationPath:     req.OperationPath,
		OperationModule:   req.OperationModule,
		OperationContent:  req.OperationContent,
		OperationStatus:   req.OperationStatus,
		OperationLatency:  req.OperationLatency,
		OperationRequest:  req.OperationRequest,
		OperationResponse: req.OperationResponse,
		CreatedAt:         time.Now(),
	}

	return s.logRepo.CreateOperationLog(ctx, log)
}

func (s *logService) ListOperationLogs(ctx *gin.Context, page, pageSize int) ([]*dto.OperationLogInfo, int64, error) {
	logs, total, err := s.logRepo.ListOperationLogs(ctx, page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	items := make([]*dto.OperationLogInfo, len(logs))
	for i, log := range logs {
		items[i] = &dto.OperationLogInfo{
			LogID:             log.LogID,
			AdminID:           log.AdminID,
			AdminName:         log.AdminName,
			OperationIP:       log.OperationIP,
			OperationLocation: log.OperationLocation,
			OperationBrowser:  log.OperationBrowser,
			OperationOS:       log.OperationOS,
			OperationMethod:   log.OperationMethod,
			OperationPath:     log.OperationPath,
			OperationModule:   log.OperationModule,
			OperationContent:  log.OperationContent,
			OperationStatus:   log.OperationStatus,
			OperationLatency:  log.OperationLatency,
			OperationRequest:  log.OperationRequest,
			OperationResponse: log.OperationResponse,
			CreatedAt:         log.CreatedAt.Format("2006-01-02 15:04:05"),
		}
	}

	return items, total, nil
}

// ... 继续实现其他方法
