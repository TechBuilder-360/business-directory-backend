package dto


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