package model

// Service  ...
type Service struct {
	Base

	OrganisationID string
	Name           string `gorm:"not null"`
	Description    string `gorm:"not null"`
	Image          *string
}