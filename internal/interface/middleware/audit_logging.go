package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/talon-one/talon-backend-assingment/internal/domain/entity"
	"github.com/talon-one/talon-backend-assingment/internal/domain/queue"
	"go.uber.org/zap"
	"io"
	"net/http"
	"time"
)

type AuditLogging struct {
	Producer queue.Interface
}

func NewAuditLogging(producer queue.Interface) *AuditLogging {
	return &AuditLogging{
		Producer: producer,
	}
}

func (al AuditLogging) Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log Request
		body, _ := io.ReadAll(r.Body)
		r.Body = io.NopCloser(bytes.NewBuffer(body)) // Restore body for next handler

		// Capture response using a custom ResponseWriter.
		responseRecorder := &responseCapture{ResponseWriter: w, body: bytes.NewBuffer(nil)}

		// Capture request duration time.
		now := time.Now()
		next.ServeHTTP(responseRecorder, r)
		duration := time.Since(now)

		// Create entity of request and response for database.
		auditLogEntry := entity.AuditLog{
			ID:                   uuid.New(),
			Timestamp:            time.Now(),
			DurationMilliSeconds: duration.Milliseconds(),
			StatusCode:           responseRecorder.status,
			Method:               r.Method,
			Path:                 r.URL.Path,
			IPAddress:            r.RemoteAddr,
			UserAgent:            r.UserAgent(),
			RequestBody:          string(body),
			Response:             responseRecorder.body.String(),
		}
		auditLogEntryBytes, _ := json.Marshal(auditLogEntry)
		if err := al.Producer.Produce(context.Background(), auditLogEntryBytes); err != nil {
			zap.L().Error("Failed to produce audit log entry", zap.String("error", err.Error()))
		}

		zap.L().Info("The response from the audit log entry ",
			zap.Int("Status", responseRecorder.status),
			zap.String("Body", responseRecorder.body.String()),
		)
	})
}

type responseCapture struct {
	http.ResponseWriter
	status int
	body   *bytes.Buffer
}

func (rc *responseCapture) WriteHeader(code int) {
	rc.status = code
	rc.ResponseWriter.WriteHeader(code)
}

func (rc *responseCapture) Write(data []byte) (int, error) {
	rc.body.Write(data)
	return rc.ResponseWriter.Write(data)
}
