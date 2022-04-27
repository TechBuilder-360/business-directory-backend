package dto

import "github.com/TechBuilder-360/business-directory-backend/models"

type CreateBranch struct {
	BranchName     string         `json:"branch_name" validate:"required"`
	Contact        models.Contact `json:"contact" `
	Address        models.Address `json:"address"`
}

type Branch struct {
	BranchName string `json:"branch_name"`
	IsHQ bool `json:"is_hq"`
	Address
	Location
	Contact
}
