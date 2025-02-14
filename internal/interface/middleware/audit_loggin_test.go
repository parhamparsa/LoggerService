package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/talon-one/talon-backend-assingment/internal/domain/entity"
	mock_queue "github.com/talon-one/talon-backend-assingment/mocks"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestAuditLoggingMiddleware(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mockedProducer := mock_queue.NewMockInterface(ctl)
	middlewareInstance := NewAuditLogging(mockedProducer)

	dummyHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(10 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
	})

	handlerWithMiddleware := middlewareInstance.Logging(dummyHandler)

	requestBody := `{"key":"value"}`
	req := httptest.NewRequest(http.MethodPost, "/test-path", bytes.NewBufferString(requestBody))
	req.Header.Set("User-Agent", "GoTest")
	req.RemoteAddr = "127.0.0.1:1234"

	recorder := httptest.NewRecorder()

	expectedAuditLog := entity.AuditLog{
		ID:                   uuid.Nil,
		Timestamp:            time.Time{},
		DurationMilliSeconds: 0,
		StatusCode:           http.StatusOK,
		Method:               "POST",
		Path:                 "/test-path",
		IPAddress:            "127.0.0.1:1234",
		UserAgent:            "GoTest",
		RequestBody:          requestBody,
		Response:             "OK",
	}

	mockedProducer.EXPECT().Produce(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, msgBytes []byte) error {
		var actualAuditLog entity.AuditLog
		require.NoError(t, json.Unmarshal(msgBytes, &actualAuditLog))

		require.NoError(t, uuid.Validate(actualAuditLog.ID.String()))
		require.NotZero(t, actualAuditLog.Timestamp)
		require.Greater(t, actualAuditLog.DurationMilliSeconds, int64(0))

		actualAuditLog.ID = uuid.Nil
		actualAuditLog.Timestamp = time.Time{}
		actualAuditLog.DurationMilliSeconds = 0
		require.Equal(t, expectedAuditLog, actualAuditLog)
		return nil
	})

	handlerWithMiddleware.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "OK", recorder.Body.String())
}
