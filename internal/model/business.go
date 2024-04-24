package model

import (
	"github.com/TechBuilder-360/business-directory-backend/internal/common/types"
)

//types Permission string

//// Permissions
//constant (
//	OWNER    Permission = "owner"         // owner has all the privileges
//	ADDADMIN Permission = "can_add_admin" // can add new admin to Business
//	CREATEAD Permission = "can_create_ad" // can create advert for Business
//)

const (
	MicroSize  types.BusinessSize = "< 10 employees"
	SmallSize  types.BusinessSize = "10 - 49 employees"
	MediumSize types.BusinessSize = "50 - 249 employees"
	LargeSize  types.BusinessSize = "> 249 employees"
)

type Business struct {
	Base

	Category           string `json:"-" gorm:"not null"`
	CountryID          string `json:"-" gorm:"not null"`
	Name               string `gorm:"column:name;unique"`
	LogoURL            *string
	PhoneNumber        *string //international format i.e 23481*******1
	SupportPhoneNumber *string //international format i.e 23481*******1
	EmailAddress       string  `json:"email_address" gorm:"not null;unique"`
	Website            *string
	BusinessSize       types.BusinessSize     `gorm:"not null"`
	Description        string                 `gorm:"not null"`
	RegistrationNumber *string                `json:"registration_number" gorm:"null;unique"`
	Location           types.LocationType     `gorm:"not null;REMOTE"`
	Verified           types.VerificationType `gorm:"default:UNVERIFIED"`
	ServiceType        string                 `json:"service_type"`
	FoundingDate       string
	Rating             float64   `json:"rating"`
	Active             bool      `gorm:"default:false"`
	PublicKey          string    `gorm:"not null"`
	SecretKey          string    `gorm:"not null"`
	Members            []Member  `gorm:"-"`
	Branch             []Branch  `gorm:"-"`
	Services           []Service `gorm:"-"`
	Products           []Product `gorm:"-"`
}
