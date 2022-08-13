package services

import (
	"errors"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/types"
	"github.com/TechBuilder-360/business-directory-backend/internal/infrastructure/sendgrid"
	"github.com/TechBuilder-360/business-directory-backend/internal/model"
	"github.com/TechBuilder-360/business-directory-backend/internal/repository"
	log "github.com/sirupsen/logrus"
)

//go:generate mockgen -destination=../mocks/services/mockService.go -package=services github.com/TechBuilder-360/business-directory-backend/services UserService
type UserService interface {
	RegisterUser(body *types.Registration, log log.Entry) (string, error)
	ActivateEmail(email string, log log.Entry) (string, error)
}

type DefaultUserService struct {
	repo     repository.UserRepository
	activity repository.ActivityRepository
}

func NewUserService() UserService {
	return &DefaultUserService{repo: repository.NewUserRepository()}
}

func (u *DefaultUserService) ActivateEmail(email string, log log.Entry) (string, error) {

	// Update with conditions
	if err := u.repo.Activate(email); err != nil {
		log.Error("An Error occurred while Activating your account, Please try again. %s", err.Error())
		return "", err
	}

	return " Your account has been activated, Please proceed to login", nil
}
func (u *DefaultUserService) RegisterUser(body *types.Registration, log log.Entry) (string, error) {
	// Check if email address exist
	ok, err := u.repo.DoesUserEmailExist(body.EmailAddress)
	if err != nil {
		log.Error("An Error occurred while checking if user email exist. %s", err.Error())
		return "", err
	}
	if ok {
		log.Info("Email address already exist. '%s'", body.EmailAddress)
		return "", errors.New("email address already exist")
	}

	// Save user details
	user := &model.User{
		FirstName:    body.FirstName,
		LastName:     body.LastName,
		DisplayName:  body.DisplayName,
		EmailAddress: body.EmailAddress,
	}

	//err = u.repo.Create(user)
	//if err != nil {
	//	log.Error("Error occurred when saving new user. %s", err.Error())
	//	return "", err
	//}

	//profile := &types.UserProfile{
	//	ID:            user.ID,
	//	FirstName:     user.FirstName,
	//	LastName:      user.LastName,
	//	DisplayName:   user.DisplayName,
	//	EmailAddress:  user.EmailAddress,
	//	PhoneNumber:   user.PhoneNumber,
	//	EmailVerified: user.EmailVerified,
	//	LastLogin:     nil,
	//}

	// Activity log
	activity := &model.Activity{By: user.ID, Message: "Registered"}
	go func() {
		if err = u.activity.Get(activity); err != nil {
			log.Error("User activity failed to log", err.Error())
		}
	}()

	//TODO: Send Activate email
	err = sendgrid.SendMail("Activate your account", body.EmailAddress, "", body.DisplayName)
	if err != nil {
		log.Error("Error occurred when sending activation email. %s", err.Error())
		return "", err
	}
	return "Account created successfully, Please check your mail to activate Account ", nil
}
