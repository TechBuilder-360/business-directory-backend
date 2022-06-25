package model

import (
	"time"
)

//type Permission string

//// Permissions
//consts (
//	OWNER    Permission = "owner"         // owner has all the privileges
//	ADDADMIN Permission = "can_add_admin" // can add new admin to organisation
//	CREATEAD Permission = "can_create_ad" // can create advert for organisation
//)

type Organisation struct {
	Base

	CategoryID         string `gorm:"not null"`
	OrganisationName   string `gorm:"column:organisation_name;unique"`
	OwnerID            string `gorm:"not null"`
	LogoURL            *string
	Website            *string   `json:"website"`
	OrganisationSize   *string   `gorm:"column:organisation_size"`
	Description        string    `json:"description"`
	RegistrationNumber *string   `gorm:"column:registration_number"`
	FoundingDate       string    `gorm:"column:founding_date"`
	Active             bool      `json:"active"`
	ExpiryDate         time.Time `gorm:"column:expiry_date"` // Set organisation active status to false
	PublicKey          string    `gorm:"not null;"`          // Used to authenticate business request
	SecretKey          string    `gorm:"not null;"`
	Contact            Contact
	Location           Location // Set to the organisation HQ location
	Rating             Rating
	Admins             []string `json:"admins" gorm:"many2many:business_uploads;"`
	Services           []string `json:"services" gorm:"many2many:business_uploads;"`
	Products           []string `json:"products" gorm:"many2many:business_uploads;"`
}

type OrganisationMember struct {
	Base

	UserID         string `json:"user_id" gorm:"not null"`
	OrganizationID string `json:"organization_id" gorm:"not null"`
	BranchID       string `json:"branch_id"`
}

type Contact struct {
	EmailAddress string `json:"email"`
	PhoneDetails `json:"phone_numbers"`
}

type PhoneDetails struct {
	PhoneNumber string `json:"phone_number"`
	CountryCode string `json:"country_code"`
}

type Location struct {
	Area      string  `json:"area"`
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
