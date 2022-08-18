package types

import "time"

// JWTResponse ...
type JWTResponse struct {
	AccessToken string      `json:"access_token"`
	Profile     UserProfile `json:"profile"`
}

type MailTemplate struct {
	Header   string
	Code     string
	Token    string
	ToEmail  string
	ToName   string
	Subject  string
	Duration time.Duration
}
