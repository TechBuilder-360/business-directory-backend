package models

import (
	"time"
)

type UserProfile struct {
	ID            string    `json:"id" bson:"_id"`
	FirstName     string    `json:"first_name" bson:"first_name"`
	LastName      string    `json:"last_name" bson:"last_name"`
	DisplayName   string    `json:"display_name" bson:"display_name"`
	EmailAddress  string    `json:"email_address" bson:"email_address"`
	PhoneNumber   string    `json:"phone_number" bson:"phone_number"`
	EmailVerified bool      `json:"email_verified" bson:"email_verified"`
	LastLogin     time.Time `json:"last_login" bson:"last_login"`
	CreatedAt     time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" bson:"updated_at"`
}
