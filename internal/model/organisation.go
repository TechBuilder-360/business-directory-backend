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

	Category           string `json:"-" gorm:"not null"`
	Country            string `json:"-" gorm:"not null"`
	UserID             string `json:"-" gorm:"not null"`
	OrganisationName   string `gorm:"column:organisation_name;unique"`
	LogoURL            *string
	PhoneNumber        *string //international format i.e 23481*******1
	EmailAddress       string  `json:"email_address" gorm:"not null"`
	Website            *string
	City               string                 `json:"city" gorm:"not null"`
	OrganisationSize   types.OrganisationSize `gorm:"not null"`
	Description        string                 `gorm:"not null"`
	RegistrationNumber *string
	FoundingDate       string
	Rating             float64 `gorm:"default:0"`
	Active             bool    `gorm:"default:false"` // display organisation
	Status             bool    `gorm:"default:false"`
	PublicKey          string  `gorm:"not null"`
	SecretKey          string  `gorm:"not null"`
	User               User    `gorm:"-"`
	Branch             []Branch
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
	BaseP

	UserID         string       `json:"user_id"`
	OrganizationID string       `json:"organization_id" gorm:"primaryKey"`
	RoleID         string       `gorm:"primaryKey"`
	BranchID       *string      `json:"-" gorm:"null"`
	Branch         Branch       `gorm:"-"`
	User           User         `gorm:"-"`
	Role           Role         `gorm:"-"`
	Organisation   Organisation `gorm:"-"`
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
