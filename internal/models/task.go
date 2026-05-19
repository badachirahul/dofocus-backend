package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Task struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	Title       string `gorm:"not null" json:"title"`

	Completed bool `gorm:"default:false" json:"completed"`

	UserID uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (task *Task) BeforeCreate(tx *gorm.DB) error {
	task.ID = uuid.New()
	return nil
}