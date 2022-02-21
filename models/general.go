package models

import (
	"github.com/google/uuid"
	"time"
)


type Notification struct {
	ID uuid.UUID `json:"id" bson:"_id"`
	MemberID uuid.UUID `json:"member_id" bson:"member_id"`
	Title string `json:"title" bson:"title"`
	Message string `json:"message" bson:"message"`
	Seen bool `json:"seen" bson:"seen"`
	IsDeleted bool `json:"is_deleted" bson:"is_deleted"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	DeletedAt time.Time `json:"deleted_at" bson:"deleted_at"`
}

type Token struct {
	ID uuid.UUID `json:"id" bson:"_id"`
	UserID uuid.UUID `json:"user_id" bson:"user_id"`
	Token string `json:"token" bson:"token"`
	ExpireAt int `json:"expire_at" bson:"expire_at"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}
