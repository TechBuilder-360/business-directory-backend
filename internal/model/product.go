package model

// Product  ...
type Product struct {
	Base

	OrganisationID string
	Name           string `gorm:"not null"`
	Description    string `gorm:"not null"`
	Image          *string
}
