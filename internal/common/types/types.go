package types

import "time"

type (
	LoginResponse struct {
		Authentication Authentication       `json:"authentication"`
		Profile        UserProfile          `json:"profile"`
		Organisations  []OrganisationMember `json:"organisations"`
	}

	OrganisationMember struct {
		ID string `json:"id"`
	}

	UserProfile struct {
		ID            string    `json:"id"`
		FirstName     string    `json:"first_name"`
		LastName      string    `json:"last_name"`
		DisplayName   string    `json:"display_name"`
		EmailAddress  string    `json:"email_address"`
		PhoneNumber   string    `json:"phone_number"`
		EmailVerified bool      `json:"email_verified"`
		LastLogin     time.Time `json:"last_login"`
	}

	Authentication struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}

	RefreshTokenRequest struct {
		RefreshToken string `json:"refresh_token" validate:"required"`
	}

	Query struct {
		Page     uint
		PageSize uint
		Search   string
	}
)

type ENVIRONMENT string
type OrganisationSize string
type RoleType string
type Directory string
type Hash string
