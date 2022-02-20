package dto

type AuthRequest struct {
	Email string `json:"email"`
	Token string `json:"token"`
}
