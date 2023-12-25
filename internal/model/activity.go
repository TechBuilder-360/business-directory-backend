package model

import (
	"github.com/google/uuid"
)

// Activity ...
type Activity struct {
	Base

	Message string
	By      string
	For     string
}

type Notification struct {
	Base

	UserID  uuid.UUID `json:"member_id"`
	Title   string    `json:"title"`
	Message string    `json:"message"`
	Seen    bool      `json:"seen"`
}

type Token struct {
	Base

	UserID   string `json:"user_id"`
	Token    string `json:"token"`
	ExpireAt int64  `json:"expire_at"`
}
