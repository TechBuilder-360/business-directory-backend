package types

// AuthRequest ...
type AuthRequest struct {
	EmailAddress string `json:"email_address" validate:"required"`
	Otp          string `json:"otp" validate:"required"`
}

// EmailRequest ...
type EmailRequest struct {
	EmailAddress string `json:"email_address" validate:"required,email"`
}

type UpgradeUserTierRequest struct {
	IdentityNumber string `json:"identity_number" validate:"required" `
	IdentityName   string `json:"identity_name" validate:"required"`
	IdentityImage  string `json:"identity_image" validate:"required" `
}
