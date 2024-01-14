package auth_server_sdk

import (
	"context"
	"errors"
	"fmt"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/types"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/utils"
	"github.com/TechBuilder-360/business-directory-backend/internal/configs"
	log "github.com/sirupsen/logrus"
	"time"
)

type RegistrationResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    struct {
		UserId string `json:"user_id"`
	} `json:"data"`
}

type refreshTokenRequest struct {
	Token        string `json:"token" validate:"required"`
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type LoginResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    Login  `json:"data"`
}

type Login struct {
	Authentication Authentication `json:"authentication"`
	Profile        struct {
		Id            string    `json:"id"`
		FirstName     string    `json:"first_name"`
		LastName      string    `json:"last_name"`
		DisplayName   string    `json:"display_name"`
		EmailAddress  string    `json:"email_address"`
		PhoneNumber   string    `json:"phone_number"`
		EmailVerified bool      `json:"email_verified"`
		LastLogin     time.Time `json:"last_login"`
	} `json:"profile"`
}

type AuthResponse struct {
	Status  bool           `json:"status"`
	Message string         `json:"message"`
	Data    Authentication `json:"data"`
}

type Authentication struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpireAt     int64  `json:"expire_at"`
}

type emailRequest struct {
	EmailAddress string `json:"email_address"`
}

type authRequest struct {
	EmailAddress string `json:"email_address"`
	Otp          string `json:"otp"`
}

type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type Error struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type UserResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    User   `json:"data"`
}

