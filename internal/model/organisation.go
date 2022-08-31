package model

import (
	"github.com/TechBuilder-360/business-directory-backend/internal/common/types"
)

//types Permission string

//// Permissions
//constant (
//	OWNER    Permission = "owner"         // owner has all the privileges
//	ADDADMIN Permission = "can_add_admin" // can add new admin to organisation
//	CREATEAD Permission = "can_create_ad" // can create advert for organisation
//)

const (
	MicroSize  types.OrganisationSize = "< 10 employees"
	SmallSize  types.OrganisationSize = "10 - 49 employees"
	MediumSize types.OrganisationSize = "50 - 249 employees"
	LargeSize  types.OrganisationSize = "> 249 employees"
)

type Organisation struct {
	Base

	CategoryID         string `gorm:"not null"`
	CountryID          string `gorm:"not null"`
	OrganisationName   string `gorm:"column:organisation_name;unique"`
	UserID             string `gorm:"not null"`
	LogoURL            *string
	PhoneNumber        string //international format i.e 23481*******1
	EmailAddress       string
	Website            *string
	City               string `json:"city"`
	OrganisationSize   types.OrganisationSize
	Description        string
	RegistrationNumber *string
	FoundingDate       string
	Active             bool
	PublicKey          string
	SecretKey          string
	//Members            []OrganisationMember
	//Services           []OrganisationService
	//Products           []OrganisationProduct
}

type OrganisationService struct {
	Base

	OrganisationID string
	Service        string
	Image          string
}

type OrganisationProduct struct {
	Base

	OrganisationID string
	Product        string
	Image          string
}

type OrganisationMember struct {
	Base

	UserID         string `json:"user_id"`
	OrganizationID string `json:"organization_id" gorm:"primaryKey"`
	RoleID         string `gorm:"primaryKey"`
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
