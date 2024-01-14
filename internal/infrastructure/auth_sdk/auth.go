package auth_server_sdk

import (
	"context"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/types"
	"github.com/go-resty/resty/v2"
)

type auth struct {
}

type IAuth interface {
	Registration(ctx context.Context, payload types.Registration) (*string, error)
	ActivateEmail(ctx context.Context, token string) error
	Authenticate(ctx context.Context, emailAddress string) error
	Login(ctx context.Context, emailAddress, otp string) (*Login, error)
	RefreshToken(ctx context.Context, token, refreshToken string) (*Authentication, error)
	Logout(ctx context.Context, token string) error
	ValidateToken(ctx context.Context, token string) bool
	GetUser(ctx context.Context, id string) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
}

func New() IAuth {
	return &auth{}
}

var client *resty.Client

func init() {
	client = resty.New()
}
