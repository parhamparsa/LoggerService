package audit_log_repository_pg

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/talon-one/talon-backend-assingment/internal/domain/entity"
	"github.com/talon-one/talon-backend-assingment/internal/domain/repository/audit_log_repository"
)

var _ audit_log_repository.AuditLogRepository = &AuditLogRepositoryPostgres{}

type AuditLogRepositoryPostgres struct {
	db *pgxpool.Pool
}

func NewAuditLogRepositoryPostgres(db *pgxpool.Pool) *AuditLogRepositoryPostgres {
	return &AuditLogRepositoryPostgres{db: db}
}

func (r *AuditLogRepositoryPostgres) Save(ctx context.Context, auditLog *entity.AuditLog) error {
	query := `
		INSERT INTO audit_logs (id, timestamp, status_code, duration_ms, method, path, ip_address, user_agent, request_body, response_body)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	_, err := r.db.Exec(ctx,
		query,
		auditLog.ID,
		auditLog.Timestamp,
		auditLog.StatusCode,
		auditLog.DurationMilliSeconds,
		auditLog.Method,
		auditLog.Path,
		auditLog.IPAddress,
		auditLog.UserAgent,
		auditLog.RequestBody,
		auditLog.Response,
	)

	return err
}