type User struct {
	Id            string    `json:"id"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	DisplayName   string    `json:"display_name"`
	EmailAddress  string    `json:"email_address"`
	PhoneNumber   string    `json:"phone_number"`
	EmailVerified bool      `json:"email_verified"`
	LastLogin     time.Time `json:"last_login"`
}

func (c *auth) Registration(ctx context.Context, payload types.Registration) (*string, error) {
	result := new(RegistrationResponse)
	e := new(Error)

	resp, err := client.R().
		EnableTrace().
		SetHeader("x-auth", utils.AddToStr(configs.Instance.AuthServerSecretKey)).
		SetContext(ctx).
		SetBody(&payload).
		SetResult(result).
		SetError(e).
		Post(fmt.Sprintf("%s/auth/register", utils.AddToStr(configs.Instance.AuthServerBaseURL)))

	if err != nil {
		log.Error("AUTH:: registration error %+v", err)
		return nil, errors.New("request failed")
	}
	if resp.IsError() {
		log.Error("AUTH:: registration error %+v", e)
		return nil, errors.New(e.Message)
	}

	return &result.Data.UserId, nil
}

func (c *auth) ActivateEmail(ctx context.Context, token string) error {
	result := new(Response)
	e := new(Error)

	resp, err := client.R().
		EnableTrace().
		SetHeader("x-auth", utils.AddToStr(configs.Instance.AuthServerSecretKey)).
		SetContext(ctx).
		SetResult(result).
		SetError(result).
		SetPathParam("token", token).
		Get(fmt.Sprintf("%s/auth/activate/{token}", utils.AddToStr(configs.Instance.AuthServerBaseURL)))

	if err != nil {
		log.Error("AUTH:: activate email error %+v", err)
		return errors.New("request failed")
	}
	if resp.IsError() {
		log.Error("AUTH:: activate email error %+v", e)
		return errors.New(e.Message)
	}

	return nil
}

func (c *auth) Authenticate(ctx context.Context, emailAddress string) error {
	result := new(Response)
	e := new(Error)

	payload := emailRequest{
		EmailAddress: emailAddress,
	}

	resp, err := client.R().
		EnableTrace().
		SetHeader("x-auth", utils.AddToStr(configs.Instance.AuthServerSecretKey)).
		SetContext(ctx).
		SetBody(&payload).
		SetResult(result).
		SetError(result).
		Post(fmt.Sprintf("%s/auth/authentication", utils.AddToStr(configs.Instance.AuthServerBaseURL)))

	if err != nil {
		log.Error("AUTH:: authenticate error %+v", err)
		return errors.New("request failed")
	}
	if resp.IsError() {
		log.Error("AUTH:: authenticate error %+v", e)
		return errors.New(e.Message)
	}

	return nil
}

func (c *auth) Login(ctx context.Context, emailAddress, otp string) (*Login, error) {
	result := new(LoginResponse)
	e := new(Error)

	payload := authRequest{
		EmailAddress: emailAddress,
		Otp:          otp,
	}

	resp, err := client.R().
		EnableTrace().
		SetHeader("x-auth", utils.AddToStr(configs.Instance.AuthServerSecretKey)).
		SetContext(ctx).
		SetBody(&payload).
		SetResult(result).
		SetError(result).
		Post(fmt.Sprintf("%s/auth/login", utils.AddToStr(configs.Instance.AuthServerBaseURL)))

	if err != nil {
		log.Error("AUTH:: login error %+v", err)
		return nil, errors.New("request failed")
	}
	if resp.IsError() {
		log.Error("AUTH:: login error %+v", e)
		return nil, errors.New(e.Message)
	}

	return &result.Data, nil
}

func (c *auth) RefreshToken(ctx context.Context, token, refreshToken string) (*Authentication, error) {
	result := new(AuthResponse)
	e := new(Error)

	payload := refreshTokenRequest{
		Token:        token,
		RefreshToken: refreshToken,
	}

	resp, err := client.R().
		EnableTrace().
		SetHeader("x-auth", utils.AddToStr(configs.Instance.AuthServerSecretKey)).
		SetContext(ctx).
		SetBody(&payload).
		SetResult(result).
		SetError(result).
		Post(fmt.Sprintf("%s/auth/refresh", utils.AddToStr(configs.Instance.AuthServerBaseURL)))

	if err != nil {
		log.Error("AUTH:: refresh token error %+v", err)
		return nil, errors.New("request failed")
	}
	if resp.IsError() {
		log.Error("AUTH:: refresh token error %+v", e)
		return nil, errors.New(e.Message)
	}

	return &result.Data, nil
}

func (c *auth) Logout(ctx context.Context, token string) error {
	result := new(Response)
	e := new(Error)

	resp, err := client.R().
		EnableTrace().
		SetAuthToken(token).
		SetHeader("x-auth", utils.AddToStr(configs.Instance.AuthServerSecretKey)).
		SetContext(ctx).
		SetResult(result).
		SetError(result).
		Put(fmt.Sprintf("%s/auth/logout", utils.AddToStr(configs.Instance.AuthServerBaseURL)))

	if err != nil {
		log.Error("AUTH:: logout error %+v", err)
		return errors.New("request failed")
	}
	if resp.IsError() {
		log.Error("AUTH:: logout error %+v", e)
		return errors.New(e.Message)
	}

	return nil
}

func (c *auth) ValidateToken(ctx context.Context, token string) bool {
	result := new(Response)

	resp, err := client.R().
		EnableTrace().
		SetAuthToken(token).
		SetHeader("x-auth", utils.AddToStr(configs.Instance.AuthServerSecretKey)).
		SetContext(ctx).
		SetResult(result).
		SetError(result).
		Get(fmt.Sprintf("%s/auth/validate-token", utils.AddToStr(configs.Instance.AuthServerBaseURL)))

	if err != nil || resp.IsError() {
		return false
	}

	return true
}

func (c *auth) GetUser(ctx context.Context, id string) (*User, error) {
	result := new(UserResponse)
	e := new(Error)

	resp, err := client.R().
		EnableTrace().
		SetHeader("x-auth", utils.AddToStr(configs.Instance.AuthServerSecretKey)).
		SetContext(ctx).
		SetResult(result).
		SetError(result).
		SetPathParam("id", id).
		Get(fmt.Sprintf("%s/users/{id}", utils.AddToStr(configs.Instance.AuthServerBaseURL)))

	if err != nil {
		log.Error("AUTH:: fetch user error %+v", err)
		return nil, errors.New("request failed")
	}
	if resp.IsError() {
		log.Error("AUTH:: fetch user error %+v", e)
		return nil, errors.New(e.Message)
	}

	return &result.Data, nil
}

func (c *auth) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	result := new(UserResponse)
	e := new(Error)

	resp, err := client.R().
		EnableTrace().
		SetHeader("x-auth", utils.AddToStr(configs.Instance.AuthServerSecretKey)).
		SetContext(ctx).
		SetResult(result).
		SetError(result).
		SetQueryParam("email", email).
		Get(fmt.Sprintf("%s/users", utils.AddToStr(configs.Instance.AuthServerBaseURL)))

	if err != nil {
		log.Error("AUTH:: fetch user error %+v", err)
		return nil, errors.New("request failed")
	}
	if resp.IsError() {
		log.Error("AUTH:: fetch user error %+v", e)
		return nil, errors.New(e.Message)
	}

	return &result.Data, nil
}
