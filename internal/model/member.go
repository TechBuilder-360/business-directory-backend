package model

import "github.com/TechBuilder-360/business-directory-backend/internal/common/types"

// Member ...
type Member struct {
	Base

	UserId     string       `json:"user_id"`
	BusinessID string       `json:"business_id" gorm:"primaryKey"`
	RoleID     string       `gorm:"primaryKey"`
	Status     types.Status `json:"status"`
	User       User         `gorm:"-"`
	Role       Role         `gorm:"-"`
	Business   Business     `gorm:"-"`
}
