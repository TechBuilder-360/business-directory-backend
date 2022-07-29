package model

// UserToken ...
type UserToken struct {
	Base

	UserId string `json:"user_id"`
	Token  string `json:"token"`
}
