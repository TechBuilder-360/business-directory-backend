package types

type Contact struct {
	Email        string         `json:"email" validate:"required"`
	PhoneNumbers []PhoneDetails `json:"phone_numbers"`
}

type PhoneDetails struct {
	Type        string `json:"types" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	CountryCode string `json:"country_code" validate:"required"`
}

type Location struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type Address struct {
	CountryCode string `json:"country_code"`
	ZipCode     string `json:"zip_code"`
	Street      string `json:"street"`
	City        string `json:"city"`
	State       string `json:"state"`
}

// CreateOrganisationReq ...
type CreateOrganisationReq struct {
	Name             string           `json:"organisation_name" validate:"required"`
	Category         string           `json:"category" validate:"required"`
	Country          string           `json:"country" validate:"required" example:"NG"`
	Description      string           `json:"description" validate:"required"`
	OrganisationSize OrganisationSize `json:"organisation_size" validate:"required"`
	FoundingDate     string           `json:"founding_date" validate:"required"`
}
type CreateOrganisationResponse struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	IsHQ        bool     `json:"is_hq"`
	Branch      []Branch `json:"branches"`
}

// Organisation ...
type Organisation struct {
	ID                 string           `json:"id"`
	Name               string           `json:"name"`
	LogoURL            *string          `json:"logo_url"`
	Website            *string          `json:"website"`
	OrganisationSize   OrganisationSize `json:"organisation_size"`
	ServiceType        string           `json:"service_type"`
	Description        string           `json:"description"`
	RegistrationNumber *string          `json:"registration_number"`
	Rating             float64          `json:"rating"`
	FoundingDate       string           `json:"founding_date"`
	Verified           VerificationType `json:"verified"`
	Branch             []Branch         `json:"branch"`
}

// Organisations ...
type Organisations struct {
	ID          string           `json:"id"`
	Name        string           `json:"name"`
	LogoURL     *string          `json:"logo_url"`
	Description string           `json:"description"`
	Rating      float64          `json:"rating"`
	Verified    VerificationType `json:"verified"`
}

type OrganStatus struct {
	Active bool `json:"active" `
}
