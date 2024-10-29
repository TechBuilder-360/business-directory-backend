package model

import "github.com/TechBuilder-360/business-directory-backend/internal/model"

type Country struct {
	model.Base

	Name        string `json:"name" gorm:"not null"`
	Code        string `json:"code" gorm:"not null"`
	CallingCode string `json:"calling_code"`
	Active      bool   `json:"active" gorm:"not null"`
}
