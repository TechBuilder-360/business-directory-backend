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
	EmailVerified  bool      `json:"email_verified" gorm:"default:false"`
	LastLogin      time.Time `json:"last_login" gorm:"null"`
	Tier           int       `json:"tier" gorm:"default:0"`
	IdentityNumber string    `json:"indentity_number" gorm:"null"`
	IdentityImage  string    `json:"indentity_image" gorm:"null"`
}
