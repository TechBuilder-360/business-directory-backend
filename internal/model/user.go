package model

// User ...
type User struct {
	Base

	Uid          string  `json:"uid" gorm:"not null;unique"`
	FirstName    string  `json:"first_name" gorm:"not null"`
	LastName     string  `json:"last_name" gorm:"not null"`
	DisplayName  string  `json:"display_name" gorm:"not null"`
	EmailAddress string  `json:"email_address" gorm:"not null"`
	PhoneNumber  string  `json:"phone_number" gorm:"null"`
	Avatar       *string `json:"avatar" gorm:"null"`
	Status       bool    `json:"status" gorm:"not null"`
}
