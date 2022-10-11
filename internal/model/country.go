package model

type Country struct {
	Base

	Name         string         `json:"name" gorm:"not null"`
	Code         string         `json:"code" gorm:"not null"`
	Active       *bool          `json:"active" gorm:"not null"`
	Branch       []Branch       `gorm:"-"`
	Organisation []Organisation `gorm:"-"`
}
