package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type GoLog struct {
	ID             uint      `gorm:"primarykey" json:"id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	UUID           string    `gorm:"type:uuid;uniqueIndex;not null" json:"uuid"`
	Method         string    `gorm:"type:varchar(10);not null" json:"method"`
	Path           string    `gorm:"type:varchar(500);not null" json:"path"`
	StatusCode     int       `gorm:"not null" json:"status_code"`
	ResponseTimeNs int64     `gorm:"not null;column:response_time_ns" json:"response_time_ns"`
	ResponseTimeMs float64   `gorm:"not null;column:response_time_ms" json:"response_time_ms"`
	IPAddress      string    `gorm:"type:varchar(45)" json:"ip_address,omitempty"`
	UserAgent      string    `gorm:"type:text" json:"user_agent,omitempty"`
	RequestBody    string    `gorm:"type:text" json:"request_body,omitempty"`
	ResponseBody   string    `gorm:"type:text" json:"response_body,omitempty"`
	ErrorMessage   string    `gorm:"type:text" json:"error_message,omitempty"`
}

func (GoLog) TableName() string {
	return "go_logs"
}

func (l *GoLog) BeforeCreate(tx *gorm.DB) error {
	if l.UUID == "" {
		l.UUID = uuid.New().String()
	}
	return nil
}
