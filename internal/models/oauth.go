package models

import (
	"time"
)

type Token struct {
	ID              string    `gorm:"primaryKey;type:uuid"`
	UserID          string    `gorm:"type:uuid;not null"`
	AccessToken     string    `gorm:"not null"`
	AccessExpiresAt time.Time `gorm:"not null"`
	Revoked         bool      `gorm:"not null;default:false"`
	CreatedAt       time.Time
}
