package model

import (
	"time"
)

// UserProfile ...
type UserProfile struct {
	Base

	FirstName     string     `json:"first_name"`
	LastName      string     `json:"last_name"`
	DisplayName   string     `json:"display_name"`
	EmailAddress  string     `json:"email_address"`
	PhoneNumber   string     `json:"phone_number"`
	EmailVerified bool       `json:"email_verified"`
	LastLogin     *time.Time `json:"last_login"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}

// UserToken ...
type UserToken struct {
	Base

	UserId   string
	Token    string
	ExpiryAt uint
}
