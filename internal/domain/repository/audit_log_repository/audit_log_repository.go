package audit_log_repository

import (
	"context"
	"github.com/talon-one/talon-backend-assingment/internal/domain/entity"
)

//go:generate mockgen -source=audit_log_repository.go -destination=../../../../mocks/audit_log_repository/repo.go
type AuditLogRepository interface {
	Save(ctx context.Context, user *entity.AuditLog) error
}
