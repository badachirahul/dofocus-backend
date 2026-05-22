package models

import (
	"time"

	"github.com/google/uuid"
)

type FocusSession struct {
	SessionID uuid.UUID `gorm:"type:uuid;primaryKey" json:"session_id"`

	TaskID uuid.UUID `gorm:"type:uuid;not null" json:"task_id"`

	UserID uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`

	// Example:
	// 25 min = 1500 seconds
	TimerDurationSeconds int `gorm:"not null" json:"timer_duration_seconds"`

	// Actual focused duration accumulated
	FocusedSeconds int `gorm:"default:0" json:"focused_seconds"`

	// active | paused | completed | cancelled
	Status string `gorm:"type:varchar(20);not null" json:"status"`

	// Used when session is currently active
	LastResumedAt *time.Time `json:"last_resumed_at"`

	// First time session started
	StartedAt time.Time `gorm:"not null" json:"started_at"`

	// Session completion/cancellation time
	EndedAt *time.Time `json:"ended_at"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}