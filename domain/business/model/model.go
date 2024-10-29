package model

import (
	model2 "github.com/TechBuilder-360/business-directory-backend/domain/user/model"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/types"
	"github.com/TechBuilder-360/business-directory-backend/internal/model"
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

	OWNER             = "Owner"
	OrganisationAdmin = "Organisation Admin"
	BranchManager     = "Branch Manager"
)

type Business struct {
	model.Base

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

type Member struct {
	model.Base

	UserId     string       `json:"user_id"`
	BusinessID string       `json:"business_id" gorm:"primaryKey"`
	RoleID     string       `gorm:"primaryKey"`
	Status     types.Status `json:"status"`
	User       model2.User  `gorm:"-"`
	Role       Role         `gorm:"-"`
	Business   Business     `gorm:"-"`
}

// Service  ...
type Service struct {
	model.Base

	OrganisationID string
	Name           string `gorm:"not null"`
	Description    string `gorm:"not null"`
	Image          *string
}

type Branch struct {
	model.Base

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

// Product  ...
type Product struct {
	model.Base

	OrganisationID string
	Name           string `gorm:"not null"`
	Description    string `gorm:"not null"`
	Image          *string
}

type Role struct {
	model.Base

	Name types.RoleType `gorm:"unique"`
	Permission
}

type Category struct {
	model.Base

	Name string `gorm:"not null;unique"`
}

type Permission struct {
	model.Base

	Code        string
	Description string
}

type Rating struct {
	model.Base

	OrganisationID string `gorm:"not null;unique"`
	A              uint64 `gorm:"column:a_star;default:0;"`
	B              uint64 `gorm:"column:b_star;default:0;"`
	C              uint64 `gorm:"column:c_star;default:0;"`
	D              uint64 `gorm:"column:d_star;default:0;"`
	E              uint64 `gorm:"column:e_star;default:0;"`
}

func (r Rating) Summation() uint64 {
	return r.A + r.B + r.C + r.D + r.E
}

func (r Rating) AverageRating() float64 {
	return float64(((1 * r.A) + (2 * r.B) + (3 * r.C) + (4 * r.D) + (5 * r.E)) / r.Summation())
}
