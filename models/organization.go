package models

import (
	"time"

	"github.com/google/uuid"
)

//type Permission string

//// Permissions
//const (
//	OWNER    Permission = "owner"         // owner has all the privileges
//	ADDADMIN Permission = "can_add_admin" // can add new admin to organisation
//	CREATEAD Permission = "can_create_ad" // can create advert for organisation
//)

type Organisation struct {
	Base

	OrganisationName   string      `gorm:"column:organisation_name"`
	LogoURL            string      `gorm:"column:logo_url"`
	CreatorID          string      `gorm:"column:creator_id"`
	Website            string      `json:"website"`
	OrganisationSize   string      `gorm:"column:organisation_size"`
	Description        string      `json:"description"`
	RegistrationNumber string      `gorm:"column:registration_number"`
	Rating             float64     `json:"rating"`
	FoundingDate       string      `gorm:"column:founding_date"`
	Active             bool        `json:"active"`
	ExpiryDate         time.Time   `gorm:"column:expiry_date"`
	Contact
	Location
	Admins             []string    `json:"admins" gorm:"many2many:business_uploads;"`
	Services           []string    `json:"services" gorm:"many2many:business_uploads;"`
	Products           []string    `json:"products" gorm:"many2many:business_uploads;"`
}

type Branch struct {
	Base

	OrganisationID string   `gorm:"column:organisation_id"`
	BranchName     string      `gorm:"column:branch_name"`
	IsHQ           bool        `gorm:"column:IsHQ"`
	Address
	Location
	Contact
}

type OrganisationMember struct {
	Base

	UserID     uuid.UUID    `json:"user_id" gorm:"primaryKey"`
	OrganizationID   uuid.UUID    `json:"organization_id" gorm:"primaryKey"`
	BranchID   uuid.UUID    `json:"branch_id"`
}

type Permission struct {
	Base

	Code string
	Description string
}

type Contact struct {
	EmailAddress string         `json:"email" bson:"email"`
	PhoneDetails `json:"phone_numbers" bson:"phone_numbers"`
}

type PhoneDetails struct {
	PhoneNumber string `json:"phone_number" bson:"phone_number"`
	CountryCode string `json:"country_code" bson:"country_code"`
}

type Location struct {
	Longitude float64 `json:"longitude" bson:"longitude"`
	Latitude  float64 `json:"latitude" bson:"latitude"`
}

type Address struct {
	CountryCode string `json:"country_code" bson:"country_code"`
	Country     string `json:"country" bson:"country"`
	ZipCode     int    `json:"zip_code" bson:"zip_code"`
	Street      string `json:"street" bson:"street"`
	City        string `json:"city" bson:"city"`
	State       string `json:"state" bson:"state"`
}
