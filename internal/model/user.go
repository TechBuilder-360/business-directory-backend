package model

import (
	"time"
)

// User ...
type User struct {
	Base

	FirstName      string    `json:"first_name" gorm:"not null"`
	LastName       string    `json:"last_name" gorm:"not null"`
	DisplayName    string    `json:"display_name" gorm:"not null"`
	EmailAddress   string    `json:"email_address" gorm:"not null"`
	PhoneNumber    string    `json:"phone_number" gorm:"null"`
	Avatar         *string   `json:"avatar"`
	EmailVerified  bool      `json:"email_verified" gorm:"default:false"`
	LastLogin      time.Time `json:"last_login" gorm:"null"`
	Tier           uint8     `json:"tier" gorm:"default:0"`
	IdentityNumber string    `json:"identity_number" gorm:"null"`
	IdentityName   string    `json:"identity_name" gorm:"null"`
	IdentityImage  string    `json:"identity_image" gorm:"null"`
}
