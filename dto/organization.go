package dto

import (
	"time"

	"github.com/TechBuilder-360/business-directory-backend/models"
)


type CreateOrganisation struct {
	OrganisationName string    `json:"organisation_name"`
	OrganisationSize string    `json:"organisation_size"`
	Description      string    `json:"description"`
	FoundingDate     time.Time `json:"founding_date" `
	Active bool 	`json:"active"`
	OrganisationID string    `json:"organisation_id"`
}



type CreateBranch struct {
	
	OrganisationID string   `json:"organisation_id"`
	BranchName     string      `json:"branch_name"`
	Contact        models.Contact     `json:"contact"`
	Address        models.Address     `json:"address"`
	IsHQ           bool        `json:"IsHQ"`
}



type Contact struct {
	Email        string         `json:"email" `
	PhoneNumbers []PhoneDetails `json:"phone_numbers"`
}

type PhoneDetails struct {
	Type        string `json:"type"`
	PhoneNumber string `json:"phone_number" `
	CountryCode string `json:"country_code"`
}

type Location struct {
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
// CreateOrgReq ...
type CreateOrgReq struct {
	OrganisationName string `json:"organisation_name" validate:"required"`
	OrganisationSize string `json:"organisation_size" validate:"required"`
	Description      string `json:"description" validate:"required"`
	FoundingDate     string `json:"founding_date" validate:"required"`
}

// CreateOrgResponse ...
type CreateOrgResponse struct {
	OrganisationID string `json:"organisation_id"`
}