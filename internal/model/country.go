package model

type Country struct {
	Base

	Name        string `json:"name" gorm:"not null"`
	Code        string `json:"code" gorm:"not null"`
	CallingCode string `json:"calling_code"`
	Active      bool   `json:"active" gorm:"not null"`
}
