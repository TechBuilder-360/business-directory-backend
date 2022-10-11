package types

type CreateBranchRequest struct {
	Name        string  `json:"name" validate:"required"`
	ZipCode     int     `json:"zip_code" validate:"required"`
	Street      string  `json:"street" validate:"required"`
	City        string  `json:"city" validate:"required"`
	State       string  `json:"state" validate:"required"`
	CountryCode string  `json:"country_code" validate:"required"`
	PhoneNumber string  `json:"phone_number"`
	Longitude   float64 `json:"longitude"`
	Latitude    float64 `json:"latitude"`
}

type Branch struct {
	Name        string `json:"name"`
	IsHQ        bool   `json:"is_hq"`
	PhoneNumber *string
	Country     string  `json:"country"`
	ZipCode     *string `json:"zip_code"`
	Street      *string `json:"street"`
	City        *string `json:"city"`
	State       *string `json:"state"`
	Longitude   float64 `json:"longitude"`
	Latitude    float64 `json:"latitude"`
}
