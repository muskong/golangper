package impl

import (
	"admins/api/dto"
	"admins/domain/entity"
	"admins/domain/repository"
	"context"
	"time"
)

type logService struct {
	logRepo repository.LogRepository
}

func NewLogService(logRepo repository.LogRepository) *logService {
	return &logService{logRepo: logRepo}
}

func (s *logService) CreateOperationLog(ctx context.Context, req *dto.OperationLogCreateRequest) error {
	log := &entity.OperationLog{
		UserID:            req.UserID,
		UserName:          req.UserName,
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

func (s *logService) ListOperationLogs(ctx context.Context, query dto.LogQueryRequest) (*dto.PageResponse, error) {
	logs, total, err := s.logRepo.ListOperationLogs(ctx, query)
	if err != nil {
		return nil, err
	}

	items := make([]dto.OperationLogInfo, len(logs))
	for i, log := range logs {
		items[i] = dto.OperationLogInfo{
			LogID:             log.LogID,
			UserID:            log.UserID,
			UserName:          log.UserName,
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

	return &dto.PageResponse{
		Total:    total,
		List:     items,
		PageNum:  query.PageNum,
		PageSize: query.PageSize,
	}, nil
}

// ... 继续实现其他方法