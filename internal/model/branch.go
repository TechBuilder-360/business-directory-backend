package model

type Branch struct {
	Base

	OrganisationID string  `gorm:"column:organisation_id"`
	Name           string  `gorm:"column:name"`
	IsHQ           bool    `gorm:"column:IsHQ"`
	Active         bool    `gorm:"default:true"`
	PhoneNumber    *string `json:"phone_number"`
	CountryID      string  `json:"country_id"`
	ZipCode        *string `json:"zip_code"`
	Street         *string `json:"street"`
	City           *string `json:"city"`
	State          *string `json:"state"`
	Longitude      float64 `json:"longitude"`
	Latitude       float64 `json:"latitude"`
}
