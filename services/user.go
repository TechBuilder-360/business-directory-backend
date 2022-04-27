package services

import (
	"errors"
	"github.com/TechBuilder-360/business-directory-backend/dto"
	"github.com/TechBuilder-360/business-directory-backend/repository"
	log "github.com/Toflex/oris_log"
)

//go:generate mockgen -destination=../mocks/services/mockService.go -package=services github.com/TechBuilder-360/business-directory-backend/services UserService
type UserService interface {
	RegisterUser(body *dto.Registration, log log.Logger) (*dto.UserProfile, error)
}

type DefaultUserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &DefaultUserService{repo: repo}
}

func (u *DefaultUserService) RegisterUser(request *dto.Registration, log log.Logger) (*dto.UserProfile, error) {
	// Check if email address exist
	ok,err:=u.repo.DoesUserEmailExist(request.EmailAddress)
	if err != nil {
		log.Error("An Error occurred while checking if user email exist. %s", err.Error())
		return nil, err
	}
	if ok {
		log.Info("Email address already exist. '%s'", request.EmailAddress)
		return nil, errors.New("email address already exist")
	}

	// Save user details
	user, err := u.repo.Create(request)
	if err != nil {
		log.Error("Error occurred when saving new user. %s", err.Error())
		return nil, err
	}

	profile := &dto.UserProfile{
		ID:            user.ID,
		FirstName:     user.FirstName,
		LastName:      user.LastName,
		DisplayName:   user.DisplayName,
		EmailAddress:  user.EmailAddress,
		PhoneNumber:   user.PhoneNumber,
		EmailVerified: false,
		LastLogin:     nil,
	}

	// Activity log
	//activity := &models.Activity{By: userId, Message: "Registered"}
	go func() {
		//if err = u.repo.AddActivity(activity); err!=nil {
		//	log.Error("User activity failed to log", err.Error())
		//}
	}()

	// TODO: Send Activate email

	return profile, nil
}