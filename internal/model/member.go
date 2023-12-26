package model

// Member ...
type Member struct {
	BaseP

	UserId         string       `json:"user_id"`
	OrganizationID string       `json:"organization_id" gorm:"primaryKey"`
	RoleID         string       `gorm:"primaryKey"`
	BranchID       *string      `json:"-" gorm:"null"`
	Branch         Branch       `gorm:"-"`
	User           User         `gorm:"-"`
	Role           Role         `gorm:"-"`
	Organisation   Organisation `gorm:"-"`
}
