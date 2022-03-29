package models

import "time"

// Activity ...
type Activity struct {
	ID string `bson:"_id"`
	Message string `bson:"message"`
	By  string `bson:"by"`
	For  string `bson:"For"`
	CreatedAt time.Time `bson:"createdAt"`
}
