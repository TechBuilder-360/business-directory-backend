package model

import (
	"time"
)

//types Permission string

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
	PhoneNumber        string
	EmailAddress       string
	Website            *string
	OrganisationSize   *string
	Description        string
	RegistrationNumber *string
	FoundingDate       time.Time
	Active             bool
	ExpiryDate         time.Time
	PublicKey          string
	SecretKey          string
	Location           Location // Set to the organisation HQ location
	Rating             Rating
	Admins             []string
	Services           []string
	Products           []string
}

type OrganisationMember struct {
	Base

	UserID         string `json:"user_id" gorm:"not null"`
	OrganizationID string `json:"organization_id" gorm:"not null"`
	BranchID       string `json:"branch_id"`
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
