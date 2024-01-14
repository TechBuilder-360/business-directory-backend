package types

import "time"

type MailTemplate struct {
	Header   string
	Code     string
	Token    string
	ToEmail  string
	ToName   string
	Subject  string
	Duration time.Duration
}

// Registration ...
type Registration struct {
	EmailAddress string  `json:"email_address" validate:"required,email"`
	Avatar       *string `json:"avatar"`
	FirstName    string  `json:"first_name" validate:"required"`
	LastName     string  `json:"last_name" validate:"required"`
	DisplayName  *string `json:"display_name"`
	PhoneNumber  *string `json:"phone_number" validate:"e164"`
}

type Authenticate struct {
	EmailAddress string `json:"email_address" validate:"required,email"`
	Otp          string `json:"otp"`
}

type RefreshToken struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}
