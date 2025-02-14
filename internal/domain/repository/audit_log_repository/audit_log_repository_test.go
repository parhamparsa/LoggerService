package audit_log_repository_test

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"github.com/talon-one/talon-backend-assingment/internal/application/service/audit_log_service"
	"github.com/talon-one/talon-backend-assingment/internal/domain/entity"
	mock_audit_log_repository "github.com/talon-one/talon-backend-assingment/mocks/audit_log_repository"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestSave(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := mock_audit_log_repository.NewMockAuditLogRepository(ctl)

	repo.EXPECT().Save(gomock.Any(), gomock.Any()).Return(nil).Times(1)

	service := audit_log_service.NewAuditLogService(repo)
	bytes, err := json.Marshal(&entity.AuditLog{})
	require.NoError(t, err)
	require.Nil(t, service.HandleWithRetry(bytes, 3))
}
