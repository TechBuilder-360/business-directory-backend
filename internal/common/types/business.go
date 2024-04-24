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

// BusinessReq ...
type BusinessReq struct {
	Name         string       `json:"name" validate:"required"`
	Category     string       `json:"category" validate:"required"`
	Country      string       `json:"country" validate:"required" example:"NG"`
	Description  string       `json:"description" validate:"required"`
	BusinessSize BusinessSize `json:"size" validate:"required"`
	FoundingDate string       `json:"founding_date" validate:"required"`
}

type BusinessResponse struct {
	ID           string       `json:"id"`
	Name         string       `json:"name"`
	Country      string       `json:"country" example:"NG"`
	Category     string       `json:"category"`
	Description  string       `json:"description"`
	BusinessSize BusinessSize `json:"size"`
	IsHQ         bool         `json:"is_hq"`
	Branch       []Branch     `json:"branches"`
}

type OrganStatus struct {
	Active bool `json:"active" `
}
