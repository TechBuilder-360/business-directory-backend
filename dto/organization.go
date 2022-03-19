package dto

import "time"

type CreateOrganisation struct {
	OrganisationName string    `json:"organisation_name"`
	OrganisationSize string    `json:"organisation_size"`
	Description      string    `json:"description"`
	FoundingDate     time.Time `json:"founding_date" `
}
