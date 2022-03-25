package dto

//import "github.com/go-playground/validator/v10"

// AuthRequest ...
type AuthRequest struct {
	Email string `json:"email" validate:"required,email"`
	Token string `json:"token"`
}

// EmailRequest ...
type EmailRequest struct {
	Email string `json:"email" binding:"required,email"`
}

// Registration ...
type Registration struct {
	EmailAddress string `json:"email_address" validate:"required,email"`
	FirstName    string `json:"first_name" validate:"required"`
	LastName     string `json:"last_name" validate:"required"`
	DisplayName  string `json:"display_name"`
	PhoneNumber  string `json:"phone_number"`
}
