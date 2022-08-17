package types

import "time"

//import "github.com/go-playground/validator/v10"

// AuthRequest ...
type AuthRequest struct {
	EmailAddress string `json:"email-address" validate:"required"`
	Token        string `json:"token" validate:"required"`
}

// EmailRequest ...
type EmailRequest struct {
	EmailAddress string `json:"email_address" validate:"required,email"`
}

// Registration ...
type Registration struct {
	EmailAddress string `json:"email_address" validate:"required,email"`
	FirstName    string `json:"first_name" validate:"required"`
	LastName     string `json:"last_name" validate:"required"`
	DisplayName  string `json:"display_name"`
	PhoneNumber  string `json:"phone_number"`
}

// UserProfile ...
type UserProfile struct {
	ID            string     `json:"id"`
	FirstName     string     `json:"first_name"`
	LastName      string     `json:"last_name"`
	DisplayName   string     `json:"display_name"`
	EmailAddress  string     `json:"email_address"`
	PhoneNumber   string     `json:"phone_number"`
	EmailVerified bool       `json:"email_verified"`
	LastLogin     *time.Time `json:"last_login"`
}
