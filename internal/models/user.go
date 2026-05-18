package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	UserID uuid.UUID `gorm:"type:uuid;primaryKey" json:"user_id"`

	Name     string `gorm:"not null" json:"name"`
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `gorm:"not null" json:"-"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	user.UserID = uuid.New()
	return nil
}
