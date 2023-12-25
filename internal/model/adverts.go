package model

import (
	"github.com/google/uuid"
)

type Advert struct {
	ID          uuid.UUID `json:"id" bson:"_id"`
	StartDate   string    `json:"start_date" bson:"start_date"`
	EndDate     string    `json:"end_date" bson:"end_date"`
	AdLink      string    `json:"ad_link" bson:"ad_link"`
	Description string    `json:"description" bson:"description"`
	Purpose     string    `json:"purpose" bson:"purpose"`
	Views       uint      `json:"views" bson:"views"`
	UpVote      uint      `json:"up_vote" bson:"up_vote"`
}
