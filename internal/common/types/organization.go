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
	Country     string `json:"country"`
	ZipCode     int    `json:"zip_code"`
	Street      string `json:"street"`
	City        string `json:"city"`
	State       string `json:"state"`
}

// CreateOrgReq ...
type CreateOrgReq struct {
	OrganisationName string           `json:"organisation_name" validate:"required"`
	OrganisationSize OrganisationSize `json:"organisation_size" validate:"required"`
	UserID           string           `json:"user_id" validate:"required"`
	FoundingDate     string           `json:"founding_date" validate:"required"`
}

// Organisation ...
type Organisation struct {
	OrganisationID     string  `json:"organisation_id"`
	OrganisationName   string  `json:"organisation_name"`
	LogoURL            string  `json:"logo_url"`
	Website            string  `json:"website"`
	OrganisationSize   string  `json:"organisation_size"`
	Description        string  `json:"description"`
	RegistrationNumber string  `json:"registration_number"`
	Rating             float64 `json:"rating"`
	FoundingDate       string  `json:"founding_date"`
}

// Organisations ...
type Organisations struct {
	OrganisationID   string  `json:"organisation_id"`
	OrganisationName string  `json:"organisation_name"`
	LogoURL          string  `json:"logo_url"`
	Description      string  `json:"description"`
	Rating           float64 `json:"rating"`
}

type OrganStatus struct {
	OrganisationID string `json:"organisation_id"`
	Active         bool   `json:"active" `
}

type DataView struct {
	Page     int     `json:"page"`
	Perpage  int64   `json:"perpage"`
	Total    int64   `json:"total"`
	LastPage float64 `json:"last_page"`
}
