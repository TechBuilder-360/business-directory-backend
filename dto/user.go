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