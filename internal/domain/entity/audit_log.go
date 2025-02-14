package entity

import (
	"github.com/google/uuid"
	"time"
)

type AuditLog struct {
	ID                   uuid.UUID `json:"id"`
	Timestamp            time.Time `json:"timestamp"`
	StatusCode           int       `json:"status_code"`
	DurationMilliSeconds int64     `json:"duration_ms"`
	Method               string    `json:"method"`
	Path                 string    `json:"path"`
	IPAddress            string    `json:"ip_address"`
	UserAgent            string    `json:"user_agent"`
	RequestBody          string    `json:"request_body"`
	Response             string    `json:"response"`
}
