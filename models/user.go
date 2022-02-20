package models

import (
	"github.com/google/uuid"
	"time"
)

type Profile struct {
	ID           uuid.UUID `json:"id" bson:"_id"`
	FirstName    string    `json:"first_name" bson:"first_name"`
	LastName     string    `json:"last_name" bson:"last_name"`
	DisplayName  string    `json:"display_name" bson:"display_name"`
	EmailAddress string    `json:"email_address" bson:"email_address"`
	LastLogin    time.Time `json:"last_login" bson:"last_login"`
	CreatedAt    time.Time `json:"created_at" bson:"created_at"`
}
