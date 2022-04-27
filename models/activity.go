package models

import (
	"github.com/google/uuid"
	"time"
)

// Activity ...
type Activity struct {
	ID string `bson:"_id"`
	Message string `bson:"message"`
	By  string `bson:"by"`
	For  string `bson:"For"`
	CreatedAt time.Time `bson:"createdAt"`
}

type Notification struct {
	Base

	UserID uuid.UUID `json:"member_id"`
	Title string `json:"title"`
	Message string `json:"message"`
	Seen bool `json:"seen"`
	IsDeleted bool `json:"is_deleted"`
}

type Token struct {
	Base

	UserID string `json:"user_id"`
	Token string `json:"token"`
	ExpireAt int64 `json:"expire_at"`
}