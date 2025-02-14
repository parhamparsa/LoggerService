package audit_log_service

import (
	"context"
	"encoding/json"
	"github.com/talon-one/talon-backend-assingment/internal/domain/entity"
	"github.com/talon-one/talon-backend-assingment/internal/domain/queue"
	"github.com/talon-one/talon-backend-assingment/internal/domain/repository/audit_log_repository"
	"go.uber.org/zap"
)

type AuditLogService struct {
	auditLogRepository audit_log_repository.AuditLogRepository
}

var _ queue.MessageHandler = &AuditLogService{}

func NewAuditLogService(auditLogRepository audit_log_repository.AuditLogRepository) AuditLogService {
	return AuditLogService{auditLogRepository: auditLogRepository}
}

func (s AuditLogService) Save(ctx context.Context, log *entity.AuditLog) error {
	return s.auditLogRepository.Save(ctx, log)
}

func (s AuditLogService) HandleWithRetry(body []byte, retries int) error {
	var err error
	auditLog := entity.AuditLog{}
	if err := json.Unmarshal(body, &auditLog); err != nil {
		zap.L().Error("Cannot unmarshal audit log", zap.ByteString("body", body), zap.Error(err))
		return err
	}
	for i := 0; i < retries; i++ {
		if err = s.Save(context.Background(), &auditLog); err == nil {
			break
		}
	}
	return err
}
