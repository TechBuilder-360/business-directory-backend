package services

import (
	"context"
	"errors"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/types"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/utils"
	auth "github.com/TechBuilder-360/business-directory-backend/internal/infrastructure/auth_sdk"
	"github.com/TechBuilder-360/business-directory-backend/internal/model"
	"github.com/TechBuilder-360/business-directory-backend/internal/repository"
	log "github.com/sirupsen/logrus"
)

type IAuthService interface {
	Registration(ctx context.Context, payload types.Registration, logger *log.Entry) error
	ActivateAccount(ctx context.Context, token string, logger *log.Entry) error
	Logout(ctx context.Context, token string) error
	Authenticate(ctx context.Context, payload types.Authenticate) error
	RefreshToken(ctx context.Context, payload *types.RefreshToken) (*types.Authentication, error)
	Login(ctx context.Context, payload types.Authenticate) (*types.LoginResponse, error)
}

type authService struct {
	auth           auth.IAuth
	userRepository repository.UserRepository
}

func NewAuthService() IAuthService {
	return &authService{
		auth:           auth.New(),
		userRepository: repository.NewUserRepository(),
	}
}

func (a *authService) Registration(ctx context.Context, payload types.Registration, logger *log.Entry) error {
	var userID *string

	user, err := a.userRepository.GetByEmail(payload.EmailAddress)
	if err != nil {
		logger.Error("an error occurred ", err.Error())
		return errors.New("request failed")
	}

	if user != nil {
		return errors.New("account already exist, please login")
	}

	u, err := a.auth.GetUserByEmail(ctx, payload.EmailAddress)
	if err != nil {
		logger.Error("an error occurred ", err.Error())
		return errors.New("request failed")
	}

	if u == nil {
		userID, err = a.auth.Registration(ctx, payload)
		if err != nil {
			logger.Error("an error occurred ", err.Error())
			return errors.New("request failed")
		}

		u, err = a.auth.GetUser(ctx, utils.AddToStr(userID))
		if err != nil {
			logger.Error("an error occurred ", err.Error())
			return errors.New("request failed")
		}
	}

	user = &model.User{
		Uid:          u.Id,
		FirstName:    u.FirstName,
		LastName:     u.LastName,
		DisplayName:  u.DisplayName,
		EmailAddress: u.EmailAddress,
		PhoneNumber:  u.PhoneNumber,
		Status:       u.EmailVerified,
	}

	err = a.userRepository.Create(user)
	if err != nil {
		logger.Error("an error occurred when creating user ", err.Error())
		return errors.New("request failed")
	}

	return nil
}

func (a *authService) ActivateAccount(ctx context.Context, token string, logger *log.Entry) error {
	//TODO handles redirect action from email activation
	panic("implement me")
}

func (a *authService) Logout(ctx context.Context, token string) error {
	return a.auth.Logout(ctx, token)
}

func (a *authService) Authenticate(ctx context.Context, payload types.Authenticate) error {
	return a.auth.Authenticate(ctx, payload.EmailAddress)
}

func (a *authService) RefreshToken(ctx context.Context, payload *types.RefreshToken) (*types.Authentication, error) {
	response, err := a.auth.RefreshToken(ctx, payload.Token, payload.RefreshToken)
	if err != nil {
		return nil, err
	}

	return &types.Authentication{
		AccessToken:  response.AccessToken,
		RefreshToken: response.RefreshToken,
		ExpireAt:     response.ExpireAt,
	}, nil
}

func (a *authService) Login(ctx context.Context, payload types.Authenticate) (*types.LoginResponse, error) {
	response, err := a.auth.Login(ctx, payload.EmailAddress, payload.Otp)
	if err != nil {
		return nil, err
	}

	return &types.LoginResponse{
		Profile: types.UserProfile{
			Id:            response.Profile.Id,
			FirstName:     response.Profile.FirstName,
			LastName:      response.Profile.LastName,
			DisplayName:   response.Profile.DisplayName,
			EmailAddress:  response.Profile.EmailAddress,
			PhoneNumber:   response.Profile.PhoneNumber,
			EmailVerified: response.Profile.EmailVerified,
			LastLogin:     response.Profile.LastLogin,
		},
		Authentication: types.Authentication{
			AccessToken:  response.Authentication.AccessToken,
			RefreshToken: response.Authentication.RefreshToken,
			ExpireAt:     response.Authentication.ExpireAt,
		}}, nil

}
