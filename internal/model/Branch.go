package model

type Branch struct {
	Base

	OrganisationID string `gorm:"column:organisation_id"`
	Name           string `gorm:"column:name"`
	IsHQ           bool   `gorm:"column:is_hq"`
	Address
	Location
}
