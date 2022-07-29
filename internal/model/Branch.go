package model

type Branch struct {
	Base

	OrganisationID string `gorm:"column:organisation_id"`
	BranchName     string `gorm:"column:branch_name"`
	IsHQ           bool   `gorm:"column:IsHQ"`
	Address
	Location
	Contact
}
