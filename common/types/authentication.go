package types

// JWTResponse ...
type JWTResponse struct {
	AccessToken string      `json:"access_token"`
	Profile     UserProfile `json:"profile"`
}
