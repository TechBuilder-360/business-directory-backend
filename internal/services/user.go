package services

import (
	"errors"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/types"
	"github.com/TechBuilder-360/business-directory-backend/internal/model"
	"github.com/TechBuilder-360/business-directory-backend/internal/repository"
	log "github.com/sirupsen/logrus"
)

//go:generate mockgen -destination=../mocks/services/mockService.go -package=services github.com/TechBuilder-360/business-directory-backend/services UserService
type UserService interface {
	RegisterUser(body *types.Registration, log log.Entry) (*types.UserProfile, error)
}

type DefaultUserService struct {
	repo repository.UserRepository
}

func NewUserService() UserService {
	return &DefaultUserService{repo: repository.NewUserRepository()}
}

func (u *DefaultUserService) RegisterUser(body *types.Registration, log log.Entry) (*types.UserProfile, error) {
	// Check if email address exist
	ok, err := u.repo.DoesUserEmailExist(body.EmailAddress)
	if err != nil {
		log.Error("An Error occurred while checking if user email exist. %s", err.Error())
		return nil, err
	}
	if ok {
		log.Info("Email address already exist. '%s'", body.EmailAddress)
		return nil, errors.New("email address already exist")
	}

	// Save user details
	user := &model.User{
		FirstName:    body.FirstName,
		LastName:     body.LastName,
		DisplayName:  body.DisplayName,
		EmailAddress: body.EmailAddress,
	}

	err = u.repo.Create(user)
	if err != nil {
		log.Error("Error occurred when saving new user. %s", err.Error())
		return nil, err
	}

	profile := &types.UserProfile{
		ID:            user.ID,
		FirstName:     user.FirstName,
		LastName:      user.LastName,
		DisplayName:   user.DisplayName,
		EmailAddress:  user.EmailAddress,
		PhoneNumber:   user.PhoneNumber,
		EmailVerified: user.EmailVerified,
		LastLogin:     nil,
	}

	// Activity log
	//activity := &model.Activity{By: userId, Message: "Registered"}
	go func() {
		//if err = u.repo.AddActivity(activity); err!=nil {
		//	log.Error("User activity failed to log", err.Error())
		//}
	}()

	// TODO: Send Activate email

	return profile, nil
}
