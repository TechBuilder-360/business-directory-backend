package models

import (
	"time"

	"github.com/google/uuid"
)

type Permission string

// Permissions
const (
	OWNER    Permission = "owner"         // owner has all the privileges
	ADDADMIN Permission = "can_add_admin" // can add new admin to organisation
	CREATEAD Permission = "can_create_ad" // can create advert for organisation
)

type Organisation struct {
	ID                 string   `json:"id" bson:"_id"`
	OrganisationName   string      `json:"organisation_name" bson:"organisation_name"`
	LogoURL            string      `json:"logo_url" bson:"logo_url"`
	CreatorID          uuid.UUID   `json:"creator_id" bson:"creator_id"`
	GeoLocation        Location    `json:"geo_location" bson:"geo_location"`
	Admins             []uuid.UUID `json:"admins" bson:"admins"`
	CreatedAt          time.Time   `json:"created_at" bson:"created_at"`
	UpdatedAt          time.Time   `json:"updated_at" bson:"updated_at"`
	Website            string      `json:"website" bson:"website"`
	OrganisationSize   string      `json:"organisation_size" bson:"organisation_size"`
	Description        string      `json:"description" bson:"description"`
	RegistrationNumber string      `json:"registration_number" bson:"registration_number"`
	Rating             float64     `json:"rating" bson:"rating"`
	FoundingDate       string      `json:"founding_date" bson:"founding_date"`
	Contact            Contact     `json:"contact" bson:"contact"`
	Products           []string    `json:"products" bson:"products"`
	Services           []string    `json:"services" bson:"services"`
	Active             bool        `json:"active" bson:"active"`
	ExpiryDate         time.Time   `json:"expiry_date" bson:"expiry_date"`
}

type Branch struct {
	ID             uuid.UUID   `json:"id" bson:"_id"`
	OrganisationID uuid.UUID   `json:"organisation_id" bson:"organisation_id"`
	BranchName     string      `json:"branch_name" bson:"branch_name"`
	Admins         []uuid.UUID `json:"admins" bson:"admins"`
	Contact        Contact     `json:"contact" bson:"contact"`
	GeoLocation    Location    `json:"geo_location" bson:"geo_location"`
	Address        Address     `json:"address" bson:"address"`
	CreatedAt      time.Time   `json:"created_at" bson:"created_at"`
	IsHQ           bool        `json:"IsHQ" bson:"IsHQ"`
}

type OrganisationMember struct {
	ID         uuid.UUID    `json:"id" bson:"_id"`
	MemberID   uuid.UUID    `json:"member_id" bson:"member_id"`
	BranchID   uuid.UUID    `json:"branch_id" bson:"branch_id"`
	Permission []Permission `json:"permission" bson:"permission"`
}

type Contact struct {
	Email        string         `json:"email" bson:"email"`
	PhoneNumbers []PhoneDetails `json:"phone_numbers" bson:"phone_numbers"`
}

type PhoneDetails struct {
	Type        string `json:"type" bson:"type"`
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
