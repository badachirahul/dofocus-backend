package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OTPVerification struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Email     string    `gorm:"unique;not null" json:"email"`
	OTP       string    `gorm:"not null" json:"-"`
	Verified  bool      `gorm:"default:false" json:"verified"`
	ExpiresAt time.Time `json:"expires_at"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

func (otp *OTPVerification) BeforeCreate(tx *gorm.DB) error {
	otp.ID = uuid.New()
	return nil
}