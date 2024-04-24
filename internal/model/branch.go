package model

type Branch struct {
	Base

	BusinessID  string  `gorm:"business_id"`
	Name        string  `gorm:"name"`
	IsHQ        bool    `gorm:"is_HQ"`
	Active      bool    `gorm:"default:false"`
	PhoneNumber *string `json:"phone_number"`
	CountryID   string  `json:"country_id"`
	ZipCode     *string `json:"zip_code"`
	Street      *string `json:"street"`
	City        *string `json:"city"`
	State       *string `json:"state"`
	Longitude   float64 `json:"longitude"`
	Latitude    float64 `json:"latitude"`
}
